package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - a representation of the http handler
type Handler struct {
	Router  *mux.Router
	Service MailerService
	Server  *http.Server
}

// NewHandler - returns a new handler with the router initialized and the
// service configured
func NewHandler(service MailerService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.Router.Use(CorsMiddleware)
	h.Router.Use(JSONMiddleware)
	h.Router.Use(LoggingMiddleware)
	h.Router.Use(TimeoutMiddleware)
	h.Router.Use(ApiKeyMiddleware)
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/alive", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "I am alive")
	})

	h.Router.HandleFunc("/api/v1/sendmail", h.SendEmail).Methods(http.MethodPost, http.MethodOptions)
}

// Serve - starts the server to listen for requests handles gracefully shut downs
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := h.Server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("shut down gracefully")
	return nil
}
