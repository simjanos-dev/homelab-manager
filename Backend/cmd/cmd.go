package cmd

import (
	"errors"
	"os/exec"
	"strings"
)

func GetCommandResultByLines(removeFirstLine bool, command string, arguments ...string) ([]string, error) {
	var cmd = exec.Command(command, arguments...)

	var cmdOutput, cmdError = cmd.Output()

	if cmdError != nil {
		return nil, errors.New("command execution error")
	}

	trimmedCmdOutput := strings.ReplaceAll(string(cmdOutput), "\r\n", "\n")
	lines := strings.Split(trimmedCmdOutput, "\n")

	if removeFirstLine {
		lines = lines[1:]
	}

	return lines, nil
}
