package gocli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func Run(command string) (error) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return errors.New(fmt.Sprintf("Command not worked: %s", command))
	}

	return nil
}
