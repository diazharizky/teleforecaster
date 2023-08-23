package handlers

import (
	"fmt"
	"strconv"

	"github.com/diazharizky/teleforecaster/internal/app"
	tb "gopkg.in/telebot.v3"
)

const invalidInputMsg = "Input tidak sesuai."

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
				return ctx.Send(
					resolveErrMessage(err),
				)
			}

			stateInt, _ := strconv.Atoi(txt)
			stateInt--

			if invalidInput(len(states), stateInt) {
				delete(userSession, senderID)
				return ctx.Send(invalidInputMsg)
			}

			session.state = states[stateInt]
			userSession[senderID] = session

			cities, err := appCtx.GetCitiesModule.Call(country, session.state)
			if err != nil {
				delete(userSession, senderID)
				return ctx.Send(
					resolveErrMessage(err),
				)
			}

			pickCityMsg := "Pilih kota/kabupaten:\n"
			for i, city := range cities {
				pickCityMsg += fmt.Sprintf("%d. %s\n", i+1, city)
			}

			return ctx.Send(pickCityMsg)
		}
		defer func() {
			delete(userSession, senderID)
		}()

		if session.city == "" {
			cities, err := appCtx.GetCitiesModule.Call(country, session.state)
			if err != nil {
				return ctx.Send(
					resolveErrMessage(err),
				)
			}

			cityInt, _ := strconv.Atoi(txt)
			cityInt--

			if invalidInput(len(cities), cityInt) {
				return ctx.Send(invalidInputMsg)
			}

			session.city = cities[cityInt]
			userSession[senderID] = session // TODO: This may unnecessary, will remove it later
		}

		data, err := appCtx.AirVisualClient.GetDataByCity(country, session.state, session.city)
		if err != nil {
			return ctx.Send(
				resolveErrMessage(err),
			)
		}

		return ctx.Send(
			data.Report(),
		)
	}
}

func invalidInput(optsLen int, input int) bool {
	return input < 0 || input > optsLen-1
}
