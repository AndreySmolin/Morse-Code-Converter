package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server:", log.LstdFlags)
	server := server.NewRouter(logger)
	logger.Print("Start port:8080")
	err := server.Http.ListenAndServe()
	if err != nil {
		logger.Fatalf("Fatal error:%v", err)
	}
}
