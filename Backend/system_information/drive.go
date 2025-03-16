package system_information

import (
	"servermanager/cmd"
	"servermanager/config"
	"slices"
	"strconv"
	"strings"
)

type Drive struct {
	Name       string
	Unit       string
	TotalSpace int
	UsedSpace  int
	FreeSpace  int
}

func GetDrives(conf config.Config) ([]Drive, error) {
	var drives []Drive

	cmdParameters := slices.Concat([]string{"-B", "GB"}, conf.Drives)

	lines, cmdError := cmd.GetCommandResultByLines(true, "df", cmdParameters...)

	if cmdError != nil {
		return nil, cmdError
	}

	for _, line := range lines {
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
