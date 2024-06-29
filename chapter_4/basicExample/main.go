package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	// Create a web service
	webservice := new(restful.WebService)
	// Create a route and attach it to handler in the service
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// Add the service to application
	restful.Add(webservice)

	fmt.Println("Server run port :3111")
	log.Fatal(http.ListenAndServe(":3111", nil))
}

func pingTime(req *restful.Request, resp *restful.Response) {
	// Write to the response
	_, _ = io.WriteString(resp, fmt.Sprintf("%s", time.Now().Format(time.UnixDate)))
}
