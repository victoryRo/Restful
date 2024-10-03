package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

// filterContentType middleware
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			_, _ = w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// setServerTimeCookie middleware
func setServerTimeCookie(handler http.Handler) http.Handler {
	cookie := http.Cookie{
		Name:  "ServerTimeUTC",
		Value: strconv.FormatInt(time.Now().Unix(), 10),
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
		handler.ServeHTTP(w, r)
	})
}

// handle Handler request
func handle(w http.ResponseWriter, r *http.Request) {
	var tempCity city

	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		fmt.Printf("Got %s city of area %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}

// main tiene una ligera variación en el mapeo de una ruta hasta el controlador.
// Utiliza llamadas a funciones anidadas para encadenar middleware, como se puede ver aquí:
func main() {
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/city", filterContentType(setServerTimeCookie(originalHandler)))

	fmt.Println("Run server port: 3011")
	log.Fatal(http.ListenAndServe(":3011", nil))
}

// Request
// curl -H "Content-Type: application/json" -X POST http://localhost:3011/city -d '{"name":"New York", "area":304}'
