package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var statCmd = Command{
	Help:  "",
	Usage: "",
	Func:  stat,
}

func stat(update tgbotapi.Update) {
	id := update.Message.Chat.ID
	msg := tgbotapi.NewMessage(id, "")
	count := dbApi.Count(id)

	msg.Text = fmt.Sprintf("🔎 Предложений в базе данных %d\n🧐 Если эта цифра не меняется, то убедитесь, что бот имеет доступ к сообщениям (админка)", count)
	bot.Send(msg)
}
