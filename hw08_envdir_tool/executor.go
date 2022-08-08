package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, envValue := range env {
		if envValue.NeedRemove {
			os.Unsetenv(name)
		} else {
			os.Setenv(name, envValue.Value)
		}
	}
	var out bytes.Buffer
	c := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec

	c.Stdout = &out
	err := c.Run()
	if err != nil {
		fmt.Println(err)
		errExitCode := &exec.ExitError{}
		if errors.As(err, &errExitCode) {
			return errExitCode.ExitCode()
		}
		return 1
	}
	fmt.Println(out.String())
	return
}
