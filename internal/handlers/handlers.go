package handlers

import (
	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/internal/models"
)

func Init(appCtx app.Ctx) (h []models.Handler) {
	return []models.Handler{
		{
			Endpoint:    "/aq",
			Description: "Get air quality data",
			Fn:          aq(appCtx),
		},
	}
}
