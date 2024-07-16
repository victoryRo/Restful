package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/victoryRo/Restful/chapter_4/dbutils"
)

// DB Driver visible to whole program
var DB *sql.DB

type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

type ScheduleResource struct {
	ID          int
	TrainID     int
	StatationID int
	ArrivalTime time.Time
}

// Register adds paths and routes to container
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))

	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")

	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)

	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)

	statement, _ := DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) values (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) removeTrain(req *restful.Request, rep *restful.Response) {
	id := req.PathParameter("train-id")

	statement, _ := DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err == nil {
		rep.WriteHeader(http.StatusOK)
	} else {
		rep.AddHeader("Content-Type", "text/plain")
		rep.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	// Connect to database
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	// create tables
	dbutils.Initialize(DB)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)

	log.Printf("start listening on localhost:3001")
	server := &http.Server{Addr: ":3001", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

/**
curl -X POST \
     http://localhost:3001/v1/trains \
    -H 'cache-control: no-cache' \
    -H 'content-type: application/json' \
    -d '{"driverName": "Veronica", "operatingStatus": true}'
*/

// curl -X GET "http://localhost:3001/v1/trains/1"

// curl -X DELETE "http://localhost:3001/v1/trains/1"
