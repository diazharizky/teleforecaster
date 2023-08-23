package app

import (
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
)

type Ctx struct {
	AirVisualClient airvisual.Client

	StateRepository IStateRepository
	CityRepository  ICityRepository

	GetDataByLocationModule IGetDataByLocationModule
	GetStatesModule         IGetStatesModule
	GetCitiesModule         IGetCitiesModule
}
