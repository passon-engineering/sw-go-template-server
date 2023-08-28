package webserver

import (
	"log"
	"net/http"
	"sw-go-template-server/modules/application"

	"github.com/passon-engineering/sw-go-logger-lib/logger"
)

func handleRedirectToHttps(app *application.Application) {
	redirectServer := &http.Server{
		Addr: app.Config.HttpAddress,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			app.Logger.Entry(logger.Container{
				Status:      logger.STATUS_INFO,
				Source:      "redirectToHTTPS",
				HttpRequest: r,
			})

			u := "https://" + app.Config.HttpsAddress + r.RequestURI
			http.Redirect(w, r, u, http.StatusMovedPermanently)
		}),
	}

	go func() {
		if err := redirectServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error initializing redirect server: %s", err)
			return
		}
	}()

}
