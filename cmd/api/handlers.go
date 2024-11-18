package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

// Possible improvement: use Redis cache to reduce lookup time
func getFileUrlById(app *Config, fileId string) (name string, e error) {
	log.Printf("Requested file id: %v\n", fileId)

	// simulate network delay
	time.Sleep(time.Second * 1)

	randomNumber := rand.IntN(len(app.fileUrls))
	fileUrl := app.fileUrls[randomNumber]

	log.Printf("using %v, from index %v", fileUrl, randomNumber)

	return fileUrl, nil
}

func (app *Config) FileDownloadStream(w http.ResponseWriter, r *http.Request) {
	fileId := r.PathValue("fileId")

	url, err := getFileUrlById(app, fileId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("File not found"))
		return
	}

	resp, err := http.Get(url)

	if err != nil {
		log.Println("HTTP Get error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not load file due to error"))
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Reading response body error:", err)
		}

		log.Println("HTTP Response Body:", string(body))

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not load file due to status code"))
		return
	}

	log.Println("Sending file")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+app.fileName)

	// See: https://itnext.io/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-ebe1b93fb72f
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
