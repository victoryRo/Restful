package helper

import (
	"fmt"
	"log"

	// _ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 54321
	user     = "victoJSON"
	password = "secretoJSONB"
	dbname   = "store"
)

// type Shipment struct {
// 	gorm.Model
// 	Packages []Package
// 	Data     string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
// }

type Package struct {
	gorm.Model
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM creates tables with plural names. Use this to suppress it
// func (Shipment) TableName() string {
// 	return "Shipment"
// }

func (Package) TableName() string {
	return "Package"
}

func InitDB() (*gorm.DB, error) {
	var err error
	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		log.Println("Database Connection ok")
	}

	err = db.AutoMigrate(&Package{})
	if err != nil {
		return nil, err
	} else {
		log.Println("Tables Package, Shipment Created.")
	}
	return db, nil
}
