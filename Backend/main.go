package main

import (
	"homelabmanager/config"
	"homelabmanager/web"
	"net/http"
)

func main() {
	// load config
	var config config.Config
	config.Load()

	// start web server
	web.CreateRoutes(config)
	http.ListenAndServe(":3333", nil)
}
