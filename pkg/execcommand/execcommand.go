package execcommand

import (
	"errors"
	"fmt"
	"os/exec"
)

const ShellToUse = "/bin/sh"

func Shellout(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(fmt.Sprint(err) + ": " + string(output))
	}
	return nil
}
