package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

	commands "tgmarkovbreadgo/commands"
	db "tgmarkovbreadgo/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	token = "5157281753:AAEIKTXt_3k5_upTjdlOSFJ15tc57mkHr6o"
	// token    = "624086120:AAHfg1F2At9lgfLcGdZv-vVIuUEPh91wK3Y"
	dbApi    *db.Api
	bot      *tgbotapi.BotAPI
	command  commands.CommandList
	rg, _    = regexp.Compile(`(\s+)`)
	replacer = strings.NewReplacer("\n", " ")
	lastMsg  = make(map[int64]int64)
)

func main() {
	_bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot = _bot

	log.Println("Authorized at @" + bot.Self.UserName)

	dbApi = db.Init()
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
		if dbApi.Count(id) < 7000 {
			dbApi.AddMsg(id, txt)
		}
	}
}

func tryToGen(update tgbotapi.Update) {
	if update.Message.NewChatMembers == nil {
		rnd := rand.Intn(19)
		rand.Seed(time.Now().Unix())
		rndsec := rand.Intn(19)
		if rnd == rndsec {
			command["gen"].Func(update)
		}
	}
}

func spamErr(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.FromChat().ID, "")
	msg.Text = fmt.Sprintf("⌛️ Подожди %d сек...", 10-(time.Now().Unix()-lastMsg[update.Message.From.ID]))
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func handle(update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.Chat.IsPrivate() == false && update.Message.Time().Add(time.Minute*15).Unix() > time.Now().Unix() {
			if update.Message.IsCommand() && strings.HasSuffix(update.Message.CommandWithAt(), "@"+bot.Self.UserName) {
				if cmd, ok := command[update.Message.Command()]; ok {
					if lastMsg[update.Message.From.ID]+9 < time.Now().Unix() {
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

			if update.Message.IsCommand() == false {
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
	}
}
