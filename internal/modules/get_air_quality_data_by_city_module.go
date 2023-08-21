package modules

import (
	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
)

type getAirQualityDataByCityModule struct {
	appCtx app.Ctx
}

func NewGetAirQualityDataByCityModule(appCtx app.Ctx) getAirQualityDataByCityModule {
	return getAirQualityDataByCityModule{appCtx}
}

func (m getAirQualityDataByCityModule) Call(country, state, city string) (*airvisual.CityData, error) {
	data, err := m.appCtx.AirVisualClient.GetCityData("Bandung", "West Java", "Indonesia")
	if err != nil {
		return nil, err
	}

	return data, nil
}
