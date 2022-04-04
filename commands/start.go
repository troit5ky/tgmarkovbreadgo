package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var startCmd = Command{
	Help:  "",
	Usage: "",
	Func:  start,
}

func start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "👋 Запущен")
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
