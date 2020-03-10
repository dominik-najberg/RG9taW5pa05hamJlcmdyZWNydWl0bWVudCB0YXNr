package server

import (
	"context"
	"fmt"
	"github.com/dominik-najberg/RG9taW5pa05hamJlcmdyZWNydWl0bWVudCB0YXNr/client"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func NewApiServer(port int) *http.Server {
	h := newHandlers(Cache()(client.HttpGetFetcher))
	http.HandleFunc("/weather/city/", h.weatherByCitiesViewHandler)

	return &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: nil,
	}
}

func Start(port int) {
	end := make(chan bool, 1)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	apiServer := NewApiServer(port)
	go GracefullyShutDownServer(apiServer, stop, end)
	log.Printf("server listening on %s", apiServer.Addr)

	if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("apiServer error: %v", err)
	}

	<-end
}

func GracefullyShutDownServer(server *http.Server, stopChannel chan os.Signal, thatsAll chan bool) {
	<-stopChannel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("error while shutting down the server: %v", err)
	}
	close(thatsAll)
}
