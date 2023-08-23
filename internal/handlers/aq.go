package handlers

import (
	"fmt"

	"github.com/diazharizky/teleforecaster/internal/app"
	tb "gopkg.in/telebot.v3"
)

func aq(appCtx app.Ctx) func(ctx tb.Context) error {
	return func(ctx tb.Context) error {
		senderID := ctx.Sender().ID
		_, ok := userSession[senderID]
		if !ok {
			userSession[senderID] = session{}
		}

		states, err := appCtx.GetStatesModule.Call(country)
		if err != nil {
			return err
		}

		if len(states) <= 0 {
			return ctx.Send("Gagal mengambil data provinsi.")
		}

		stateMsg := "Pilih provinsi:\n"
		for i, s := range states {
			stateMsg += fmt.Sprintf("%d. %s\n", i+1, s)
		}

		return ctx.Send(stateMsg)
	}
}
