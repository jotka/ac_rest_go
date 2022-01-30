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
var devices []string
var httpClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	apiUrl = os.Getenv("API_URL")
	apiToken = os.Getenv("API_TOKEN")
	devices = strings.Split(os.Getenv("DEVICES"), ",") //all devices, comma separated
	currentStatus = make(map[string]in.State)

	//update all periodically
	cron := cron.New()
	cron.AddFunc("*/60 * * * * *", func() {
		updateAllFromCloud(devices)
	})
	cron.Start()

	fmt.Printf("ac_rest_go started with API URL %s\n", apiUrl)
	updateAllFromCloud(devices)

	router := gin.Default()
	router.GET("/device/:device/status", status)
	router.POST("/device/:device/power", power)
	router.POST("/device/:device/ac_mode", ac_mode)
	router.POST("/device/:device/fan_mode", fan_mode)
	router.POST("/device/:device/fan_oscillation_mode", fan_oscillation_mode)
	router.POST("/device/:device/volume", volume)
	router.POST("/device/:device/preset", preset)
	router.POST("/device/:device/temperature", temperature)

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "alive")
	})

	router.Run(":8080")
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
	executeCommand(device, "airConditionerMode", "setAirConditionerMode", request.Value)
	state := getCurrentStatus(device)
	state.Components.Main.AirConditionerMode.Value = request.Value
	c.JSON(http.StatusOK, state)
}

//POST /fan_mode
func fan_mode(c *gin.Context) {
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
	c.JSON(http.StatusOK, state)
}

//POST /fan_oscillation_mode
func fan_oscillation_mode(c *gin.Context) {
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
	c.JSON(http.StatusOK, state)
}

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
