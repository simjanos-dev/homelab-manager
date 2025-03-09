package system_information

import (
	"errors"
	"os/exec"
	"servermanager/config"
	"slices"
	"strconv"
	"strings"
)

func GetDrives(conf config.Config) ([]Drive, error) {
	var drives []Drive

	cmdParameters := slices.Concat([]string{"-B", "GB"}, conf.Drives)

	var cmd = exec.Command("df", cmdParameters...)
	var cmdOutput, cmdError = cmd.Output()

	if cmdError != nil {
		return nil, errors.New("command execution error")
	}

	trimmedCmdOutput := strings.ReplaceAll(string(cmdOutput), "\r\n", "\n")
	lines := strings.Split(trimmedCmdOutput, "\n")

	for lineIndex, line := range lines {
		if lineIndex == 0 {
			continue
		}

		lineData := strings.Fields(line)

		if len(lineData) < 4 {
			continue
		}

		total, _ := strconv.Atoi(strings.ReplaceAll(lineData[1], "GB", ""))
		used, _ := strconv.Atoi(strings.ReplaceAll(lineData[2], "GB", ""))
		free, _ := strconv.Atoi(strings.ReplaceAll(lineData[3], "GB", ""))
		drives = append(drives, Drive{
			Name:       lineData[0],
			Unit:       "GB",
			TotalSpace: total,
			UsedSpace:  used,
			FreeSpace:  free,
		})
	}

	return drives, nil
}
