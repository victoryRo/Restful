package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverErr := dbDriver.Prepare(train)
	if driverErr != nil {
		log.Println(driverErr)
	}

	// Create train table
	_, statementErr := statement.Exec()
	if statementErr != nil {
		log.Println("Table already exists!")
	}

	statement, _ = dbDriver.Prepare(station)
	statement.Exec()
	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()

	log.Println("All tables created/initialized successfully!")
}
