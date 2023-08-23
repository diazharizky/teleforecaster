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
	states := m.appCtx.StateRepository.List(country)
	if len(states) <= 0 {
		data, err := m.appCtx.AirVisualClient.GetStates(country)
		if err != nil {
			return nil, err
		}

		states = make([]string, len(data))
		for i, d := range data {
			states[i] = d.State
		}

		m.appCtx.StateRepository.Save(country, states)
	}

	return states, nil
}
