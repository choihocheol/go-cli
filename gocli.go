package gocli

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func Run(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	var output bytes.Buffer
	cmd.Stdout = &output

	err := cmd.Run()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Command not worked: %s", command))
	}

	return output.String(), nil
}
