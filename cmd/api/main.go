package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	fileUrl  string
	fileName string
}

func main() {
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fileUrl := os.Getenv("FILE_URL")
	if fileUrl == "" {
		log.Panic("FILE_URL not provided")
	}
	fileName := os.Getenv("FILE_NAME")
	if fileUrl == "" {
		log.Panic("FILE_NAME not provided")
	}

	app := Config{
		fileUrl:  fileUrl,
		fileName: fileName,
	}

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
