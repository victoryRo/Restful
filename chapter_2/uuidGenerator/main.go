package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
)

// UUID es un multiplexor personalizado
type UUID struct{}

func (p *UUID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandomUUID(w, r)
		return
	}
	http.NotFound(w, r)
}

func giveRandomUUID(w http.ResponseWriter, r *http.Request) {
	c := 10
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%x", b)
}

func main() {
	mux := &UUID{}
	fmt.Println("Run server port :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
