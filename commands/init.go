package commands

import (
	"log"

	db "tgmarkovbreadgo/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	Command struct {
		Help  string
		Usage string
		Func  func(update tgbotapi.Update)
		Init  func()
	}
	CommandList map[string]Command
)

var (
	bot      *tgbotapi.BotAPI
	dbApi    *db.Api
	commands = make(CommandList)
)

func Init(b *tgbotapi.BotAPI, db *db.Api) CommandList {
	bot = b
	dbApi = db

	// cmds
	commands["start"] = startCmd
	commands["help"] = helpCmd
	commands["gen"] = genCmd

	for cmdName, cmd := range commands {
		if cmd.Init != nil {
			cmd.Init()
			cmd.Init = nil
		}
		log.Println("command", "'"+cmdName+"'", "initialized!")
	}

	return commands
}
