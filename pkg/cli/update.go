package cli

import (
	"github.com/AdeAttwood/SoleMail/pkg/config"
	"github.com/AdeAttwood/SoleMail/pkg/database"
	"github.com/AdeAttwood/SoleMail/pkg/email"
)

func Update() {
	user_config := config.NewConfig(config.GetDefaultConfigPath())
	db, err := database.Open(config.GetDefaultDataPath())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	for _, account := range user_config.Accounts {
		service, err := email.Create(&account, db)
		if err != nil {
			panic(err)
		}

		err = service.Update()
		if err != nil {
			panic(err)
		}
	}
}
