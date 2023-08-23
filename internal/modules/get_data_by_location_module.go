package modules

import (
	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
)

type getDataByLocationModule struct {
	appCtx app.Ctx
}

func NewGetDataByLocationModule(appCtx app.Ctx) getDataByLocationModule {
	return getDataByLocationModule{appCtx}
}

func (m getDataByLocationModule) Call(lat, lng float32) (*airvisual.CityData, error) {
	data, err := m.appCtx.AirVisualClient.GetDataByLocation(lat, lng)
	if err != nil {
		return nil, err
	}

	return data, nil
}
