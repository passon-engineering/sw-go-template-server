package main

import (
	"sw-go-template-server/modules/application"
	"sw-go-template-server/modules/webserver"
)

func main() {
	app := application.Init()

	app.Config.HttpAddress = "example.com:80"
	app.Config.HttpsAddress = "example.com:443"
	app.Config.WebTlsCert = "/etc/letsencrypt/live/example.com/fullchain.pem"
	app.Config.WebTlsKey = "/etc/letsencrypt/live/example.com/privkey.pem"
	app.Config.WebDirectory = "/frontend/"
	webserver.Init(app)
}
