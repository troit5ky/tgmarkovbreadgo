package main

import (
	"log"
	"math/rand"
	"regexp"
	"strings"

	commands "tgmarkovbreadgo/commands"
	db "tgmarkovbreadgo/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	token    = "5157281753:AAEIKTXt_3k5_upTjdlOSFJ15tc57mkHr6o"
	dbApi    *db.Api
	command  commands.CommandList
	rg, _    = regexp.Compile(`(\s+)`)
	replacer = strings.NewReplacer("\n", " ")
)

func main() {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Authorized at " + bot.Self.UserName)

	dbApi = db.Init()
	command = commands.Init(bot, dbApi)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		go handle(update)
	}
}

func addMsg(id int64, txt string) {
	txt = replacer.Replace(txt)

	if count := rg.FindAllString(txt, -1); len(count) > 0 {
		dbApi.AddMsg(id, txt)
	}
}

func tryToGen(update tgbotapi.Update) {
	rnd := rand.Intn(8)
	if rnd == 3 {
		command["gen"].Func(update)
	}
}

func handle(update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Chat.IsPrivate() == false {
			if update.Message.IsCommand() {
				if cmd, ok := command[update.Message.Command()]; ok {
					cmd.Func(update)
				}
			} else if update.Message.Text != "" {
				addMsg(update.FromChat().ID, update.Message.Text)
			} else if update.Message.Caption != "" {
				addMsg(update.FromChat().ID, update.Message.Caption)
			}

			if update.Message.IsCommand() == false {
				tryToGen(update)
			}
		}
	}
}
