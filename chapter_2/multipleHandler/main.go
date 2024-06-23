package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/random-float", float)
	newMux.HandleFunc("/random-int", integer)

	fmt.Println("Run server in port :3000")
	log.Fatal(http.ListenAndServe(":3000", newMux))
}

func float(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Float64())
}

func integer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, rand.Intn(100))
}
