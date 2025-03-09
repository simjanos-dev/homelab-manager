package system_information

type Drive struct {
	Name       string
	Unit       string
	TotalSpace int
	UsedSpace  int
	FreeSpace  int
}
