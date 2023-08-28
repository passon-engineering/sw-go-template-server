package webserver

import (
	"net/http"
	"sw-go-template-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleRoot(app *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		path := r.URL.Path[1:]
		if path == "" {
			path = "index.html"
		}

		http.ServeFile(w, r, app.ServerPath+app.Config.WebDirectory+path)
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Source:         "handleRoot",
			Info:           "served static file",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
