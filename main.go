package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	commands "tgmarkovbreadgo/commands"
	"tgmarkovbreadgo/config"
	db "tgmarkovbreadgo/database"
	markov "tgmarkovbreadgo/generate"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	dbApi    *db.Api
	bot      *tgbotapi.BotAPI
	command  commands.CommandList
	rg, _    = regexp.Compile(`(\s+)`)
	replacer = strings.NewReplacer("\n", " ")
	lastMsg  = make(map[int64]int64)
)

func main() {
	log.Println("Starting...")

	// CFG init
	if err := config.Init(); err != nil {
		return
	}

	_bot, err := tgbotapi.NewBotAPI(config.Config.Token)
	if err != nil {
		log.Println("Maybe check Bot Token?")
		log.Panic(err)
	}

	bot = _bot

	log.Println("Authorized at @" + bot.Self.UserName)

	dbApi = db.Init()
	markov.Init(dbApi)
	command = commands.Init(bot, dbApi)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		go handle(update)
	}
}

func addMsg(id int64, txt string, ents []tgbotapi.MessageEntity) {
	if len(ents) > 0 {
		return
	}

	txt = replacer.Replace(txt)

	if count := rg.FindAllString(txt, -1); len(count) > 0 {
		if dbApi.Count(id) < config.Config.DatabaseLimit {
			dbApi.AddMsg(id, txt)
		}
	}
}

func tryToGen(update tgbotapi.Update) {
	if update.Message.NewChatMembers == nil && update.Message.ForwardFrom == nil {
		rnd := rand.Intn(10)

		if rnd == 0 {
			result := markov.Generate(update.Message.Chat.ID)
			if result != "" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				msg.Text = result
				bot.Send(msg)
			}
		}
	}
}

func spamErr(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "")
	msg.Text = fmt.Sprintf("⌛️ Подожди %d сек...", (config.Config.Cooldown+1)-(time.Now().Unix()-lastMsg[update.Message.From.ID]))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func handle(update tgbotapi.Update) {
	if update.Message != nil {
		if !update.Message.Chat.IsPrivate() && update.Message.Time().Add(time.Minute*15).Unix() > time.Now().Unix() {
			if update.Message.IsCommand() && strings.HasSuffix(update.Message.CommandWithAt(), "@"+bot.Self.UserName) {
				if cmd, ok := command[update.Message.Command()]; ok {
					if lastMsg[update.Message.From.ID]+config.Config.Cooldown < time.Now().Unix() {
						cmd.Func(update)
						lastMsg[update.Message.From.ID] = time.Now().Unix()
					} else {
						spamErr(update)
					}
				}
			} else if update.Message.Text != "" {
				addMsg(update.FromChat().ID, update.Message.Text, update.Message.Entities)
			} else if update.Message.Caption != "" {
				addMsg(update.FromChat().ID, update.Message.Caption, update.Message.CaptionEntities)
			}

			if !update.Message.IsCommand() {
				tryToGen(update)
			}
		}

		if update.Message.NewChatMembers != nil {
			for _, newMember := range update.Message.NewChatMembers {
				if newMember.UserName == bot.Self.UserName {
					command["start"].Func(update)
				}
			}
		}

		if update.Message.LeftChatMember != nil {
			if update.Message.LeftChatMember.ID == bot.Self.ID {
				dbApi.Wipe(update.Message.Chat.ID)
			}
		}
	}
}
