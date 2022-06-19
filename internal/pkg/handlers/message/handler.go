package message

import tele "gopkg.in/telebot.v3"

type Handler interface {
	Handle(c tele.Context) error
}
