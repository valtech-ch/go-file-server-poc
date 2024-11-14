package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const fileUrl string = "<file-url>"

func getFileUrlById(fileId string) (name string, e error) {
	log.Printf("Requested file id: %v\n", fileId)

	// simulate network delay
	time.Sleep(time.Second * 1)

	return fileUrl, nil
}

func (app *Config) FileDownloadStream(w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("fileId")

	url, err := getFileUrlById(fileId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("File not found"))
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not load file"))
		return
	} else {
		log.Println("Sending file")
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=sample-pdf-download-10-mb.pdf")
	io.Copy(w, resp.Body)
}

func (app *Config) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(map[string]bool{"pong": true})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
		return
	}

	w.Write(jsonResp)
}
