package app

import "github.com/diazharizky/teleforecaster/pkg/airvisual"

type IGetAirQualityDataByCityModule interface {
	Call(country, state, city string) (*airvisual.CityData, error)
}
