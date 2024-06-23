package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s\n", queryParams["category"][0])
}

func main() {
	timeOut := 15 * time.Second
	router := mux.NewRouter()

	router.HandleFunc("/articles", QueryHandler)
	router.Queries("category", "id")

	serve := &http.Server{
		Handler:      router,
		Addr:         "localhost:3001",
		WriteTimeout: timeOut,
		ReadTimeout:  timeOut,
	}

	fmt.Println("Run server on port 3001")
	log.Fatal(serve.ListenAndServe())
}
