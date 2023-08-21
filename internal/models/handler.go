package models

import (
	tb "gopkg.in/telebot.v3"
)

type Handler struct {
	Endpoint    string
	Description string
	Fn          func(tb.Context) error
}
