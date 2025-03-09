package system_information

import (
	"errors"
	"os/exec"
	"strconv"
	"strings"
)

func GetMemory() (Memory, error) {
	var cmd = exec.Command("free", "-m")
	var cmdOutput, cmdError = cmd.Output()

	if cmdError != nil {
		return Memory{}, errors.New("command execution error")
	}

	trimmedCmdOutput := strings.ReplaceAll(string(cmdOutput), "\r\n", "\n")
	lines := strings.Split(trimmedCmdOutput, "\n")

	for lineIndex, line := range lines {
		if lineIndex == 0 {
			continue
		}

		lineData := strings.Fields(line)

		if len(lineData) < 4 || lineData[0] != "Mem:" {
			continue
		}

		total, _ := strconv.Atoi(lineData[1])
		used, _ := strconv.Atoi(lineData[2])
		free, _ := strconv.Atoi(lineData[3])
		return Memory{
			Unit:  "MB",
			Total: total,
			Used:  used,
			Free:  free,
		}, nil
	}

	return Memory{}, errors.New("memory information not found")
}
