package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB stores the database session imformation. Needs to be initialized once
type DB struct {
	collection *mongo.Collection
}

type Movie struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	Name      string      `json:"name" bson:"name"`
	Year      string      `json:"year" bson:"year"`
	Directors []string    `json:"directors" bson:"directors"`
	Writers   []string    `json:"writers" bson:"writers"`
	BoxOffice BoxOffice   `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

// GetMovie fetches a movie with a given ID
func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}
	err := db.collection.FindOne(context.TODO(), filter).Decode(&movie)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}
}

// PostMovie adds a new movie to our MongoDB collection
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	postBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(postBody, &movie)

	result, err := db.collection.InsertOne(context.TODO(), movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(response)
	}
}

// UpdateMovie modifies the data of given resource
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	putBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(putBody, &movie)

	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": &movie}
	_, err := db.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("Updated succesfully!"))
	}
}

// DeleteMovie removes the data from the db
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	objectID, _ := primitive.ObjectIDFromHex(vars["id"])
	filter := bson.M{"_id": objectID}

	_, err := db.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte("Deleted succesfully!"))
	}
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB successfully")

	collection := client.Database("appDB").Collection("movies")
	db := &DB{collection: collection}

	// --------------------------------------------------------

	route := mux.NewRouter()
	route.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.GetMovie).Methods("GET")
	route.HandleFunc("/v1/movies", db.PostMovie).Methods("POST")
	route.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.UpdateMovie).Methods("PUT")
	route.HandleFunc("/v1/movies/{id:[a-zA-Z0-9]*}", db.DeleteMovie).Methods("DELETE")

	srv := &http.Server{
		Handler:      route,
		Addr:         "127.0.0.1:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Run server port :8001")
	log.Fatal(srv.ListenAndServe())
}

/**
curl -X POST \
 http://localhost:8001/v1/movies \
 -H 'cache-control: no-cache' \
 -H 'content-type: application/json' \
 -d '{ "name" : "Avatar", "year" : "2014", "directors" : [ "Cameron Diaz" ], "writers" : [ "Jonathan Diaz", "Martin Mar" ], "boxOffice" : { "budget" : 195000000, "gross" : 933316061 }
}'
*/

// curl -X GET http://localhost:8000/v1/movies/5cfd6cf0c281945c6cfefaab
