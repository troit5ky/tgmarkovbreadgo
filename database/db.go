package database

type Api struct{}

func (Api) AddMsg(id int64, txt string) {
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
	var msgsRaw []Message
	var msgs []string

	db.Where("group_id = ?", id).Find(&msgsRaw)
	for _, raw := range msgsRaw {
		msgs = append(msgs, raw.Text)
	}

	return msgs
}
