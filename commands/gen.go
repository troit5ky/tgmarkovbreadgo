package commands

import (
	markov "tgmarkovbreadgo/generate"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var genCmd = Command{
	Help:  "",
	Usage: "",
	Func:  gen,
}

func gen(update tgbotapi.Update) {
	id := update.Message.Chat.ID
	result, _ := markov.Generate(dbApi, id)

	msg := tgbotapi.NewMessage(id, result)
	bot.Send(msg)
}
