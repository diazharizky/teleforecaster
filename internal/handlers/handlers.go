package handlers

import (
	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/internal/models"
	"gopkg.in/telebot.v3"
)

type session struct {
	state string
	city  string
}

const country = "Indonesia"

var userSession = map[int64]session{}

func Init(appCtx app.Ctx) (h []models.Handler) {
	return []models.Handler{
		{
			Endpoint:    telebot.OnLocation,
			Description: "Get air quality data of the nearest supported city",
			Fn:          getNearestCityData(appCtx),
		},
		{
			Endpoint:    "/aq",
			Description: "Get air quality data based on user input",
			Fn:          aq(appCtx),
		},
		{
			Endpoint:    telebot.OnText,
			Description: "Part part of `/aq` handler",
			Fn:          aqStepper(appCtx),
		},
	}
}
