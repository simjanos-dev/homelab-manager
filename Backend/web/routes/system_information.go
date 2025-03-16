package routes

import (
	"encoding/json"
	"homelabmanager/config"
	"homelabmanager/system_information"
	"io"
	"net/http"
)

func SystemInformationRoute(config config.Config, responseWriter http.ResponseWriter, request *http.Request) {
	systemInformation, _ := system_information.GetSystemInformation(config)
	formattedJson, _ := json.MarshalIndent(systemInformation, "", "\t")

	io.Writer.Write(responseWriter, formattedJson)
}
