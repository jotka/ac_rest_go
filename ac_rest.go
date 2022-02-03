package main

import (
	"ac_rest_go/in"
	"ac_rest_go/out"
	"bytes"
	"encoding/json"
	"errors"
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
		state, err := getStatusFromCloud(singleDevice)
		if err != nil {
			fmt.Printf("Error updating %s from the cloud: %s\n", singleDevice, err.Error())
		} else {
			currentStatus[singleDevice] = *state
		}
	}
}

//GET /status
func status(c *gin.Context) {
	device := c.Param("device")
	fmt.Printf("/status for device %s\n", device)
	state, err := getCurrentStatus(device)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, state)
	}
}

//POST /power
func power(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := executeCommand(device, "switch", param, "")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.Switch.Switch.Value = param
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}
	}
}

//POST /temperature
func temperature(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	desiredTemp, _ := strconv.ParseFloat(param, 64)
	response, err := executeCommand(device, "thermostatCoolingSetpoint", "setCoolingSetpoint", desiredTemp)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.ThermostatCoolingSetpoint.CoolingSetpoint.Value = desiredTemp //update the cache
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}
	}
}

//POST /ac_mode
func acMode(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := executeCommand(device, "airConditionerMode", "setAirConditionerMode", param)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.AirConditionerMode.AirConditionerMode.Value = param
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}

	}
}

//POST /fan_mode
func fanMode(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := executeCommand(device, "airConditionerFanMode", "setFanMode", param)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.AirConditionerFanMode.FanMode.Value = param //update the cache
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}
	}
}

//POST /fan_oscillation_mode
func fanOscillationMode(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := executeCommand(device, "fanOscillationMode", "setFanOscillationMode", param)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.FanOscillationMode.FanOscillationMode.Value = param //update the cache
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}
	}
}

//POST /volume
func volume(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	volume, _ := strconv.Atoi(param)
	response, err := executeCommand(device, "audioVolume", "setVolume", volume)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.AudioVolume.Volume.Value = volume
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}
	}
}

//POST /preset
func preset(c *gin.Context) {
	device, param, err := parseRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	response, err := executeCommand(device, "custom.airConditionerOptionalMode", "setAcOptionalMode", param)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		state, err := getCurrentStatus(device)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			state.Components.Main.CustomAirConditionerOptionalMode.AcOptionalMode.Value = param //update the cache
			currentStatus[device] = *state
			c.JSON(http.StatusOK, response)
		}
	}
}

func parseRequest(c *gin.Context) (string, string, error) {
	var request in.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		return "", "", err
	}
	device := c.Param("device")
	param := request.Value
	return device, param, nil
}

/**
 * get the device status from cache. If not found, refresh from Samsung SmartThings cloud
 */
func getCurrentStatus(device string) (*in.State, error) {
	if deviceStatus, found := currentStatus[device]; found {
		fmt.Printf("State for %s, found in cache.\n", device)
		return &deviceStatus, nil
	} else {
		fmt.Printf("No current state for %s, updating from the cloud.\n", device)
		return getStatusFromCloud(device)
	}
}

/**
* GETs the device state from Samsung SmartThings cloud API
 */
func getStatusFromCloud(device string) (*in.State, error) {
	req, err := http.NewRequest("GET", apiUrl+device+"/status", nil)
	req.Header.Add("Authorization", apiToken)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("No response from request %s\n", apiUrl+device)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		samsungResponse := new(in.State)
		err := json.Unmarshal(body, &samsungResponse)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Device state updated from cloud: %s\n", device)
		return samsungResponse, nil
	}
}

/**
 * POST a command do Samsung SmartThings cloud
 */
func executeCommand(device string, capability string, command string, param interface{}) (*in.SamsungResponse, error) {
	fmt.Printf("capability: %s, command %s, device %s: %s\n", capability, command, device, param)
	cmd := out.Command{Component: "main", Capability: capability, Command: command}
	if param != "" {
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
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	samsungResponse := new(in.SamsungResponse)
	if err := json.Unmarshal(body, &samsungResponse); err != nil {
		if jsonErr, ok := err.(*json.SyntaxError); ok {
			problemPart := body[jsonErr.Offset-10 : jsonErr.Offset+10]
			err = fmt.Errorf("%w ~ error near '%s' (offset %d)", err, problemPart, jsonErr.Offset)
			return nil, err
		}
	}
	if samsungResponse.Error != nil {
		return samsungResponse, errors.New(samsungResponse.Error.Message)
	}
	return samsungResponse, nil
}
