# Wink: Golang API for Wink

Package `wink_exporter` is a [Prometheus](http://prometheus.io/) metrics exporter written in [Go](https://golang.org/)
for [Wink](http://www.wink.com) devices.

## Status

[![Travis Build Status](https://travis-ci.org/joelanford/wink_exporter.png)](https://travis-ci.org/joelanford/wink_exporter)

## Installation 

### Requirements

Wink requires Go 1.5 or later.


### Development

```
go get github.com/joelanford/wink_exporter
```

### Run

```
$ ./wink_exporter -h
Usage of ./wink_exporter:
  -web.listen-address string
    	Address on which to expose metrics and web interface. (default ":9200")
  -web.telemetry-path string
    	Path under which to expose metrics. (default "/metrics")
```

```
$ WINK_CLIENT_ID=<my_wink_client_id> WINK_CLIENT_SECRET=<my_wink_secret> WINK_USERNAME=<my_wink_username> WINK_PASSWORD=<my_wink_password> ./wink_exporter
INFO[0001] Wink authentication successful.
INFO[0001] Listening on :9200 at /metrics.
```

```
$ curl http://localhost:9200/metrics
connection{device_type="hub",model_name="HUB",name="Home"} 1.000000
kidde_radio_code{name="Home",device_type="hub",model_name="HUB"} 0.000000
pairing_mode_duration{device_type="hub",model_name="HUB",name="Home"} 0.000000
remote_pairable{device_type="hub",model_name="HUB",name="Home"} 0.000000
update_needed{device_type="hub",model_name="HUB",name="Home"} 0.000000
updating_firmware{device_type="hub",model_name="HUB",name="Home"} 0.000000
brightness{device_type="light_bulb",model_name="Dimmer",name="Family Room"} 1.000000
connection{device_type="light_bulb",model_name="Dimmer",name="Family Room"} 1.000000
powered{device_type="light_bulb",model_name="Dimmer",name="Family Room"} 1.000000
brightness{name="Lamp",device_type="light_bulb",model_name="GE Light Bulb"} 0.750000
connection{device_type="light_bulb",model_name="GE Light Bulb",name="Lamp"} 1.000000
powered{model_name="GE Light Bulb",name="Lamp",device_type="light_bulb"} 0.000000
connection{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 1.000000
deadband{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 1.500000
eco_target{name="Home Family Room Thermostat",device_type="thermostat",model_name="Learning Thermostat"} 0.000000
external_temperature{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 0.000000
fan_duration{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 0.000000
fan_timer_active{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 0.000000
has_fan{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 1.000000
max_max_set_point{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 22.777778
max_min_set_point{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 19.444444
max_set_point{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 22.777778
min_min_set_point{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 19.444444
min_max_set_point{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 22.777778
min_set_point{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 19.444444
powered{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 1.000000
smart_schedule_enabled{model_name="Learning Thermostat",name="Home Family Room Thermostat",device_type="thermostat"} 0.000000
temperature{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 22.500000
users_away{device_type="thermostat",model_name="Learning Thermostat",name="Home Family Room Thermostat"} 0.000000
brightness{device_type="light_bulb",model_name="Dimmer",name="Master Bedroom"} 0.750000
connection{device_type="light_bulb",model_name="Dimmer",name="Master Bedroom"} 1.000000
powered{name="Master Bedroom",device_type="light_bulb",model_name="Dimmer"} 1.000000
```


## Support

This is an unofficial Prometheus exporter for the Wink API. Please do not contact
Wink or Prometheus unless you're sure the problem is not in this package. Please
submit issues and/or pull requests if you find issues with the package.


## Contributing

Contributions in the form of Pull Requests are gladly accepted.

If you have a Wink device that's not yet supported, please create a new issue
in the [wink](https://github.com/joelanford/wink) project containing the full
JSON device output or a pull request that adds support for your device. 


## License

This is Free Software, released under the terms of the [GPL v3](http://www.gnu.org/licenses/gpl-3.0.en.html).
