package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	/* GET takes a route and a handler function
	   Handler takes the gin context object
	*/
	r.GET("/pingTime", controller)

	r.Run(":3007")
}

func controller(c *gin.Context) {
	c.JSON(200, gin.H{
		"serverTime": time.Now().UTC(),
	})
}
