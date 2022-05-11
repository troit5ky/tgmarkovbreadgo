package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GID      int64
	Messages []Message `gorm:"foreignKey:GroupID; references:GID"`
}

type Message struct {
	gorm.Model
	Text    string
	GroupID int64
}

var (
	db *gorm.DB
)

func Init() *Api {
	_db, err := gorm.Open(sqlite.Open("bot.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	_db.AutoMigrate(&Group{})
	_db.AutoMigrate(&Message{})

	db = _db
	_db = nil

	return new(Api)
}
