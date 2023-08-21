package airvisual

import (
	"encoding/json"
	"fmt"
)

type BaseResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

var airQualityLevels = []string{
	"Baik",
	"Sedang",
	"Kurang baik",
	"Tidak baik",
	"Sangat tidak baik",
	"Berbahaya",
}

func (r BaseResponse) Decode(dest interface{}) error {
	jsonStr, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonStr, dest); err != nil {
		return err
	}

	return nil
}

type CityData struct {
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Location `json:"location"`
	Current  `json:"current"`
}

func (data CityData) Print() {
	fmt.Println("City", data.City)
	fmt.Println("State", data.State)
	fmt.Println("Country", data.Country)

	fmt.Println("Location.Type:", data.Location.Type)
	fmt.Println("Location.Coordinates:", data.Location.Coordinates)

	fmt.Println("Pollution.Timestamp:", data.Current.Pollution.Timestamp)
	fmt.Println("Pollution.Aqius:", data.Current.Pollution.Aqius)
	fmt.Println("Pollution.Mainus:", data.Current.Pollution.Mainus)
	fmt.Println("Pollution.Aqicn:", data.Current.Pollution.Aqicn)
	fmt.Println("Pollution.Maincn:", data.Current.Pollution.Maincn)

	fmt.Println("Weather.Timestamp:", data.Current.Weather.Timestamp)
	fmt.Println("Weather.Temperature:", data.Current.Weather.Temperature)
	fmt.Println("Weather.AirPressure:", data.Current.Weather.AirPressure)
	fmt.Println("Weather.Humidity:", data.Current.Weather.Humidity)
	fmt.Println("Weather.WindSpeed:", data.Current.Weather.WindSpeed)
	fmt.Println("Weather.WindDirection:", data.Current.Weather.WindDirection)
	fmt.Println("Weather.Icon:", data.Current.Weather.Icon)
}

func (data CityData) AirQualityLevel() (aql string) {
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
