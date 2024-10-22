package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/victoryRo/Restful/chapter_7/jsonstore/helper"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

// PackageResponse is the response to be send back for Package
type PackageResponse struct {
	Package helper.Package `json:"Package"`
}

func (driver *DBClient) GetPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	vars := mux.Vars(r)

	driver.db.First(&Package, vars["id"])

	var PackageData any
	err := json.Unmarshal([]byte(Package.Data), &PackageData)
	if err != nil {
		log.Fatal(err)
	}

	var response = PackageResponse{Package: Package}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resJSON, _ := json.Marshal(response)
	_, _ = w.Write(resJSON)
}

func (driver *DBClient) GetPackagesbyWeight(w http.ResponseWriter, r *http.Request) {
	var packages []helper.Package
	weight := r.FormValue("weight")

	var query = "select * from \"Package\" where data->>'weight'=?"
	driver.db.Raw(query, weight).Scan(&packages)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(packages)
	_, _ = w.Write(respJSON)
}

func (driver *DBClient) PostPackage(w http.ResponseWriter, r *http.Request) {
	var Package = helper.Package{}
	postBody, _ := io.ReadAll(r.Body)

	Package.Data = string(postBody)
	driver.db.Save(&Package)

	responseMap := map[string]interface{}{"id": Package.ID}
	w.Header().Set("Content-Type", "application/json")
	response, _ := json.Marshal(responseMap)
	_, _ = w.Write(response)
}

func main() {
	db, err := helper.InitDB()
	if err != nil {
		panic(err)
	}

	dbclient := &DBClient{db: db}

	route := mux.NewRouter()
	route.HandleFunc("/v1/package/{id:[a-zA-Z0-9]*}", dbclient.GetPackage).Methods(http.MethodGet)
	route.HandleFunc("/v1/package", dbclient.PostPackage).Methods(http.MethodPost)
	route.HandleFunc("/v1/package", dbclient.GetPackagesbyWeight).Methods(http.MethodGet)

	times := 15 * time.Second

	srv := &http.Server{
		Handler:      route,
		Addr:         "127.0.0.1:3002",
		WriteTimeout: times,
		ReadTimeout:  times,
	}

	log.Println("Runnig server on port :3002")
	log.Fatal(srv.ListenAndServe())
}

// curl -X POST \
// http://localhost:3002/v1/package \
// -H 'cache-control: no-cache' \
// -H 'content-type: application/json' \
// -d '{
//     "dimensions": {
//     "width": 21,
//     "height": 12
//     },
//     "weight": 10,
//     "is_damaged": false,
//     "status": "In transit"
// }'
