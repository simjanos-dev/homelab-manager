package system_information

import (
	"errors"
	"os/exec"
	"strings"
)

func GetUptime() (string, error) {
	var cmd = exec.Command("uptime")
	var cmdOutput, cmdError = cmd.Output()

	if cmdError != nil {
		return "", errors.New("command execution error")
	}

	uptime := strings.TrimSpace(string(cmdOutput))
	uptime = strings.Fields(uptime)[0]
	return uptime, nil
}
