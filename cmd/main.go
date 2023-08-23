package main

import (
	"log"
	"time"

	"github.com/diazharizky/teleforecaster/internal/app"
	"github.com/diazharizky/teleforecaster/internal/handlers"
	"github.com/diazharizky/teleforecaster/internal/modules"
	"github.com/diazharizky/teleforecaster/internal/repositories"
	"github.com/diazharizky/teleforecaster/pkg/airvisual"
	"github.com/diazharizky/teleforecaster/pkg/cache"
	"gopkg.in/telebot.v3"
)

var appCtx app.Ctx = app.Ctx{}
var bot *telebot.Bot

func init() {
	appCtx.AirVisualClient = airvisual.New()

	cacheStore := cache.New("localhost", "6379", "", 0)

	appCtx.StateRepository = repositories.NewStateRepository(cacheStore)
	appCtx.CityRepository = repositories.NewCityRepository(cacheStore)

	appCtx.GetDataByLocationModule = modules.NewGetDataByLocationModule(appCtx)
	appCtx.GetStatesModule = modules.NewGetStatesModule(appCtx)
	appCtx.GetCitiesModule = modules.NewGetCitiesModule(appCtx)

	initBot()
	initHandlers()
}

func main() {
	log.Println("Bot is listening...")

	bot.Start()
}

func initBot() {
	settings := telebot.Settings{
		Token:  "6694802209:AAGAb0YIVHu6AqIiUzyr2E6rn7YGCcF4bvM",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	bot, err = telebot.NewBot(settings)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func initHandlers() {
	hs := handlers.Init(appCtx)
	for _, h := range hs {
		bot.Handle(h.Endpoint, h.Fn)
	}
}
