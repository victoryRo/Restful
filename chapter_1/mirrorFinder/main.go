package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/victoryRo/Restful/chapter_1/mirrors"
)

type response struct {
	FastestURL string        `json:"fastest_url"`
	Latency    time.Duration `json:"latency"`
}

func findFastest(urls []string) response {
	urlCh := make(chan string)
	latencyCh := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			log.Println("Started probing: ", mirrorURL)
			start := time.Now()

			_, err := http.Get(mirrorURL + "/README")
			latency := time.Since(start) / time.Millisecond
			if err == nil {
				urlCh <- mirrorURL
				latencyCh <- latency
			}
			log.Printf("Got the best mirror: %s with latency: %s\n", mirrorURL, latency)
		}()
	}
	return response{<-urlCh, <-latencyCh}
}

func main() {
	http.HandleFunc("/fastest-mirror", callMirror)
	port := ":8080"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func callMirror(w http.ResponseWriter, r *http.Request) {
	response := findFastest(mirrors.MirrorList)
	respJSON, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(respJSON)
}
