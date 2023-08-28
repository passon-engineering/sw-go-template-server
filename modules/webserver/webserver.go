package webserver

import (
	"net/http"
	"os"
	"sw-go-template-server/modules/application"
	"time"

	"github.com/gorilla/mux"
	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func Init(app *application.Application) {
	router := mux.NewRouter()

	server := http.Server{
		Handler:      router,
		Addr:         app.Config.HttpsAddress,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	router.NotFoundHandler = http.HandlerFunc(handleRoot(app))

	handleRedirectToHttps(app)

	err := server.ListenAndServeTLS(app.Config.WebTlsCert, app.Config.WebTlsKey)
	if err != nil {
		app.Logger.Entry(logger.Container{
			Status: logger.STATUS_ERROR,
			Error:  "Could not initialize http server: " + err.Error(),
		})
		time.Sleep(2 * time.Second)
		os.Exit(1)
	}

}
