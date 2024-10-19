package helper

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 54320
	user     = "victory"
	password = "secretoSafe"
	dbname   = "urldb"
)

func InitDB() (*sql.DB, error) {
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// dbURI := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable" ,user, password, host, port, dbname)

	var err error

	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}

	// verificar conexion
	if err := db.Ping(); err != nil {
		log.Fatalln("Error from database ping:", err)
	} else {
		log.Println("Database connection successful")
	}

	table := "web_url(ID SERIAL PRIMARY KEY, url TEXT NOT NULL);"

	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS " + table)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, err
	}

	return db, nil
}
