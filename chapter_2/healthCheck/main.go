package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// HealthCheck controller "handler"
// returns date time to client
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	_, _ = io.WriteString(w, currentTime.String())
}

func main() {
	http.HandleFunc("/health", HealthCheck)

	fmt.Println("Running server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
