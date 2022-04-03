package config

import (
	"fmt"
	"os"
	"runtime"
)

func GetDefaultConfigPath() string {
	if runtime.GOOS == "linux" {
		return fmt.Sprintf("%s/.config/sole-mail", os.Getenv("HOME"))

	}

	panic("Un supperted OS " + runtime.GOOS)
}

func GetDefaultDataPath() string {
	if runtime.GOOS == "linux" {
		return fmt.Sprintf("%s/.local/share/sole-mail", os.Getenv("HOME"))

	}

	panic("Un supperted OS " + runtime.GOOS)
}
