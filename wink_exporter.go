package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joelanford/wink"
)

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":9200", "Address on which to expose metrics and web interface.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")

		winkClientID     = os.Getenv("WINK_CLIENT_ID")
		winkClientSecret = os.Getenv("WINK_CLIENT_SECRET")
		winkUsername     = os.Getenv("WINK_USERNAME")
		winkPassword     = os.Getenv("WINK_PASSWORD")
	)
	flag.Parse()

	c := wink.NewClient(winkClientID, winkClientSecret)
	err := c.Authenticate(winkUsername, winkPassword)
	if err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Wink authentication successful.")

	r := mux.NewRouter()
	r.HandleFunc(*metricsPath, func(w http.ResponseWriter, r *http.Request) {
		metrics, _ := getMetrics(c)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		for _, metric := range metrics {
			fmt.Fprintf(w, "%s\n", metric)
		}
	})

	log.Infoln("Listening on " + *listenAddress + " at " + *metricsPath + ".")
	http.ListenAndServe(*listenAddress, handlers.LoggingHandler(os.Stdout, r))
}

func getMetrics(c *wink.Client) ([]string, error) {
	devices, err := c.GetDevices()
	var metrics []string
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if d, ok := device.(*wink.LightBulb); ok {
			labels := map[string]string{
				"device_type": "light_bulb",
				"model_name":  d.ModelName,
				"name":        d.Name,
			}
			metrics = append(metrics, formatMetric("brightness", labels, d.LastReading.Brightness))
			metrics = append(metrics, formatMetric("connection", labels, boolToFloat64(d.LastReading.Connection)))
			metrics = append(metrics, formatMetric("powered", labels, boolToFloat64(d.LastReading.Powered)))
		} else if d, ok := device.(*wink.Thermostat); ok {
			labels := map[string]string{
				"device_type": "thermostat",
				"model_name":  d.ModelName,
				"name":        d.Name,
			}
			metrics = append(metrics, formatMetric("connection", labels, boolToFloat64(d.LastReading.Connection)))
			metrics = append(metrics, formatMetric("deadband", labels, d.LastReading.Deadband))
			metrics = append(metrics, formatMetric("eco_target", labels, boolToFloat64(d.LastReading.EcoTarget)))
			metrics = append(metrics, formatMetric("external_temperature", labels, d.LastReading.ExternalTemperature))
			metrics = append(metrics, formatMetric("fan_duration", labels, float64(d.LastReading.FanDuration)))
			metrics = append(metrics, formatMetric("fan_timer_active", labels, boolToFloat64(d.LastReading.FanTimerActive)))
			metrics = append(metrics, formatMetric("has_fan", labels, boolToFloat64(d.LastReading.HasFan)))
			metrics = append(metrics, formatMetric("max_max_set_point", labels, d.LastReading.MaxSetPoint))
			metrics = append(metrics, formatMetric("max_min_set_point", labels, d.LastReading.MinSetPoint))
			metrics = append(metrics, formatMetric("max_set_point", labels, d.LastReading.MaxSetPoint))
			metrics = append(metrics, formatMetric("min_min_set_point", labels, d.LastReading.MinSetPoint))
			metrics = append(metrics, formatMetric("min_max_set_point", labels, d.LastReading.MaxSetPoint))
			metrics = append(metrics, formatMetric("min_set_point", labels, d.LastReading.MinSetPoint))
			metrics = append(metrics, formatMetric("powered", labels, boolToFloat64(d.LastReading.Powered)))
			metrics = append(metrics, formatMetric("smart_schedule_enabled", labels, boolToFloat64(d.SmartScheduleEnabled)))
			metrics = append(metrics, formatMetric("temperature", labels, d.LastReading.Temperature))
			metrics = append(metrics, formatMetric("users_away", labels, boolToFloat64(d.LastReading.UsersAway)))
		} else if d, ok := device.(*wink.Hub); ok {
			labels := map[string]string{
				"device_type": "hub",
				"model_name":  d.ModelName,
				"name":        d.Name,
			}
			metrics = append(metrics, formatMetric("connection", labels, boolToFloat64(d.LastReading.Connection)))
			metrics = append(metrics, formatMetric("kidde_radio_code", labels, float64(d.Configuration.KiddeRadioCode)))
			metrics = append(metrics, formatMetric("pairing_mode_duration", labels, float64(d.LastReading.PairingModeDuration)))
			metrics = append(metrics, formatMetric("remote_pairable", labels, boolToFloat64(d.LastReading.RemotePairable)))
			metrics = append(metrics, formatMetric("update_needed", labels, boolToFloat64(d.LastReading.UpdateNeeded)))
			metrics = append(metrics, formatMetric("updating_firmware", labels, boolToFloat64(d.LastReading.UpdatingFirmware)))
		}
	}
	return metrics, nil
}

func formatMetric(name string, labels map[string]string, value float64) string {
	var labelSlice []string
	for k, v := range labels {
		labelSlice = append(labelSlice, k+"="+"\""+v+"\"")
	}
	labelString := strings.Join(labelSlice, ",")

	return fmt.Sprintf("%s{%s} %f", name, labelString, value)
}

func boolToFloat64(value bool) float64 {
	if value {
		return 1.0
	}
	return 0.0
}
