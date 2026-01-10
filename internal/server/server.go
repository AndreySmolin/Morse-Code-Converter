package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Http   *http.Server
}

func NewRouter(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc(`/`, handlers.RootHandler)
	router.HandleFunc(`/upload`, handlers.LoadHandler)
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return &Server{
		Logger: logger,
		Http:   httpServer,
	}
}
