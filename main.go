package main

import (
	"flag"
	"sw-go-template-server/modules/application"
	"sw-go-template-server/modules/webserver"
)

func main() {
	// Define flags
	httpAddress := flag.String("http", "example.com:80", "HTTP address")
	httpsAddress := flag.String("https", "example.com:443", "HTTPS address")
	webTlsCert := flag.String("tlscert", "/etc/letsencrypt/live/example.com/fullchain.pem", "Web TLS certificate")
	webTlsKey := flag.String("tlskey", "/etc/letsencrypt/live/example.com/privkey.pem", "Web TLS key")
	webDirectory := flag.String("webdir", "/frontend/", "Web directory")

	// Parse the flags
	flag.Parse()

	// Initialize the application
	app := application.Init()

	// Set configuration using flags
	app.Config.HttpAddress = *httpAddress
	app.Config.HttpsAddress = *httpsAddress
	app.Config.WebTlsCert = *webTlsCert
	app.Config.WebTlsKey = *webTlsKey
	app.Config.WebDirectory = *webDirectory

	// Initialize webserver
	webserver.Init(app)
}
