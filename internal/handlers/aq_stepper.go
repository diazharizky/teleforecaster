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

		template := "Kota/Kabupaten: %s\nProvinsi: %s\nSuhu: %s\nKualitas Udara: %s"
		msg := fmt.Sprintf(
			template,
			resp.City,
			resp.State,
			resp.TempLevel(),
			resp.AQL(),
		)

		return ctx.Send(msg)
	}
}
