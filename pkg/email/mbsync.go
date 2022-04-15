package email

import (
	"os"
	"os/exec"
	"strings"

	"github.com/AdeAttwood/SoleMail/pkg/config"
)

const CONFIG_TEMPLATE = `
IMAPStore {{ID}}-remote
SSLType IMAPS
Host {{HOST}}
User {{USER}}
Pass {{PASSWORD}}

MaildirStore {{ID}}-local
Path {{DIR}}/
Inbox {{DIR}}/INBOX
Subfolders Verbatim

Channel {{ID}}
Master :{{ID}}-remote:
Slave :{{ID}}-local:
Create Slave
Expunge Both
CopyArrivalDate yes
`

type MbSync struct {
	account *config.Account
}

func NewMbSync(account *config.Account) *MbSync {
	return &MbSync{account: account}
}

func (mbsync *MbSync) Sync() error {
	config := strings.ReplaceAll(CONFIG_TEMPLATE, "{{ID}}", mbsync.account.ID)
	config = strings.ReplaceAll(config, "{{DIR}}", mbsync.account.Directory)
	config = strings.ReplaceAll(config, "{{USER}}", mbsync.account.UserName)
	config = strings.ReplaceAll(config, "{{PASSWORD}}", mbsync.account.Password)
	config = strings.ReplaceAll(config, "{{HOST}}", mbsync.account.Host)

	config_path := "/tmp/solemail-mbsync-" + mbsync.account.ID
	err := os.WriteFile(config_path, []byte(config), 0600)
	if err != nil {
		return err
	}

	defer os.Remove(config_path)

	cmd := exec.Command("mbsync", "--config", config_path, "--quiet", "--all")
	return cmd.Run()
}
