# ac-rest-go

## the challenge
In order to integrate Samsung AC devices with HomeAssistant, you need to open Home Assistant to the Internet, as the official SmartThings integration uses webhooks.

If you want to keep your Home Assistant **not exposed** on the Internet, the https://github.com/SebuZet/samsungrac can be used as an alternative to the official integration.
However, Home Assistant will ask for each device status every 3 seconds.
For multiple AC devices, this will flood Samsung API with frequent multiple requests for each device.
This will render your Samsung API account throttled and then banned.

**ac-rest**  is an auxiliary server for processing and executing AC-conditioners requests to Samsung Smartthings API.
Tested with Samsung WindFree Avant 2,5kW and 5kW devices and https://github.com/SebuZet/samsungrac

ac-rest **uses own status cache** and skips the Samsung API calls.
It will refresh the status from the Samsung Cloud periodically also (once per minute), to sync the status if you use the mobile app to control your unit, for example.

## configuration

Expects ENV variables:
* DEVICES comma separated list of devices to refresh status periodically.
* API_TOKEN your Samsung API token with "Bearer: " prefix.
* API_URL https://api.smartthings.com/v1/devices/

## endpoints exposed:


| method | URL                                    | payload                                                                                                 | comment |
|--------|----------------------------------------|---------------------------------------------------------------------------------------------------------|---------|
| GET    | /devices/{device}/status               | response has the payload the same as https://api.smartthings.com/v1/devices/{deviceId}/status           |         |
| POST   | /devices/{device}/power                | "on", "off                                                                                              |         |
| POST   | /devices/{device}/temperature          | temperature in Celsius, double                                                                          |         |
| POST   | /devices/{device}/ac_mode              | "auto", "cool", "dry", "off"                                                                            |         |
| POST   | /devices/{device}/fan_mode             | "auto", "low", "medium", "high", "turbo"                                                                 |         |
| POST   | /devices/{device}/fan_oscillation_mode | "vertical", "horizontal", "fixed", "fixedLeft", "fixedRight", "fixedCenter", direct", "indirect", "far" |         |
| POST   | /devices/{device}/beep                 | "on", "off"                                                                                             |         |
| POST   | /devices/{device}/preset               | "off", "speed", "sleep", "windFree", "windFreeSleep"                                                    |         |

## parameters (POST payload)

{device} is your device id, taken from https://api.smartthings.com/v1/devices/

```json
{"value": "yourValueHere"}

```

## example
Set power on, then set ac-mode to cooling at 21C

```
POST /devices/{device}/power  {"value": "on"}
POST /devices/{device}/temperature  {"value": 21}
POST /devices/{device}/ac_mode  {"value": "cool"}
```

## running
Can be run in Docker, for example:

```bash
docker run \
    -d \
    --name=acrest \
    --restart=unless-stopped \
    --env=DEVICES="edf3d10c-redacted-fb7f-redacted,redacted-e88e-redcated,9d8f6a8c-redacted-ed416e8b3fd1,redacted9e7c-ad8a-redacted" \
    --env=API_TOKEN="Bearer: xxx-yyy" \
    --env=API_URL="https://api.smartthings.com/v1/devices/" \
    --network=bridge \
    -p 8080:8080 \
    --memory=200m \
    ac-rest:latest
```
Having this running, you can use the https://github.com/SebuZet/samsungrac to configure the IP climate setup in Home Assistant.
See `home_assistant/climate_ip/smartthings.yaml` as an example configuration file for https://github.com/SebuZet/samsungrac to use with ac-rest.

## building locally
```bash
go build -v -o build/package/ac-rest
```


## Docker image build
```bash
docker build . -t jotka/ac-rest
```
