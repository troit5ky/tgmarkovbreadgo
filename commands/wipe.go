package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var wipeCmd = Command{
	Help:  "accept",
	Usage: "accept",
	Func:  wipe,
}

func wipe(update tgbotapi.Update) {
	id := update.Message.Chat.ID
	from := update.Message.From.ID
	msg := tgbotapi.NewMessage(id, "")

	if update.Message.CommandArguments() != "accept" {
		msg.Text = "Чтобы подтвердить очистку БД, введи - /wipe accept"
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
		dbApi.Wipe(id)
		msg.Text = "🧹 База данных очищена!"
	} else {
		msg.Text = "😢 Ты не админ"
	}

	bot.Send(msg)
}
