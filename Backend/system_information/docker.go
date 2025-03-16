package system_information

import (
	"servermanager/cmd"
	"strings"
)

type DockerComposeFile struct {
	ComposeFile string
	Name        string
	Containers  []DockerContainer
}

type DockerContainer struct {
	Name   string
	Status string
}

func GetDockerComposeFiles() ([]DockerComposeFile, error) {
	var composeFiles []DockerComposeFile

	lines, cmdError := cmd.GetCommandResultByLines(true, "docker", "compose", "ls", "-a")

	if cmdError != nil {
		return nil, cmdError
	}

	for _, line := range lines {
		lineData := strings.Fields(line)

		if len(lineData) != 3 {
			continue
		}

		containers, _ := GetDockerContainers(lineData[2])

		composeFiles = append(composeFiles, DockerComposeFile{
			ComposeFile: lineData[2],
			Name:        lineData[0],
			Containers:  containers,
		})
	}

	return composeFiles, nil
}

func GetDockerContainers(composeFile string) ([]DockerContainer, error) {
	var containers []DockerContainer

	lines, cmdError := cmd.GetCommandResultByLines(false, "docker", "compose", "-f", composeFile, "ps", "-a", "--format", "'{{.ID}}|{{.Names}}|{{.State}}'")

	if cmdError != nil {
		return nil, cmdError
	}

	for _, line := range lines {
		lineData := strings.Split(line, "|")

		if len(lineData) != 3 {
			continue
		}

		containers = append(containers, DockerContainer{
			Name:   lineData[1],
			Status: lineData[2],
		})
	}

	return containers, nil
}
