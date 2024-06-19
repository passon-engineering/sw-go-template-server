package main

import (
	"flag"

	"github.com/tpasson/sw-go-template-server/modules/application"
	"github.com/tpasson/sw-go-template-server/modules/webserver"
)

func main() {
	// Define flags
	httpAddress := flag.String("http", "example.com:80", "HTTP address")
	httpsAddress := flag.String("https", "example.com:443", "HTTPS address")
	webTlsCert := flag.String("tlscert", "/etc/letsencrypt/live/example.com/fullchain.pem", "Web TLS certificate")
	webTlsKey := flag.String("tlskey", "/etc/letsencrypt/live/example.com/privkey.pem", "Web TLS key")
	webDirectory := flag.String("webdir", "/frontend/", "Web directory")
	logDirectory := flag.String("logdir", "", "Log directory")

	// Parse the flags
	flag.Parse()

	config := application.Config{}
	// Set configuration using flags
	config.HttpAddress = *httpAddress
	config.HttpsAddress = *httpsAddress
	config.WebTlsCert = *webTlsCert
	config.WebTlsKey = *webTlsKey
	config.WebDirectory = *webDirectory
	config.LogDirectory = *logDirectory

	// Initialize the application
	app := application.Init(config)

	// Initialize webserver
	webserver.Init(app)
}
