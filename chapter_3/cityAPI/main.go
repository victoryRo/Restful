package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type city struct {
	Name string `json:"name"`
	Area uint64 `json:"area"`
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	var tempCity city

	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&tempCity)
		if err != nil {
			panic(err)
		}

		defer r.Body.Close()
		fmt.Printf("Got %v city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/city", postHandler)
	fmt.Println("Run server on port :3001")
	log.Fatal(http.ListenAndServe(":3001", nil))
}

// Request
// curl -H "Content-Type: application/json" -X POST http://localhost:3001/city -d '{"name":"Miami", "area":522}'
