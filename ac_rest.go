package main

import (
	"ac_rest_go/in"
	"ac_rest_go/out"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var currentStatus map[string]in.State //state cache

var apiUrl string
var apiToken string
var devices []string
var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	apiUrl = os.Getenv("API_URL")
	apiToken = os.Getenv("API_TOKEN")
	devices = strings.Split(os.Getenv("DEVICES"), ",") //all devices, comma separated
	currentStatus = make(map[string]in.State)

	//update all periodically
	c := cron.New()
	_ = c.AddFunc("*/60 * * * * *", func() {
		updateAllFromCloud(devices)
	})
	c.Start()

	fmt.Printf("ac_rest_go started with API URL %s\n", apiUrl)
	updateAllFromCloud(devices)

	router := gin.Default()
	router.GET("/devices/:device/status", status)
	router.POST("/devices/:device/power", power)
	router.POST("/devices/:device/ac_mode", acMode)
	router.POST("/devices/:device/fan_mode", fanMode)
	router.POST("/devices/:device/fan_oscillation_mode", fanOscillationMode)
	router.POST("/devices/:device/volume", volume)
	router.POST("/devices/:device/preset", preset)
	router.POST("/devices/:device/temperature", temperature)

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "alive")
	})

	_ = router.Run(":8080")
}

/**
 * Updates all devices from DEVICES env variable (comma separated)
 */
func updateAllFromCloud(devices []string) {
	for _, singleDevice := range devices {
		updateStatusFromCloud(singleDevice)
	}
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
	executeCommand(device, "switch", request.Value, nil)
	state := getCurrentStatus(device)
	state.Components.Main.Switch.Switch.Value = request.Value
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

//POST /ac_mode
func acMode(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/ac_mode for device %s: %s\n", device, request.Value)
	executeCommand(device, "airConditionerMode", "setAirConditionerMode", request.Value)
	state := getCurrentStatus(device)
	state.Components.Main.AirConditionerMode.AirConditionerMode.Value = request.Value
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

//POST /fan_mode
func fanMode(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/fan_mode for device %s: %s\n", device, request.Value)
	executeCommand(device, "airConditionerFanMode", "setFanMode", request.Value)
	state := getCurrentStatus(device)
	state.Components.Main.AirConditionerFanMode.FanMode.Value = request.Value
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

//POST /fan_oscillation_mode
func fanOscillationMode(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/fan_oscillation_mode for device %s: %s\n", device, request.Value)
	executeCommand(device, "fanOscillationMode", "setFanOscillationMode", request.Value)
	state := getCurrentStatus(device)
	state.Components.Main.FanOscillationMode.FanOscillationMode.Value = request.Value
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

//POST /volume
func volume(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/volume for device %s: %s\n", device, request.Value)
	volume, _ := strconv.ParseInt(request.Value, 0, 64)
	executeCommand(device, "audioVolume", "setVolume", volume)
	state := getCurrentStatus(device)
	state.Components.Main.AudioVolume.Volume.Value = int(volume)
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

//POST /preset
func preset(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/preset for device %s: %s\n", device, request.Value)
	executeCommand(device, "custom.airConditionerOptionalMode", "setAcOptionalMode", request.Value)
	state := getCurrentStatus(device)
	state.Components.Main.CustomAirConditionerOptionalMode.AcOptionalMode.Value = request.Value
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

//POST /temperature
func temperature(c *gin.Context) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	device := c.Param("device")
	fmt.Printf("/preset for device %s: %s\n", device, request.Value)
	temp, _ := strconv.ParseFloat(request.Value, 64)
	executeCommand(device, "thermostatCoolingSetpoint", "setCoolingSetpoint", temp)
	state := getCurrentStatus(device)
	state.Components.Main.ThermostatCoolingSetpoint.CoolingSetpoint.Value = temp
	currentStatus[device] = state
	c.JSON(http.StatusOK, state)
}

/**
 * get the device status from cache. If not found, refresh from Samsung SmartThings cloud
 */
func getCurrentStatus(device string) in.State {
	if deviceStatus, found := currentStatus[device]; found {
		fmt.Printf("State for %s, found in cache.\n", device)
		return deviceStatus
	} else {
		fmt.Printf("No current state for %s, updating from the cloud.\n", device)
		return updateStatusFromCloud(device)
	}
}

/**
* Updates the device state from cloud and keeps it in cache
 */
func updateStatusFromCloud(device string) in.State {
	currentStatus[device] = *getStateFromCloud(device)
	return currentStatus[device]
}

/**
* GETs the device state from Samsung SmartThings cloud API
 */
func getStateFromCloud(device string) *in.State {
	req, err := http.NewRequest("GET", apiUrl+device+"/status", nil)
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
	fmt.Printf("Device state updated from cloud: %s\n", device)
	return samsungResponse
}

/**
 * POST a command do Samsung SmartThings cloud
 */
func executeCommand(device string, capability string, command string, param interface{}) {
	cmd := out.Command{Component: "main", Capability: capability, Command: command}
	if param != nil {
		cmd.Arguments = append(cmd.Arguments, param)
	}
	var samsungCmd out.SamsungCommand
	samsungCmd.Commands = append(samsungCmd.Commands, cmd)
	jsonValue, _ := json.Marshal(samsungCmd)
	req, err := http.NewRequest("POST", apiUrl+device+"/commands", bytes.NewBuffer(jsonValue))
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("No response from request %s\n", apiUrl+device)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	samsungResponse := new(in.SamsungResponse)
	if err := json.Unmarshal(body, &samsungResponse); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := body[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
		}
	}
}
