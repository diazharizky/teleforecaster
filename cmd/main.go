package main

import (
	"fmt"
	"log"
	"time"

	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/internal/handlers"
	"github.com/diazharizky/teleforecaster/internal/modules"
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
	tb "gopkg.in/telebot.v3"
)

var appCtx app.Ctx = app.Ctx{}
var bot *tb.Bot

func init() {
	appCtx.AirVisualClient = airvisual.New()
	appCtx.GetAirQualityByCityModule = modules.NewGetAirQualityDataByCityModule(appCtx)

	initBot()
	initHandlers()
}

func main() {
	bot.Start()
}

func initBot() {
	settings := tb.Settings{
		Token:  "6694802209:AAGAb0YIVHu6AqIiUzyr2E6rn7YGCcF4bvM",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	bot, err = tb.NewBot(settings)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func initHandlers() {
	fmt.Print("List of available endpoints:\n\n")

	hs := handlers.Init(appCtx)
	for _, h := range hs {
		bot.Handle(h.Endpoint, h.Fn)
		fmt.Printf("%s: %s\n", h.Endpoint, h.Description)
	}
}
