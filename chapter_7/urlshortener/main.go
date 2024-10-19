package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/victoryRo/Restful/chapter_7/urlshortener/helper"
	base62 "github.com/victoryRo/Restful/chapter_7/urlshortener/utils"
)

// DBClient stores the database session information. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

// Record Model is a HTTP response
type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

// GetOriginalURL fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)

	// Get ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])

	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)

	// Handle response details
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		_, _ = w.Write(response)
	}
}

// GenerateShortURL adds URL to DB and gives back shortened string
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(postBody, &record)
	if err != nil {
		log.Fatal(err)
	}

	err = driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.URL).Scan(&id)
	responseMap := map[string]string{"encoded_string": base62.ToBase62(id)}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		_, _ = w.Write(response)
	}
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}

	dbclient := &DBClient{db: db}
	defer db.Close()

	// create a new router
	router := mux.NewRouter()

	// attach path with handler
	router.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	router.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")

	// Good practice: enforce timeouts for servers you create!
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Run server on port: 8080")
	log.Fatal(srv.ListenAndServe())
}

// curl -X POST \
// http://localhost:8080/v1/short \
// -H 'cache-control: no-cache' \
// -H 'content-type: application/json' \
// -d '{
//   "url": "https://www.packtpub.com/eu/game-development/unreal-engine-4-shaders-and-effects-cookbook"
// }'

// https://subscription.packtpub.com/book/web-development/9781838643577/7/ch07lvl1sec57/implementing-a-url-shortening-service-using-postgresql-and-pq
// https://github.com/PacktPublishing/Hands-On-Restful-Web-services-with-Go/blob/master/chapter7/urlshortener/main_test.go
// https://pixabay.com/sound-effects/
