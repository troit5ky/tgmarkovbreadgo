package database

import "strings"

type Api struct{}

func (Api) AddMsg(id int64, txt string) {
	txt = strings.ToLower(txt)

	var gr *Group
	var msg *Message

	db.Model(&gr).Where("g_id = ?", id).First(&gr)
	if gr.ID == 0 {
		gr.GID = id
		db.Create(&gr)
	}

	db.Find(&msg, "text = ? AND group_id = ?", txt, id)
	if msg.ID == 0 {
		db.Create(&Message{
			Text:    txt,
			GroupID: id,
		})
	}
}

func (Api) GetMessages(id int64) []string {
	var msgs []string

	db.Model(&Message{}).Select("text").Where("group_id = ?", id).Find(&msgs)

	return msgs
}

func (Api) Count(id int64) int64 {
	var count int64

	db.Model(&Message{}).Where("group_id = ?", id).Count(&count)

	return count
}

func (Api) Wipe(id int64) {
	db.Model(&Message{}).Where("group_id = ?", id).Unscoped().Delete(&Message{})
}
