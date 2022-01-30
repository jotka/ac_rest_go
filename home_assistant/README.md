# HomeAssistant custom component

## usage
place in your home_assistant/custom_components directory

and then cofigure:
```yaml
climate: 
- platform: climate_ip
  config_file: 'smartthings.yaml'
  name: 'klima_pracownia'
  ip_address: 'ip and port of the ac-rest running'
  device_id: 'your device ID here'
  debug: false
```
