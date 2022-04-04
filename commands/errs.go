package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func notEnoughtArgs(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "🧐 Аргумены неверно заданы!")
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
