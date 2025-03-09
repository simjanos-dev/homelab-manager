package main

import (
	"encoding/json"
	"fmt"
	"servermanager/config"
	"servermanager/system_information"
)

func main() {
	var config config.Config
	config.Load()

	systemInformation, _ := system_information.GetSystemInformation(config)

	jcart, _ := json.MarshalIndent(systemInformation, "", "\t")
	fmt.Println(string(jcart))
}
