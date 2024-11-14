package main

import (
	"net/http"

	"github.com/rs/cors"
)

func (app *Config) routes() http.Handler {
	mux := http.NewServeMux()

	// health check
	mux.HandleFunc("GET /ping", app.Ping)

	// File logic
	mux.HandleFunc("GET /download/{fileId}", app.FileDownloadStream)

	handler := cors.AllowAll().Handler(mux)

	return handler
}
