package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type city struct {
	Name string
	Area uint64
}

// filterContentType Middleware
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently check content type Middleware")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, _ = w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// setServerTimeCookie Middleware
func setServerTimeCookie(handler http.Handler) http.Handler {
	cookie := http.Cookie{
		Name:  "Server-Time(UTC)",
		Value: strconv.FormatInt(time.Now().Unix(), 10),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		http.SetCookie(w, &cookie)
		log.Println("Currently server time Middleware")
	})
}

// handle Handler
func handle(w http.ResponseWriter, r *http.Request) {
	var tempCity city

	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		log.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 Method Not Allowed"))
	}
}

func main() {
	originalHandler := http.HandlerFunc(handle)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(originalHandler)
	http.Handle("/city", chain)

	fmt.Println("Server port :3007")
	log.Fatal(http.ListenAndServe(":3007", nil))
}
