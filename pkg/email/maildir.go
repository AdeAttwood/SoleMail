package email

import (
	"errors"
	"io/ioutil"
	"net/mail"
	"strings"

	"github.com/AdeAttwood/SoleMail/pkg/config"
	"github.com/AdeAttwood/SoleMail/pkg/database"
	"github.com/DusanKasan/parsemail"

	"github.com/emersion/go-maildir"
	. "github.com/emersion/go-maildir"
)

func NewMailDirService(account *config.Account, db *database.Database) (EmailService, error) {
	return MaildirService{
		account: account,
		db:      db,
	}, nil
}

type MaildirService struct {
	account *config.Account
	db      *database.Database
}

func (md MaildirService) GetContent(message database.Message) (EmailContent, error) {
	content := EmailContent{}
	mailboxes, err := md.getMailboxes()
	if err != nil {
		return content, err
	}

	for _, mailbox := range mailboxes {
		maildir := Dir(mailbox)
		email, err := parseMessage(maildir, message.MessageKey)
		if err != nil {
			continue
		}

		content.Html = email.HTMLBody
		content.Text = email.TextBody

		return content, nil
	}

	return content, errors.New("Message not found")
}

func (md MaildirService) Update() error {
	mbsync := NewMbSync(md.account)
	err := mbsync.Sync()
	if err != nil {
		return err
	}

	mailboxes, err := md.getMailboxes()
	if err != nil {
		return err
	}

	for _, mailbox := range mailboxes {
		maildir := Dir(mailbox)
		keys, err := maildir.Unseen()
		if err != nil {
			return err
		}

		for _, key := range keys {
			email, err := parseMessage(maildir, key)
			if err != nil {
				// TODO(ade): Sort out when the email fails to parse I had an
				// email that was UTF-7 encoded and failed to decode, Currently
				// this is just skipped with no warning.
				continue
			}

			message := md.toMessage(email, key)
			thread_id := md.db.GetThreadId(message)
			if thread_id == -1 {
				thread_id = md.db.NextThreadID()
			}

			thread, err := md.db.GetThread(thread_id)
			if err != nil {
				thread = database.Thread{
					ThreadID: thread_id,
					Date:     message.Date,
					Subject:  message.Subject,
					FromName: message.FromName,
					From:     message.From,
					Tags:     message.Tags,
					Messages: []string{},
				}
			}

			thread.Messages = append(thread.Messages, message.MessageID)
			thread.Tags = database.ApplyTags(thread.Tags, "+inbox +unread")
			if message.Date.After(thread.Date) {
				thread.Date = message.Date
			}

			md.db.WriteMessage(message)
			md.db.WriteThread(thread)

			// Ensure all of the messages have a entry in the database so we can
			// thread messages that we do not have the entire thread
			for _, reference := range message.References {
				_, err := md.db.GetMessage(reference)
				if err != nil {
					md.db.WriteMessage(database.Message{MessageID: reference, ThreadID: thread_id})
				}
			}
		}
	}

	return nil
}

func (md MaildirService) Import() error {
	mailboxes, err := md.getMailboxes()
	if err != nil {
		return err
	}

	for _, mailbox := range mailboxes {
		maildir := Dir(mailbox)
		keys, err := maildir.Keys()
		if err != nil {
			return err
		}

		for _, key := range keys {
			email, err := parseMessage(maildir, key)
			if err != nil {
				// TODO(ade): Sort out when the email fails to parse I had an
				// email that was UTF-7 encoded and failed to decode, Currently
				// this is just skipped with no warning.
				continue
			}

			message := md.toMessage(email, key)
			md.db.WriteMessage(message)

			// Ensure all of the messages have a entry in the database so we can
			// thread messages that we do not have the entire thread
			for _, reference := range message.References {
				_, err := md.db.GetMessage(reference)
				if err != nil {
					md.db.WriteMessage(database.Message{MessageID: reference})
				}
			}
		}
	}

	md.db.UpdateThreads()

	return nil
}

func (md MaildirService) toMessage(email parsemail.Email, key string) database.Message {
	if len(email.From) == 0 {
		email.From = []*mail.Address{
			{
				Name:    "Unknown",
				Address: "unknown@example.com",
			},
		}
	}

	message := database.Message{
		MessageID:  email.MessageID,
		Account:    md.account.ID,
		MessageKey: key,
		Subject:    email.Subject,
		FromName:   email.From[0].Name,
		From:       email.From[0].Address,
		Date:       email.Date,
		Tags:       []string{"inbox", "unread"},
	}

	for _, message_id := range email.References {
		message.References = append(message.References, message_id)
	}

	for _, reply_to := range email.InReplyTo {
		// TODO(ade): fix parsing of in reply to header this is a propa hack
		if strings.Contains(reply_to, "@") && !strings.Contains(reply_to, "(") {
			message.References = append(message.References, reply_to)
		}
	}

	return message
}

func parseMessage(md maildir.Dir, key string) (parsemail.Email, error) {
	reader, err := md.Open(key)
	if err != nil {
		return parsemail.Email{}, err
	}

	email, err := parsemail.Parse(reader)
	reader.Close()

	return email, err
}

// Gets all of the valid maildir mailboxes in the account. This will be each
// folder in the user email account
func (md MaildirService) getMailboxes() ([]string, error) {
	mailboxes := []string{}
	folders, err := ioutil.ReadDir(md.account.Directory)
	if err != nil {
		return mailboxes, err
	}

	for _, folder := range folders {
		// TODO(ade): Fix this quick hack to exclude Evolution mails cmeta
		// directory as it is not a valid maildir. We should probably skip any
		// directory that dose not contain the {tmp,cur,new} directories.
		if strings.HasSuffix(folder.Name(), ".cmeta") {
			continue
		}

		if folder.Name() == "[Gmail]" {
			continue
		}

		mailboxes = append(mailboxes, md.account.Directory+"/"+folder.Name())
	}

	return mailboxes, nil
}
