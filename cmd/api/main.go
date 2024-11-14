package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct{}

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	app := Config{}

	log.Printf("Starting file service on port %s\n", httpPort)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
