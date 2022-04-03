package email

import (
	"errors"

	"github.com/AdeAttwood/SoleMail/pkg/config"
	"github.com/AdeAttwood/SoleMail/pkg/database"
)

type EmailContent struct {
	// The text content on an email
	Text string
	// The html content of an email.
	Html string
}

type EmailService interface {
	// Import all of the email into the database from an account. This will
	// update all the email data for all of the email in the account preserving
	// tags for any email that has already been tagged
	Import() error
	// Gets the content of an email from the service. This will need to get it
	// for each message as that raw content is not stored in the internal
	// databse only the subject is stored in the database and an index of some
	// content for optimising search
	GetContent(database.Message) (EmailContent, error)
}

// Creates a new instance of the email service from the type of account.
// NOTE: Only maildir is currently supported
func Create(c *config.Account, db *database.Database) (EmailService, error) {
	switch c.Type {
	case string(config.AccountTypeMaildir):
		return NewMailDirService(c, db)
	default:
		return nil, errors.New("Invalid Email service")
	}
}
