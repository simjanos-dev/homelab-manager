package system_information

import (
	"errors"
	"servermanager/cmd"
	"strconv"
	"strings"
)

type Memory struct {
	Unit  string
	Total int
	Used  int
	Free  int
}

func GetMemory() (Memory, error) {
	lines, cmdError := cmd.GetCommandResultByLines(true, "free", "-m")

	if cmdError != nil {
		return Memory{}, cmdError
	}

	for _, line := range lines {
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
