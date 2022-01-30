package in

import "time"

type State struct {
	Components Components `json:"components"`
}
type Humidity struct {
	Value     int       `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type RelativeHumidityMeasurement struct {
	Humidity Humidity `json:"humidity"`
}
type AirConditionerOdorControllerProgress struct {
	Value interface{} `json:"value"`
}
type AirConditionerOdorControllerState struct {
	Value interface{} `json:"value"`
}
type CustomAirConditionerOdorController struct {
	AirConditionerOdorControllerProgress AirConditionerOdorControllerProgress `json:"airConditionerOdorControllerProgress"`
	AirConditionerOdorControllerState    AirConditionerOdorControllerState    `json:"airConditionerOdorControllerState"`
}
type MinimumSetpoint struct {
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type MaximumSetpoint struct {
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomThermostatSetpointControl struct {
	MinimumSetpoint MinimumSetpoint `json:"minimumSetpoint"`
	MaximumSetpoint MaximumSetpoint `json:"maximumSetpoint"`
}
type SupportedAcModes struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type AirConditionerMode struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type SpiMode struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomSpiMode struct {
	SpiMode SpiMode `json:"spiMode"`
}
type AirQuality struct {
	Value interface{} `json:"value"`
}
type AirQualitySensor struct {
	AirQuality AirQuality `json:"airQuality"`
}
type SupportedAcOptionalMode struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type AcOptionalMode struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomAirConditionerOptionalMode struct {
	SupportedAcOptionalMode SupportedAcOptionalMode `json:"supportedAcOptionalMode"`
	AcOptionalMode          AcOptionalMode          `json:"acOptionalMode"`
}
type Switch_ struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Switch struct {
	Switch Switch_ `json:"switch"`
}
type AcTropicalNightModeLevel struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomAirConditionerTropicalNightMode struct {
	AcTropicalNightModeLevel AcTropicalNightModeLevel `json:"acTropicalNightModeLevel"`
}
type St struct {
	Value interface{} `json:"value"`
}
type Mndt struct {
	Value interface{} `json:"value"`
}
type Mnfv struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnhw struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Di struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnsl struct {
	Value interface{} `json:"value"`
}
type Dmv struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type N struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnmo struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Vid struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnmn struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnml struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnpv struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Mnos struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Pi struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Icv struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Ocf struct {
	St   St   `json:"st"`
	Mndt Mndt `json:"mndt"`
	Mnfv Mnfv `json:"mnfv"`
	Mnhw Mnhw `json:"mnhw"`
	Di   Di   `json:"di"`
	Mnsl Mnsl `json:"mnsl"`
	Dmv  Dmv  `json:"dmv"`
	N    N    `json:"n"`
	Mnmo Mnmo `json:"mnmo"`
	Vid  Vid  `json:"vid"`
	Mnmn Mnmn `json:"mnmn"`
	Mnml Mnml `json:"mnml"`
	Mnpv Mnpv `json:"mnpv"`
	Mnos Mnos `json:"mnos"`
	Pi   Pi   `json:"pi"`
	Icv  Icv  `json:"icv"`
}
type FanMode struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type SupportedAcFanModes struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type AirConditionerFanMode struct {
	FanMode             FanMode             `json:"fanMode"`
	SupportedAcFanModes SupportedAcFanModes `json:"supportedAcFanModes"`
}
type AlarmThreshold struct {
	Value     int       `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type SupportedAlarmThresholds struct {
	Value     []int     `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type SamsungceDustFilterAlarm struct {
	AlarmThreshold           AlarmThreshold           `json:"alarmThreshold"`
	SupportedAlarmThresholds SupportedAlarmThresholds `json:"supportedAlarmThresholds"`
}
type ElectricHepaFilterCapacity struct {
	Value interface{} `json:"value"`
}
type ElectricHepaFilterUsageStep struct {
	Value interface{} `json:"value"`
}
type ElectricHepaFilterLastResetDate struct {
	Value interface{} `json:"value"`
}
type ElectricHepaFilterStatus struct {
	Value interface{} `json:"value"`
}
type ElectricHepaFilterUsage struct {
	Value interface{} `json:"value"`
}
type ElectricHepaFilterResetType struct {
	Value interface{} `json:"value"`
}
type CustomElectricHepaFilter struct {
	ElectricHepaFilterCapacity      ElectricHepaFilterCapacity      `json:"electricHepaFilterCapacity"`
	ElectricHepaFilterUsageStep     ElectricHepaFilterUsageStep     `json:"electricHepaFilterUsageStep"`
	ElectricHepaFilterLastResetDate ElectricHepaFilterLastResetDate `json:"electricHepaFilterLastResetDate"`
	ElectricHepaFilterStatus        ElectricHepaFilterStatus        `json:"electricHepaFilterStatus"`
	ElectricHepaFilterUsage         ElectricHepaFilterUsage         `json:"electricHepaFilterUsage"`
	ElectricHepaFilterResetType     ElectricHepaFilterResetType     `json:"electricHepaFilterResetType"`
}
type DisabledCapabilities struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomDisabledCapabilities struct {
	DisabledCapabilities DisabledCapabilities `json:"disabledCapabilities"`
}
type OcfResourceUpdatedTime struct {
	Value interface{} `json:"value"`
}
type OcfResourceVersion struct {
	Value interface{} `json:"value"`
}
type CustomOcfResourceVersion struct {
	OcfResourceUpdatedTime OcfResourceUpdatedTime `json:"ocfResourceUpdatedTime"`
	OcfResourceVersion     OcfResourceVersion     `json:"ocfResourceVersion"`
}
type VersionNumber struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type SamsungceDriverVersion struct {
	VersionNumber VersionNumber `json:"versionNumber"`
}
type SupportedFanOscillationModes struct {
	Value interface{} `json:"value"`
}
type FanOscillationMode_ struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type FanOscillationMode struct {
	SupportedFanOscillationModes SupportedFanOscillationModes `json:"supportedFanOscillationModes"`
	FanOscillationMode           FanOscillationMode_          `json:"fanOscillationMode"`
}
type Temperature struct {
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type TemperatureMeasurement struct {
	Temperature Temperature `json:"temperature"`
}
type DustLevel struct {
	Value interface{} `json:"value"`
}
type FineDustLevel struct {
	Value interface{} `json:"value"`
}
type DustSensor struct {
	DustLevel     DustLevel     `json:"dustLevel"`
	FineDustLevel FineDustLevel `json:"fineDustLevel"`
}
type ReportStateRealtimePeriod struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type Value struct {
	State string `json:"state"`
}
type ReportStateRealtime struct {
	Value     Value     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type ReportStatePeriod struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomDeviceReportStateConfiguration struct {
	ReportStateRealtimePeriod ReportStateRealtimePeriod `json:"reportStateRealtimePeriod"`
	ReportStateRealtime       ReportStateRealtime       `json:"reportStateRealtime"`
	ReportStatePeriod         ReportStatePeriod         `json:"reportStatePeriod"`
}
type AutomaticExecutionSetting struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type AutomaticExecutionMode struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type SupportedAutomaticExecutionSetting struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type SupportedAutomaticExecutionMode struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type PeriodicSensing struct {
	Value interface{} `json:"value"`
}
type PeriodicSensingInterval struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type LastSensingTime struct {
	Value interface{} `json:"value"`
}
type LastSensingLevel struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type PeriodicSensingStatus struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomPeriodicSensing struct {
	AutomaticExecutionSetting          AutomaticExecutionSetting          `json:"automaticExecutionSetting"`
	AutomaticExecutionMode             AutomaticExecutionMode             `json:"automaticExecutionMode"`
	SupportedAutomaticExecutionSetting SupportedAutomaticExecutionSetting `json:"supportedAutomaticExecutionSetting"`
	SupportedAutomaticExecutionMode    SupportedAutomaticExecutionMode    `json:"supportedAutomaticExecutionMode"`
	PeriodicSensing                    PeriodicSensing                    `json:"periodicSensing"`
	PeriodicSensingInterval            PeriodicSensingInterval            `json:"periodicSensingInterval"`
	LastSensingTime                    LastSensingTime                    `json:"lastSensingTime"`
	LastSensingLevel                   LastSensingLevel                   `json:"lastSensingLevel"`
	PeriodicSensingStatus              PeriodicSensingStatus              `json:"periodicSensingStatus"`
}
type CoolingSetpoint struct {
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type ThermostatCoolingSetpoint struct {
	CoolingSetpoint CoolingSetpoint `json:"coolingSetpoint"`
}

type DrlcStatus struct {
	Value     Value     `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type DemandResponseLoadControl struct {
	DrlcStatus DrlcStatus `json:"drlcStatus"`
}
type Volume struct {
	Value     int       `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type AudioVolume struct {
	Volume Volume `json:"volume"`
}
type PowerConsumption struct {
	Value interface{} `json:"value"`
}
type PowerConsumptionReport struct {
	PowerConsumption PowerConsumption `json:"powerConsumption"`
}
type AutoCleaningMode struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomAutoCleaningMode struct {
	AutoCleaningMode AutoCleaningMode `json:"autoCleaningMode"`
}
type LockState struct {
	Value interface{} `json:"value"`
}
type SamsungceIndividualControlLock struct {
	LockState LockState `json:"lockState"`
}
type Refresh struct {
}
type Payload struct {
	Rt                                  []string `json:"rt"`
	If                                  []string `json:"if"`
	XComSamsungDaInstantaneousPower     string   `json:"x.com.samsung.da.instantaneousPower"`
	XComSamsungDaCumulativePower        string   `json:"x.com.samsung.da.cumulativePower"`
	XComSamsungDaCumulativeDate         string   `json:"x.com.samsung.da.cumulativeDate"`
	XComSamsungDaCumulativeDateUTC      string   `json:"x.com.samsung.da.cumulativeDateUTC"`
	XComSamsungDaCumulativeUnit         string   `json:"x.com.samsung.da.cumulativeUnit"`
	XComSamsungDaInstantaneousPowerUnit string   `json:"x.com.samsung.da.instantaneousPowerUnit"`
}

type DustFilterUsageStep struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type DustFilterUsage struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type DustFilterLastResetDate struct {
	Value interface{} `json:"value"`
}
type DustFilterStatus struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type DustFilterCapacity struct {
	Value     int       `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type DustFilterResetType struct {
	Value     []string  `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomDustFilter struct {
	DustFilterUsageStep     DustFilterUsageStep     `json:"dustFilterUsageStep"`
	DustFilterUsage         DustFilterUsage         `json:"dustFilterUsage"`
	DustFilterLastResetDate DustFilterLastResetDate `json:"dustFilterLastResetDate"`
	DustFilterStatus        DustFilterStatus        `json:"dustFilterStatus"`
	DustFilterCapacity      DustFilterCapacity      `json:"dustFilterCapacity"`
	DustFilterResetType     DustFilterResetType     `json:"dustFilterResetType"`
}
type OdorLevel struct {
	Value interface{} `json:"value"`
}
type OdorSensor struct {
	OdorLevel OdorLevel `json:"odorLevel"`
}
type RemoteControlEnabled struct {
	Value interface{} `json:"value"`
}
type RemoteControlStatus struct {
	RemoteControlEnabled RemoteControlEnabled `json:"remoteControlEnabled"`
}
type DeodorFilterLastResetDate struct {
	Value interface{} `json:"value"`
}
type DeodorFilterCapacity struct {
	Value interface{} `json:"value"`
}
type DeodorFilterStatus struct {
	Value interface{} `json:"value"`
}
type DeodorFilterResetType struct {
	Value interface{} `json:"value"`
}
type DeodorFilterUsage struct {
	Value interface{} `json:"value"`
}
type DeodorFilterUsageStep struct {
	Value interface{} `json:"value"`
}
type CustomDeodorFilter struct {
	DeodorFilterLastResetDate DeodorFilterLastResetDate `json:"deodorFilterLastResetDate"`
	DeodorFilterCapacity      DeodorFilterCapacity      `json:"deodorFilterCapacity"`
	DeodorFilterStatus        DeodorFilterStatus        `json:"deodorFilterStatus"`
	DeodorFilterResetType     DeodorFilterResetType     `json:"deodorFilterResetType"`
	DeodorFilterUsage         DeodorFilterUsage         `json:"deodorFilterUsage"`
	DeodorFilterUsageStep     DeodorFilterUsageStep     `json:"deodorFilterUsageStep"`
}
type EnergyType struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type EnergySavingSupport struct {
	Value     bool      `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type DrMaxDuration struct {
	Value     int       `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
}
type EnergySavingOperation struct {
	Value interface{} `json:"value"`
}
type EnergySavingOperationSupport struct {
	Value     bool      `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomEnergyType struct {
	EnergyType                   EnergyType                   `json:"energyType"`
	EnergySavingSupport          EnergySavingSupport          `json:"energySavingSupport"`
	DrMaxDuration                DrMaxDuration                `json:"drMaxDuration"`
	EnergySavingOperation        EnergySavingOperation        `json:"energySavingOperation"`
	EnergySavingOperationSupport EnergySavingOperationSupport `json:"energySavingOperationSupport"`
}
type OtnDUID struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type AvailableModules struct {
	Value     []interface{} `json:"value"`
	Timestamp time.Time     `json:"timestamp"`
}
type NewVersionAvailable struct {
	Value     bool      `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type SamsungceSoftwareUpdate struct {
	OtnDUID             OtnDUID             `json:"otnDUID"`
	AvailableModules    AvailableModules    `json:"availableModules"`
	NewVersionAvailable NewVersionAvailable `json:"newVersionAvailable"`
}
type VeryFineDustLevel struct {
	Value interface{} `json:"value"`
}
type VeryFineDustSensor struct {
	VeryFineDustLevel VeryFineDustLevel `json:"veryFineDustLevel"`
}
type VeryFineDustFilterStatus struct {
	Value interface{} `json:"value"`
}
type VeryFineDustFilterResetType struct {
	Value interface{} `json:"value"`
}
type VeryFineDustFilterUsage struct {
	Value interface{} `json:"value"`
}
type VeryFineDustFilterLastResetDate struct {
	Value interface{} `json:"value"`
}
type VeryFineDustFilterUsageStep struct {
	Value interface{} `json:"value"`
}
type VeryFineDustFilterCapacity struct {
	Value interface{} `json:"value"`
}
type CustomVeryFineDustFilter struct {
	VeryFineDustFilterStatus        VeryFineDustFilterStatus        `json:"veryFineDustFilterStatus"`
	VeryFineDustFilterResetType     VeryFineDustFilterResetType     `json:"veryFineDustFilterResetType"`
	VeryFineDustFilterUsage         VeryFineDustFilterUsage         `json:"veryFineDustFilterUsage"`
	VeryFineDustFilterLastResetDate VeryFineDustFilterLastResetDate `json:"veryFineDustFilterLastResetDate"`
	VeryFineDustFilterUsageStep     VeryFineDustFilterUsageStep     `json:"veryFineDustFilterUsageStep"`
	VeryFineDustFilterCapacity      VeryFineDustFilterCapacity      `json:"veryFineDustFilterCapacity"`
}
type DoNotDisturb struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type StartTime struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type EndTime struct {
	Value     string    `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
type CustomDoNotDisturbMode struct {
	DoNotDisturb DoNotDisturb `json:"doNotDisturb"`
	StartTime    StartTime    `json:"startTime"`
	EndTime      EndTime      `json:"endTime"`
}
type Main struct {
	RelativeHumidityMeasurement           RelativeHumidityMeasurement           `json:"relativeHumidityMeasurement"`
	CustomAirConditionerOdorController    CustomAirConditionerOdorController    `json:"custom.airConditionerOdorController"`
	CustomThermostatSetpointControl       CustomThermostatSetpointControl       `json:"custom.thermostatSetpointControl"`
	AirConditionerMode                    AirConditionerMode                    `json:"airConditionerMode"`
	CustomSpiMode                         CustomSpiMode                         `json:"custom.spiMode"`
	AirQualitySensor                      AirQualitySensor                      `json:"airQualitySensor"`
	CustomAirConditionerOptionalMode      CustomAirConditionerOptionalMode      `json:"custom.airConditionerOptionalMode"`
	Switch                                Switch                                `json:"switch"`
	CustomAirConditionerTropicalNightMode CustomAirConditionerTropicalNightMode `json:"custom.airConditionerTropicalNightMode"`
	Ocf                                   Ocf                                   `json:"ocf"`
	AirConditionerFanMode                 AirConditionerFanMode                 `json:"airConditionerFanMode"`
	SamsungceDustFilterAlarm              SamsungceDustFilterAlarm              `json:"samsungce.dustFilterAlarm"`
	CustomElectricHepaFilter              CustomElectricHepaFilter              `json:"custom.electricHepaFilter"`
	CustomDisabledCapabilities            CustomDisabledCapabilities            `json:"custom.disabledCapabilities"`
	CustomOcfResourceVersion              CustomOcfResourceVersion              `json:"custom.ocfResourceVersion"`
	SamsungceDriverVersion                SamsungceDriverVersion                `json:"samsungce.driverVersion"`
	FanOscillationMode                    FanOscillationMode                    `json:"fanOscillationMode"`
	TemperatureMeasurement                TemperatureMeasurement                `json:"temperatureMeasurement"`
	DustSensor                            DustSensor                            `json:"dustSensor"`
	CustomDeviceReportStateConfiguration  CustomDeviceReportStateConfiguration  `json:"custom.deviceReportStateConfiguration"`
	CustomPeriodicSensing                 CustomPeriodicSensing                 `json:"custom.periodicSensing"`
	ThermostatCoolingSetpoint             ThermostatCoolingSetpoint             `json:"thermostatCoolingSetpoint"`
	DemandResponseLoadControl             DemandResponseLoadControl             `json:"demandResponseLoadControl"`
	AudioVolume                           AudioVolume                           `json:"audioVolume"`
	PowerConsumptionReport                PowerConsumptionReport                `json:"powerConsumptionReport"`
	CustomAutoCleaningMode                CustomAutoCleaningMode                `json:"custom.autoCleaningMode"`
	SamsungceIndividualControlLock        SamsungceIndividualControlLock        `json:"samsungce.individualControlLock"`
	Refresh                               Refresh                               `json:"refresh"`
	CustomDustFilter                      CustomDustFilter                      `json:"custom.dustFilter"`
	OdorSensor                            OdorSensor                            `json:"odorSensor"`
	RemoteControlStatus                   RemoteControlStatus                   `json:"remoteControlStatus"`
	CustomDeodorFilter                    CustomDeodorFilter                    `json:"custom.deodorFilter"`
	CustomEnergyType                      CustomEnergyType                      `json:"custom.energyType"`
	SamsungceSoftwareUpdate               SamsungceSoftwareUpdate               `json:"samsungce.softwareUpdate"`
	VeryFineDustSensor                    VeryFineDustSensor                    `json:"veryFineDustSensor"`
	CustomVeryFineDustFilter              CustomVeryFineDustFilter              `json:"custom.veryFineDustFilter"`
	CustomDoNotDisturbMode                CustomDoNotDisturbMode                `json:"custom.doNotDisturbMode"`
}
type Components struct {
	Main Main `json:"main"`
}
