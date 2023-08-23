package handlers

import (
	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/internal/models"
	"gopkg.in/telebot.v3"
)

func Init(appCtx app.Ctx) (h []models.Handler) {
	return []models.Handler{
		{
			Endpoint:    telebot.OnLocation,
			Description: "Get air quality data of the nearest supported city",
			Fn:          getNearestCityData(appCtx),
		},
		{
			Endpoint:    "/aq",
			Description: "Get air quality data",
			Fn:          aq(appCtx),
		},
		{
			Endpoint: telebot.OnText,
			Fn:       proc(appCtx),
		},
	}
}
