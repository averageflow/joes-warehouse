package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/averageflow/joes-warehouse/internal/infrastructure/app"
)

func main() {
	applicationServer := app.NewApplicationServer(nil)

	go func() {
		// we initialize the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		if err := applicationServer.State.HTTPServer.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Println("Gracefully closing server..")
			} else {
				log.Println(err.Error())
			}
		}
	}()

	log.Println("joe's warehouse application server started listening on port 7000")
	app.TerminationSignalWatcher(applicationServer.State.HTTPServer)
}
