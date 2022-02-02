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

//POST /temperature
func temperature(c *gin.Context) {
	state, device, param, err := processRequest(c, "thermostatCoolingSetpoint", "setCoolingSetpoint")
	if err != nil {
		state.Components.Main.ThermostatCoolingSetpoint.CoolingSetpoint.Value = param //update the cache
		currentStatus[device] = state
		c.JSON(http.StatusOK, state)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//POST /ac_mode
func acMode(c *gin.Context) {
	state, device, param, err := processRequest(c, "airConditionerMode", "setAirConditionerMode")
	if err != nil {
		state.Components.Main.AirConditionerMode.AirConditionerMode.Value = param //update the cache
		currentStatus[device] = state
		c.JSON(http.StatusOK, state)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//POST /fan_mode
func fanMode(c *gin.Context) {
	state, device, param, err := processRequest(c, "airConditionerFanMode", "setFanMode")
	if err != nil {
		state.Components.Main.AirConditionerFanMode.FanMode.Value = param //update the cache
		currentStatus[device] = state
		c.JSON(http.StatusOK, state)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//POST /fan_oscillation_mode
func fanOscillationMode(c *gin.Context) {
	state, device, param, err := processRequest(c, "fanOscillationMode", "setFanOscillationMode")
	if err != nil {
		state.Components.Main.FanOscillationMode.FanOscillationMode.Value = param //update the cache
		currentStatus[device] = state
		c.JSON(http.StatusOK, state)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//POST /volume
func volume(c *gin.Context) {
	state, device, param, err := processRequest(c, "audioVolume", "setVolume")
	if err != nil {
		state.Components.Main.AudioVolume.Volume.Value = param //update the cache
		currentStatus[device] = state
		c.JSON(http.StatusOK, state)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

//POST /preset
func preset(c *gin.Context) {
	state, device, param, err := processRequest(c, "custom.airConditionerOptionalMode", "setAcOptionalMode")
	if err != nil {
		state.Components.Main.CustomAirConditionerOptionalMode.AcOptionalMode.Value = param //update the cache
		currentStatus[device] = state
		c.JSON(http.StatusOK, state)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func processRequest(c *gin.Context, capability string, command string) (in.State, string, string, error) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return in.State{}, request.Value, "", err
	}
	device := c.Param("device")
	fmt.Printf("capability: %s, command %s, device %s: %s\n", capability, command, device, request.Value)
	executeCommand(device, capability, command, request.Value)
	state := getCurrentStatus(device)

	return state, device, request.Value, nil
}

/**
 * get the device status from cache. If not found, refresh from Samsung SmartThings cloud
 */
func getCurrentStatus(device string) *in.State {
	if deviceStatus, found := currentStatus[device]; found {
		fmt.Printf("State for %s, found in cache.\n", device)
		return deviceStatus
	} else {
		fmt.Printf("No current state for %s, updating from the cloud.\n", device)
		return updateStatusFromCloud(device)
	}
}

/**
* GETs the device state from Samsung SmartThings cloud API
 */
func updateStatusFromCloud(device string) in.State {
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
	currentStatus[device] = samsungResponse

	return samsungResponse, err
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
