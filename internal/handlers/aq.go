package handlers

import (
	"fmt"

	"github.com/diazharizky/teleforecaster/internal/app"
	tb "gopkg.in/telebot.v3"
)

func aq(appCtx app.Ctx) func(ctx tb.Context) error {
	return func(ctx tb.Context) error {
		data, err := appCtx.GetAirQualityByCityModule.Call("Indonesia", "West Java", "Bandung")
		if err != nil {
			return ctx.Send("Error has happened")
		}

		template := `
Kota: %s
Provinsi: %s
Suhu: %d
Kualitas Udara: %s
		`

		msg := fmt.Sprintf(
			template,
			data.City,
			data.State,
			data.Weather.Temperature,
			data.AirQualityLevel(),
		)

		return ctx.Send(msg)
	}
}
