package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("ac_rest_go started.")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}