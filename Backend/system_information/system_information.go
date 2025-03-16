package system_information

import "homelabmanager/config"

type SystemInformation struct {
	Uptime string
	Memory Memory
	Drives []Drive
}

func GetSystemInformation(conf config.Config) (SystemInformation, error) {
	systemInformation := SystemInformation{}
	systemInformation.Drives, _ = GetDrives(conf)
	systemInformation.Uptime, _ = GetUptime()
	systemInformation.Memory, _ = GetMemory()

	return systemInformation, nil
}
