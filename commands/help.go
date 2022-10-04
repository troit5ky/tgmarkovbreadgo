package commands

import (
	"bytes"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var helpCmd = Command{
	Help:  "[команда]",
	Usage: "weather",
	Func:  help,
}

func help(update tgbotapi.Update) {
	arg := strings.ToLower(update.Message.CommandArguments())
	cmds := new(bytes.Buffer)
	msg := tgbotapi.NewMessage(update.FromChat().ID, "")
	msg.ReplyToMessageID = update.Message.MessageID

	if len(arg) > 0 {
		if cmd, isValid := commands[arg]; isValid {
			msg.Text = fmt.Sprintf("📚 Команда /%s\n\nИспользование:\n /%s %s\n\nПример:\n /%s %s", arg, arg, cmd.Help, arg, cmd.Usage)
			bot.Send(msg)
		} else {
			msg.Text = "🤔 Я не нашёл такой команды"
			bot.Send(msg)
		}
		return
	}

	for cmd := range commands {
		fmt.Fprintf(cmds, "/%s\n", cmd)
	}

	msg.Text = fmt.Sprintf("📚 Список команд:\n%s\nПодробнее: /help команда | /help gen", cmds.String())
	bot.Send(msg)
}
