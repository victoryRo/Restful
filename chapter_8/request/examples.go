package main

import (
	"log"

	"github.com/levigross/grequests"
)

// Basic example
// func main() {
// 	resp, err := grequests.Get("http://httpbin.org/get", nil)
// 	if err != nil {
// 		log.Fatalln("Unable to make request: ", err)
// 	}
//
// 	log.Println(resp.String())
// }

// Json example
func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	var returnData map[string]interface{}
	err = resp.JSON(&returnData)
	if err != nil {
		log.Fatalln("JSON response with err: ", err)
	}
	log.Println(returnData)
}
