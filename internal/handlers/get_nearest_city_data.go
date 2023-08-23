package handlers

import (
	"github.com/diazharizky/teleforecaster/internal/app"
	tb "gopkg.in/telebot.v3"
)

func getNearestCityData(appCtx app.Ctx) func(ctx tb.Context) error {
	return func(ctx tb.Context) error {
		loc := ctx.Message().Location

		data, err := appCtx.GetDataByLocationModule.Call(loc.Lat, loc.Lng)
		if err != nil {
			return ctx.Send("Terjadi kesalahan.")
		}

		return ctx.Send(data.Report())
	}
}
