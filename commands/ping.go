package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var pingCmd = Command{
	Help:    "",
	Usage: "",
	Func:    ping,
}

func ping(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "🏓 pong")
	msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
