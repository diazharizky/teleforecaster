package handlers

import (
	"fmt"

	"github.com/diazharizky/teleforecaster/internal/app"
	tb "gopkg.in/telebot.v3"
)

func getNearestCityData(appCtx app.Ctx) func(ctx tb.Context) error {
	return func(ctx tb.Context) error {
		loc := ctx.Message().Location

		data, err := appCtx.GetDataByLocationModule.Call(loc.Lat, loc.Lng)
		if err != nil {
			return ctx.Send("Error has happened")
		}

		template := "Kota/Kabupaten: %s\nProvinsi: %s\nSuhu: %s\nKualitas Udara: %s"
		msg := fmt.Sprintf(
			template,
			data.City,
			data.State,
			data.TempLevel(),
			data.AQL(),
		)

		return ctx.Send(msg)
	}
}
