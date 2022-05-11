package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var startCmd = Command{
	Help:  "",
	Usage: "",
	Func:  start,
}

func start(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "👋 Привет, я умею генерировать сообщения, изучая ваши разговоры.\n📚 Посмотреть команды можно в /help\n\nP.S. я отправляю сообщения сам, но их можно сгенерироавть с помощью /gen")

	bot.Send(msg)
}
