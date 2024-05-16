package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World from Go Gin!")
	})

	router.GET("/admin", func(c *gin.Context) {
		c.String(200, "Admin area")
	})

	router.Run(":5000") // listen and serve on 0.0.0.0:5000
}
