package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/victoryRo/Restful/chapter_4/dbutils"
)

var DB *sql.DB

type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

func GetStation(c *gin.Context) {
	var station StationResource
	id := c.Param("station_id")

	sqlQuery := "select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?"
	query := DB.QueryRow(sqlQuery, id)
	err := query.Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource

	if err := c.BindJSON(&station); err == nil {
		statement, _ := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newID, _ := result.LastInsertId()
			station.ID = int(newID)
			c.JSON(http.StatusOK, gin.H{
				"result": station,
			})
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station_id")
	statement, _ := DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./station.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(DB)

	// ---------------------------

	router := gin.Default()

	router.GET("/v1/stations/:station_id", GetStation)
	router.POST("/v1/stations", CreateStation)
	router.DELETE("/v1/stations/:station_id", RemoveStation)

	err = router.Run(":4000")
	if err != nil {
		log.Println("Run err listen and sever", err)
	}
}

/**
curl -X POST \
  http://localhost:4000/v1/stations \
 -H 'cache-control: no-cache' \
 -H 'content-type: application/json' \
 -d '{"name":"Brooklyn", "opening_time":"8:12:00", "closing_time":"18:23:00"}'
*/

// curl -X GET "http://localhost:4000/v1/stations/1"

// curl -X DELETE "http://localhost:4000/v1/stations/1"
