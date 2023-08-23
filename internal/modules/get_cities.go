package modules

import (
	"github.com/diazharizky/teleforecaster/internal/app"
)

type getCitiesModule struct {
	appCtx app.Ctx
}

func NewGetCitiesModule(appCtx app.Ctx) getCitiesModule {
	return getCitiesModule{appCtx}
}

func (m getCitiesModule) Call(country, state string) ([]string, error) {
	cities := m.appCtx.CityRepository.List(country, state)
	if len(cities) <= 0 {
		data, err := m.appCtx.AirVisualClient.GetCities(country, state)
		if err != nil {
			return nil, err
		}

		if len(data) <= 0 {
			return []string{}, nil
		}

		cities = make([]string, len(data))
		for i, d := range data {
			cities[i] = d.City
		}

		m.appCtx.CityRepository.Save(country, state, cities)
	}

	return cities, nil
}
