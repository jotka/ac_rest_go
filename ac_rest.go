package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type req struct {
	value string
}

type _switch struct {
	fanmode fanmode
	value   string
}

type fanmode struct {
	value float32
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
	device := c.Param("device")
	var reqBody req
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("/power for device %s: %s%n", device, reqBody.value)
	c.JSON(http.StatusOK, gin.H{"status": "power"})
}

func ac_mode(context *gin.Context) {

}
