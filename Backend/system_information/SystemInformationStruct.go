package system_information

type SystemInformation struct {
	Uptime string
	Memory Memory
	Drives []Drive
}
