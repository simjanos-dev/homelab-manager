package system_information

import "servermanager/config"

type SystemInformation struct {
	Uptime           string
	Memory           Memory
	Drives           []Drive
	DockerContainers []DockerComposeFile
}

func GetSystemInformation(conf config.Config) (SystemInformation, error) {
	systemInformation := SystemInformation{}
	systemInformation.Drives, _ = GetDrives(conf)
	systemInformation.Uptime, _ = GetUptime()
	systemInformation.Memory, _ = GetMemory()
	systemInformation.DockerContainers, _ = GetDockerComposeFiles()

	return systemInformation, nil
}
