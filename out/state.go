package out

type FanMode struct {
	value string
}

type AirConditionerMode struct {
	fanMode FanMode
}

type CustomAirConditionerOptionalMode struct {
	fanMode FanMode
}

type Switch struct {
	fanMode FanMode
	value   string
}

type AirConditionerFanMode struct {
	fanMode FanMode
}

type FanOscillationMode struct {
	fanMode FanMode
}

type Volume struct {
}

type TemperatureMeasurement struct {
	temperature Volume
}

type ThermostatCoolingSetpoint struct {
	coolingSetpoint Volume
}

type AudioVolume struct {
	volume Volume
}

type Main struct {
	airConditionerMode               AirConditionerMode
	customAirConditionerOptionalMode CustomAirConditionerOptionalMode
	_switch                          Switch `form:"switch" json:"switch" xml:"switch"`
	airConditionerFanMode            AirConditionerFanMode
	fanOscillationMode               FanOscillationMode
	temperatureMeasurement           TemperatureMeasurement
	thermostatCoolingSetpoint        ThermostatCoolingSetpoint
	audioVolume                      AudioVolume
}

type Component struct {
	main Main
}

type State struct {
	components []Component
}
