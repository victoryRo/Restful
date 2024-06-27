package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	_, _ = w.Write([]byte("OK"))
	log.Println("Finished processing request")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handle)
	loggerRouter := handlers.LoggingHandler(os.Stdout, router)

	fmt.Println("Server port :3001")
	log.Fatal(http.ListenAndServe(":3001", loggerRouter))
}
