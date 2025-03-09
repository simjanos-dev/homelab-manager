package system_information

import "servermanager/config"

func GetSystemInformation(conf config.Config) (SystemInformation, error) {
	systemInformation := SystemInformation{}
	systemInformation.Drives, _ = GetDrives(conf)
	systemInformation.Uptime, _ = GetUptime()
	systemInformation.Memory, _ = GetMemory()

	return systemInformation, nil
}
