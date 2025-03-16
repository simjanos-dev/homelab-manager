package routes

import (
	"encoding/json"
	"homelabmanager/config"
	"homelabmanager/system_information"
	"io"
	"net/http"
)

func DockerRoute(config config.Config, responseWriter http.ResponseWriter, request *http.Request) {
	dockerInformation, _ := system_information.GetDockerComposeFiles()
	formattedJson, _ := json.MarshalIndent(dockerInformation, "", "\t")

	io.Writer.Write(responseWriter, formattedJson)
}
