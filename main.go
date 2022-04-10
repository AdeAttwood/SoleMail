package main

import (
	"embed"
	"log"
	"os"

	"github.com/AdeAttwood/SoleMail/pkg/cli"
	"github.com/AdeAttwood/SoleMail/pkg/wails"
)

//go:embed build/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	if len(os.Args) > 1 && os.Args[1] == "import" {
		cli.Import()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "update" {
		cli.Update()
		return
	}

	err := wails.Run(assets, icon)

	if err != nil {
		log.Fatal(err)
	}
}
