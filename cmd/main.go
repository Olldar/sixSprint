package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	//"sixSprint/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Server starting on :8080")

	if err := srv.HTTP.ListenAndServe(); err != nil {
		logger.Fatalf("server failed: %v", err)
	}

}
