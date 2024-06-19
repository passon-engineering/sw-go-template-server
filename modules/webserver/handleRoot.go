package webserver

import (
	"net/http"
	"os"
	"path/filepath"
	"sw-go-template-server/modules/application"
	"time"

	"github.com/tpasson/sw-go-logger-lib/logger"
)

func handleRoot(app *application.Application) http.HandlerFunc {
	basePath := filepath.Join(app.ServerPath, app.Config.WebDirectory)

	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		path := r.URL.Path[1:]
		if path == "" {
			path = "index.html"
		}

		fullPath := filepath.Join(basePath, path)

		// Check if the requested file exists and is not a directory
		info, err := os.Stat(fullPath)
		if os.IsNotExist(err) || err != nil || info.IsDir() {
			// Log the error if it's not a simple case of the file not existing
			if err != nil && !os.IsNotExist(err) {
				app.Logger.Entry(logger.Container{
					Status:         logger.STATUS_ERROR,
					Source:         "handleRoot",
					Info:           "error checking file: " + err.Error(),
					HttpRequest:    r,
					ProcessingTime: time.Since(startTime),
				})
			}

			// Serve the default file if the requested file does not exist or any error occurs
			http.ServeFile(w, r, filepath.Join(basePath, "index.html"))
		} else {
			// Serve the requested file
			http.ServeFile(w, r, fullPath)
		}

		// Log successful file serve
		app.Logger.Entry(logger.Container{
			Status:         logger.STATUS_INFO,
			Source:         "handleRoot",
			Info:           "served " + path,
			HttpRequest:    r,
			ProcessingTime: time.Since(startTime),
		})
	}
}
