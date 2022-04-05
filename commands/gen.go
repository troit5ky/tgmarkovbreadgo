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
	result := markov.Generate(dbApi, id)
	msg := tgbotapi.NewMessage(id, "")

	if result == "" {
		msg.Text = "🧐 Мало данных для генерации"
		bot.Send(msg)
		return
	}

	msg.Text = result
	bot.Send(msg)
}
