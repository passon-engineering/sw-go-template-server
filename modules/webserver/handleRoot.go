package webserver

import (
	"net/http"
	"os"
	"path/filepath"
	"sw-go-template-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleRoot(app *application.Application) http.HandlerFunc {
	// Precompute the base path
	basePath := filepath.Join(app.ServerPath, app.Config.WebDirectory)

	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		path := r.URL.Path[1:]
		if path == "" {
			path = "index.html"
		}

		fullPath := filepath.Join(basePath, path)

		_, err := os.Stat(fullPath)
		if err != nil && os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(basePath, "index.html"))

			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Source:         "handleRoot",
				Info:           "served",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})
			return
		}

		http.ServeFile(w, r, fullPath)

		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Source:         "handleRoot",
			Info:           "served",
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
