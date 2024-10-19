package main

import (
	"fmt"
	"log"

	"github.com/victoryRo/Restful/chapter_7/basicExample/helper"
)

func main() {
	_, err := helper.InitDB()
	if err != nil {
		fmt.Println(err)
	} else {
		log.Println("Database tables are successfully initialized.")
	}
}
