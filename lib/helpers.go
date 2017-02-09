package nocaps

import (
	"os"
	"os/exec"
)

// interactive commands (ssh, vim)
func runInteractiveCommand(name string, arg ...string) error {

	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
