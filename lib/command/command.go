package command

import (
	"os"
	"os/exec"

	apperrors "github.com/wakuwaku3/example-golang-cobra/lib/app_errors"
)

func Execute(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		return apperrors.Wrap(err)
	}

	return apperrors.Wrap(cmd.Wait())
}
