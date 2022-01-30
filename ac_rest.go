package main

import (
	"ac_rest_go/in"
	"ac_rest_go/out"
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

var currentStatus map[string]out.State

var apiUrl string
var apiToken string
var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	apiUrl = os.Getenv("API_URL")
	apiToken = os.Getenv("API_TOKEN")

	fmt.Printf("ac_rest_go started with API URL %s\n", apiUrl)
	currentStatus = make(map[string]out.State)
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

func getCurrentStatus(device string) out.State {
	if deviceStatus, found := currentStatus[device]; found {
		return deviceStatus
	} else {
		fmt.Printf("No current state for %s, updating from the cloud.\n", device)
		return updateStatusFromCloud(device)
	}
}

func getJson(url string, target interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("No response from request %s/n", url)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, target); err != nil {
		fmt.Printf("Can not unmarshal JSON from %s/n", url)
	}
}

func updateStatusFromCloud(device string) out.State {
	samsungResponse := new(out.State)
	getJson(apiUrl+"/"+device, *samsungResponse)
	currentStatus[device] = *samsungResponse
	return *samsungResponse
}
