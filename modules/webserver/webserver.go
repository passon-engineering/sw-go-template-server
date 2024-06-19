package webserver

import (
	"log"
	"net/http"
	"os"
	"sw-go-template-server/modules/application"
	"time"

	"github.com/gorilla/mux"
	"github.com/tpasson/sw-go-logger-lib/logger"
)

// customLogWriter redirects log messages to the application's custom logger.
type customLogWriter struct {
	app *application.Application
}

func (w *customLogWriter) Write(p []byte) (n int, err error) {
	// Here, you might want to parse the log message or directly log it as is.
	// This example logs the message directly as an informational entry.
	w.app.Logger.Entry(logger.Container{
		Status: logger.STATUS_INFO,
		Info:   string(p),
	})
	return len(p), nil
}

func Init(app *application.Application) {
	// Redirect standard log output to the custom logger.
	log.SetOutput(&customLogWriter{app: app})

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(handleRoot(app))

	server := http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	// Check if SSL certificates are provided
	if app.Config.WebTlsCert != "" && app.Config.WebTlsKey != "" {
		// HTTPS configuration
		server.Addr = app.Config.HttpsAddress
		handleRedirectToHttps(app) // Assuming this sets up HTTP to HTTPS redirection

		// Start HTTPS server
		err := server.ListenAndServeTLS(app.Config.WebTlsCert, app.Config.WebTlsKey)
		if err != nil {
			logFatal(app, "Could not initialize HTTPS server: "+err.Error())
		}
	} else {
		// HTTP configuration
		server.Addr = app.Config.HttpAddress

		// Start HTTP server
		err := server.ListenAndServe()
		if err != nil {
			logFatal(app, "Could not initialize HTTP server: "+err.Error())
		}
	}
}

func logFatal(app *application.Application, message string) {
	app.Logger.Entry(logger.Container{
		Status: logger.STATUS_ERROR,
		Error:  message,
	})
	time.Sleep(2 * time.Second)
	os.Exit(1)
}
