package generate

import (
	db "tgmarkovbreadgo/database"

	markov "cpl.li/go/markov"
)

func Generate(db *db.Api, id int64) string {
	var result string

	defer func() {
		if rec := recover(); rec != nil {
			return
		}
	}()

	ch := markov.NewChain(2)
	ch.Add(db.GetMessages(id))

	b := ch.NewBuilder(nil)
	b.Generate(100 - ch.PairSize)
	result = b.String()

	return result
}
