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
* DEVICES comma separated list of your devices, use https://api.smartthings.com/v1/devices/ to get them
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

## running in Docker

```bash
docker run --name=acrest \
--env=API_URL=https://api.smartthings.com/v1/devices/ \
--env=DEVICES=edf3d10c-0c57-fb7f-removed,removed-e88e-a7f13bb87e3b,removed-b03f-ed416e8b3fd1,837249dd-9e7c-ad8a-removed-6b8cf5da271b \
--env=GIN_MODE=release \
--env='API_TOKEN=Bearer: your_smartthings_token_here' \
--network=bridge \
-p 8888:8080 \
--restart=unless-stopped \
--detach=true jotka/ac-rest:latest
```
Having this running, you can use the https://github.com/SebuZet/samsungrac to configure the IP climate setup in Home Assistant.
See `home_assistant/` as an example configuration with https://github.com/SebuZet/samsungrac custom component.
`climate_ip/smartthings.yaml` contains the template for ac-rest middleware.

## building locally
```bash
go build -v -o build/package/ac-rest
```

## Docker image build
```bash
docker build . -t jotka/ac-rest
```
