package airvisual

import "fmt"

type CityData struct {
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Location `json:"location"`
	Current  `json:"current"`
}

type Location struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Current struct {
	Pollution `json:"pollution"`
	Weather   `json:"weather"`
}

type Pollution struct {
	Timestamp string `json:"ts"`
	Aqius     int16  `json:"aqius"`
	Mainus    string `json:"mainus"`
	Aqicn     int16  `json:"aqicn"`
	Maincn    string `json:"maincn"`
}

type Weather struct {
	Timestamp     string  `json:"ts"`
	Temperature   int16   `json:"tp"`
	AirPressure   int16   `json:"pr"`
	Humidity      int16   `json:"hu"`
	WindSpeed     float32 `json:"ws"`
	WindDirection int16   `json:"wd"`
	Icon          string  `json:"ic"`
}

var airQualityLevels = []string{
	"Baik 🙂",
	"Sedang 😐",
	"Kurang baik 🙁",
	"Tidak baik 😔",
	"Sangat tidak baik 😣",
	"Berbahaya 🤢",
}

var temperatureLevels = []string{
	"Dingin",
	"Sejuk",
	"Hangat",
	"Panas",
	"Sangat panas",
	"Panas ekstrim",
}

func (data CityData) TempLevel() (tl string) {
	temp := data.Weather.Temperature
	if temp > 40 {
		return fmt.Sprintf("%d °C (%s)", temp, temperatureLevels[5])
	}

	if temp > 35 {
		return fmt.Sprintf("%d °C (%s)", temp, temperatureLevels[4])
	}

	if temp > 30 {
		return fmt.Sprintf("%d °C (%s)", temp, temperatureLevels[3])
	}

	if temp > 20 {
		return fmt.Sprintf("%d °C (%s)", temp, temperatureLevels[2])
	}

	if temp > 10 {
		return fmt.Sprintf("%d °C (%s)", temp, temperatureLevels[1])
	}

	return fmt.Sprintf("%d °C (%s)", temp, temperatureLevels[0])
}

func (data CityData) AQL() (aql string) {
	pol := data.Pollution.Aqius
	if pol >= 301 {
		return airQualityLevels[5]
	}

	if pol >= 201 {
		return airQualityLevels[4]
	}

	if pol >= 151 {
		return airQualityLevels[3]
	}

	if pol >= 101 {
		return airQualityLevels[2]
	}

	if pol >= 51 {
		return airQualityLevels[1]
	}

	return airQualityLevels[0]
}
