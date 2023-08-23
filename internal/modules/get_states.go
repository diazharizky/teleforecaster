package modules

import (
	"github.com/diazharizky/teleforecaster/internal/app"
)

type getStatesModule struct {
	appCtx app.Ctx
}

func NewGetStatesModule(appCtx app.Ctx) getStatesModule {
	return getStatesModule{appCtx}
}

func (m getStatesModule) Call(country string) ([]string, error) {
	states, err := m.appCtx.StateRepository.List(country)
	if err != nil {
		return nil, err
	}

	if len(states) <= 0 {
		data, err := m.appCtx.AirVisualClient.GetStates(country)
		if err != nil {
			return nil, err
		}

		states = make([]string, len(data))
		for i, d := range data {
			states[i] = d.State
		}

		if err = m.appCtx.StateRepository.Save(country, states); err != nil {
			return nil, err
		}
	}

	return states, nil
}
