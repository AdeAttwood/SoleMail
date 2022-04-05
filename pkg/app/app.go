package app

import (
	"context"

	"github.com/AdeAttwood/SoleMail/pkg/config"
	"github.com/AdeAttwood/SoleMail/pkg/database"
	"github.com/AdeAttwood/SoleMail/pkg/email"
)

// App struct
type App struct {
	ctx context.Context
	db  *database.Database
}

// NewApp creates a new App application struct
func NewApp() *App {
	d, err := database.Open(config.GetDefaultDataPath())
	if err != nil {
		panic(err)
	}

	return &App{
		db: d,
	}
}

// startup is called at application startup
func (b *App) Startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) DomReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Gets the users current configuration
func (b *App) GetConfig() config.Config {
	return config.NewConfig(config.GetDefaultConfigPath())
}

// Gets a list of messages
func (b *App) GetMessages() ([]database.Message, error) {
	return b.db.List()
}

// Gets a list of messages
func (b *App) GetThreads(query string) ([]database.Thread, error) {
	return b.db.ListThreads(query)
}

func (b *App) ReadThread(thread_id int) ([]database.Message, error) {
	return b.db.ReadThread(thread_id)
}

func (b *App) TagThread(thread_id int, tag_string string) (database.Thread, error) {
	return b.db.TagThread(thread_id, tag_string)
}

func (b *App) TagQuery(query string, tag_string string) error {
	return b.db.TagQuery(query, tag_string)
}

func (b *App) GetMessageContent(message_id string) (email.EmailContent, error) {
	message, err := b.db.GetMessage(message_id)
	if err != nil {
		return email.EmailContent{}, err
	}

	// TODO(ade): Add the account name into the message in the internal database
	account, err := b.GetConfig().GetAccount("Gmail One")
	if err != nil {
		return email.EmailContent{}, err
	}

	service, err := email.Create(&account, b.db)
	if err != nil {
		return email.EmailContent{}, err
	}

	return service.GetContent(message)
}
