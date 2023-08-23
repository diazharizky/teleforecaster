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
	cities, err := m.appCtx.CityRepository.List(country, state)
	if err != nil {
		return nil, err
	}

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

		if err = m.appCtx.CityRepository.Save(country, state, cities); err != nil {
			return nil, err
		}
	}

	return cities, nil
}
