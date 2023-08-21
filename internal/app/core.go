package app

import (
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
)

type Ctx struct {
	AirVisualClient airvisual.Client

	GetAirQualityByCityModule IGetAirQualityDataByCityModule
}
