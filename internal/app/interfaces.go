package app

import "github.com/diazharizky/teleforecaster/pkg/airvisual"

type IGetDataByLocationModule interface {
	Call(lat, lng float32) (*airvisual.CityData, error)
}

type IGetDataByCityModule interface {
	Call(state, city string) (*airvisual.CityData, error)
}

type IGetStatesModule interface {
	Call(country string) ([]string, error)
}

type IGetCitiesModule interface {
	Call(country, state string) ([]string, error)
}

type IStateRepository interface {
	List(country string) ([]string, error)
	Save(country string, states []string) error
}

type ICityRepository interface {
	List(country, state string) ([]string, error)
	Save(country, state string, cities []string) error
}
