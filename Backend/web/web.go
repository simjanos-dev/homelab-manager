package web

import (
	"homelabmanager/config"
	"homelabmanager/web/routes"
	"log"
	"net/http"
	"time"
)

func CreateRoutes(config config.Config) {
	http.HandleFunc("/system-information", func(responseWriter http.ResponseWriter, request *http.Request) {
		start := time.Now()
		routes.SystemInformationRoute(config, responseWriter, request)
		LogRequest(request, time.Since(start))
	})

	http.HandleFunc("/docker", func(responseWriter http.ResponseWriter, request *http.Request) {
		start := time.Now()
		routes.DockerRoute(config, responseWriter, request)
		LogRequest(request, time.Since(start))
	})
}

func LogRequest(request *http.Request, duration time.Duration) {
	log.Printf("%s took %s", request.Pattern, duration)
}
