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

	if len(arg) > 1 {
		if cmd, isValid := commands[arg]; isValid {
			text := fmt.Sprintf("📚 Команда /%s\n\nИспользование:\n /%s %s\n\nПример:\n /%s %s", arg, arg, cmd.Help, arg, cmd.Usage)
			msg := tgbotapi.NewMessage(update.FromChat().ID, text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		return
	}

	for cmd, _ := range commands {
		fmt.Fprintf(cmds, "/%s\n", cmd)
	}

	text := fmt.Sprintf("📚 Список команд:\n%s\nПодробнее: /help команда | /help weather", cmds.String())
	msg := tgbotapi.NewMessage(update.FromChat().ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
