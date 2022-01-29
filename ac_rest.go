package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	Value string `form:"value" json:"value" xml:"value"  binding:"required"`
}

type _switch struct {
	fanmode fanmode
	value   string
}

type fanmode struct {
	value float32
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	fmt.Println("ac_rest_go started.")

	router := gin.Default()
	router.POST("/device/:device/power", power)
	router.POST("/device/:device/ac_mode", ac_mode)
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "alive")
	})

	router.Run(":8080")
}

func power(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/power for device %s: %s\n", device, request.Value)
	c.JSON(http.StatusOK, gin.H{"status": "power"})
}

func ac_mode(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/ac_mode for device %s: %s\n", device, request.Value)
	c.JSON(http.StatusOK, gin.H{"status": "power"})
}
