package main

import (
	"ac_rest_go/in"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

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

var currentStatus map[string]in.State

var apiUrl string
var apiToken string
var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	apiUrl = os.Getenv("API_URL")
	apiToken = os.Getenv("API_TOKEN")

	fmt.Printf("ac_rest_go started with API URL %s\n", apiUrl)
	currentStatus = make(map[string]in.State)
	router := gin.Default()
	router.GET("/device/:device/status", status)
	router.POST("/device/:device/power", power)
	router.POST("/device/:device/ac_mode", ac_mode)
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "alive")
	})

	router.Run(":8080")
}

//GET /status
func status(c *gin.Context) {
	device := c.Param("device")
	fmt.Printf("/status for device %s\n", device)
	state := getCurrentStatus(device)
	c.JSON(http.StatusOK, state)
}

//POST /power
func power(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/power for device %s: %s\n", device, request.Value)
	state := getCurrentStatus(device)

	c.JSON(http.StatusOK, state)
}

//POST /ac_mode
func ac_mode(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/ac_mode for device %s: %s\n", device, request.Value)
	c.JSON(http.StatusOK, gin.H{"status": "power"})
}

func getCurrentStatus(device string) in.State {
	if deviceStatus, found := currentStatus[device]; found {
		return deviceStatus
	} else {
		fmt.Printf("No current state for %s, updating from the cloud.\n", device)
		return updateStatusFromCloud(device)
	}
}

/**
Updates the device state from cloud and keeps it in cache
*/
func updateStatusFromCloud(device string) in.State {
	currentStatus[device] = *getStateFromCloud(device)
	return currentStatus[device]
}

/**
gets the device state from Samsung cloud API
*/
func getStateFromCloud(device string) *in.State {
	req, err := http.NewRequest("GET", apiUrl+device, nil)
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("No response from request %s\n", apiUrl+device)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	samsungResponse := new(in.State)
	if err := json.Unmarshal(body, &samsungResponse); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := body[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
	}
	fmt.Printf("Device state updated from cloud: %s (%s, %s)\n", device, samsungResponse.DeviceTypeName, samsungResponse.Label)
	return samsungResponse
}
