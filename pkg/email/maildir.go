package email

import (
	"errors"
	"io/ioutil"
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
	mailboxes, err := ioutil.ReadDir(md.account.Directory)
	if err != nil {
		return content, err
	}

	for _, mailbox := range mailboxes {
		if strings.HasSuffix(mailbox.Name(), ".cmeta") {
			continue
		}

		maildir := Dir(md.account.Directory + "/" + mailbox.Name())
		email, err := parseMessage(maildir, message.MessageKey)
		if err != nil {
			return content, err
		}

		content.Html = email.HTMLBody
		content.Text = email.TextBody

		return content, nil
	}

	return content, errors.New("Message not found")
}

func (md MaildirService) Import() error {
	mailboxes, err := ioutil.ReadDir(md.account.Directory)
	if err != nil {
		return err
	}

	for _, mailbox := range mailboxes {
		// TODO(ade): Fix this quick hack to exclude Evolution mails cmeta
		// directory as it is not a valid maildir. We should probably skip any
		// directory that dose not contain the {tmp,cur,new} directories.
		if strings.HasSuffix(mailbox.Name(), ".cmeta") {
			continue
		}

		maildir := Dir(md.account.Directory + "/" + mailbox.Name())
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

			message := database.Message{
				MessageID:  email.MessageID,
				Account:    md.account.ID,
				MessageKey: key,
				Subject:    email.Subject,
				FromName:   email.From[0].Name,
				From:       email.From[0].Address,
				Date:       email.Date,
				Tags:       []string{},
			}

			for _, message_id := range email.References {
				message.References = append(message.References, message_id)
			}

			for _, reply_to := range email.InReplyTo {
				message.References = append(message.References, reply_to)
			}

			md.db.AddMessage(message)
		}
	}

	md.db.UpdateThreads()

	return nil
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
