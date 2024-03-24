package webserver

import (
	"net/http"
	"sw-go-template-server/modules/application"
	"time"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleRedirectToHttps(app *application.Application) {
	redirectServer := &http.Server{
		Addr: app.Config.HttpAddress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			u := "https://" + app.Config.HttpsAddress + r.RequestURI
			http.Redirect(w, r, u, http.StatusMovedPermanently)

			app.Logger.Entry(logger.Container{
				Status:         logger.STATUS_INFO,
				Source:         "handleRedirectToHttps",
				HttpRequest:    r,
				ProcessingTime: time.Since(startTime),
			})
		}),
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	go func() {
		if err := redirectServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Entry(logger.Container{
				Status: logger.STATUS_ERROR,
				Source: "handleRedirectToHttps",
				Info:   "could not redirect: " + err.Error(),
			})
			return
		}
	}()
}
