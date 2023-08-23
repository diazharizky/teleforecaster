package handlers

import (
	"fmt"
	"strconv"

	"github.com/diazharizky/teleforecaster/internal/app"
	tb "gopkg.in/telebot.v3"
)

func aqStepper(appCtx app.Ctx) func(ctx tb.Context) error {
	return func(ctx tb.Context) error {
		senderID := ctx.Sender().ID
		session, ok := userSession[senderID]
		if !ok {
			return nil
		}

		txt := ctx.Text()

		if session.state == "" {
			states, err := appCtx.GetStatesModule.Call(country)
			if err != nil {
				return err
			}

			stateInt, _ := strconv.Atoi(txt)

			session.state = states[stateInt-1]
			userSession[senderID] = session

			pickCityQuestion := "Pilih kota/kabupaten:\n"

			cities, err := appCtx.GetCitiesModule.Call(country, session.state)
			if err != nil {
				return err
			}

			if len(cities) <= 0 {
				return ctx.Send("Mohon maaf data untuk provinsi tersebut belum tersedia.")
			}

			for i, c := range cities {
				pickCityQuestion += fmt.Sprintf("%d. %s\n", i+1, c)
			}

			return ctx.Send(pickCityQuestion)
		}

		if session.city == "" {
			cities, err := appCtx.GetCitiesModule.Call(country, session.state)
			if err != nil {
				return err
			}

			cityInt, _ := strconv.Atoi(txt)

			session.city = cities[cityInt-1]
			userSession[senderID] = session
		}
		defer func() {
			delete(userSession, senderID)
		}()

		resp, err := appCtx.AirVisualClient.GetDataByCity(country, session.state, session.city)
		if err != nil {
			return err
		}

		return ctx.Send(resp.Report())
	}
}
