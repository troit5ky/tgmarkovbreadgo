package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var resetCmd = Command{
	Help:  "accept",
	Usage: "accept",
	Func:  reset,
}

func reset(update tgbotapi.Update) {
	id := update.Message.Chat.ID
	from := update.Message.From.ID
	msg := tgbotapi.NewMessage(id, "")

	if update.Message.CommandArguments() != "accept" {
		msg.Text = "Чтобы подтвердить очистку БД, введи - /reset accept"
		bot.Send(msg)
		return
	}

	admins, _ := bot.GetChatAdministrators(tgbotapi.ChatAdministratorsConfig{
		ChatConfig: update.Message.Chat.ChatConfig(),
	})

	for _, admin := range admins {
		if admin.User.ID == from {
			dbApi.Reset(id)
			msg.Text = "🧹 База данных очищена!"
			bot.Send(msg)
			return
		}
	}

	msg.Text = "😢 Ты не админ"
	bot.Send(msg)
}
