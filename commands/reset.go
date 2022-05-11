package commands

import (
	"fmt"

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

	member, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{
		ChatConfigWithUser: tgbotapi.ChatConfigWithUser{
			ChatID:             id,
			SuperGroupUsername: update.Message.Chat.UserName,
			UserID:             from,
		},
	})
	if err != nil {
		msg.Text = fmt.Sprintf("произошла ошибка '%s'", err)
		bot.Send(msg)
		return
	}

	if member.IsAdministrator() || member.IsCreator() {
		dbApi.Reset(id)
		msg.Text = "🧹 База данных очищена!"
	} else {
		msg.Text = "😢 Ты не админ"
	}

	bot.Send(msg)
}
