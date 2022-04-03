package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Config struct {
	// A list of accounts this application has access to
	Accounts []Account
}

// Defines the type of account for getting mail from imap or maildir
type AccountType string

const (
	// Maildir account type https://en.wikipedia.org/wiki/Maildir
	AccountTypeMaildir AccountType = "Maildir"
)

type Account struct {
	// The id of the account.
	ID string
	// The type of account
	Type string
	// The name of the user that will be using the email. This will be used in
	// the From address
	Name string
	// The directory where your maildir mailbox is
	Directory string
	// The email address associated to this account
	Email string
	// The username for the account. Most of the time this is
	UserName string
	// The password for accessing the account
	Password string
}

// Get the users configuration a directory if one has been saved then it will
// load that.
func NewConfig(path string) Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	if _, err := os.Stat(path + "/config.json"); os.IsNotExist(err) {
		return Config{}
	}

	json_file, err := os.Open(path + "/config.json")
	if err != nil {
		panic(err)
	}

	config := Config{}

	json_content, _ := ioutil.ReadAll(json_file)
	json.Unmarshal(json_content, &config)

	return config
}

func (c Config) GetAccount(name string) (Account, error) {
	for _, account := range c.Accounts {
		if account.ID == name {
			return account, nil
		}
	}

	return Account{}, errors.New("Email account not found")
}
