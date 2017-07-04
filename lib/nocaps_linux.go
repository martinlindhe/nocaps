package nocaps

// NOTE: due to insane fragmentation of the Linux desktop,
// there dont's seem to be a single solution for executing a command on X start-up

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func Perform() error {

	content := `[Desktop Entry]
Type=Application
Name=nocaps
Comment=github.com/martinlindhe/nocaps
Exec=/usr/bin/setxkbmap -option caps:none
Terminal=false
NoDisplay=true
`

	usr, err := user.Current()
	if err != nil {
		return err
	}

	fileName := filepath.Join(usr.HomeDir, ".config/autostart/nocaps.desktop")
	parent := filepath.Dir(fileName)
	err = os.MkdirAll(parent, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, []byte(content), 0755)
	if err != nil {
		return err
	}
	fmt.Println("Generated", fileName, "which will be run each start-up")

	return runInteractiveCommand("setxkbmap", "-option", "caps:none")
}
