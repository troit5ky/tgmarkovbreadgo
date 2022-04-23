package generate

import (
	"math/rand"
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

	arr := db.GetMessages(id)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	ch.Add(arr)

	b := ch.NewBuilder(nil)
	b.Generate(30 - ch.PairSize)
	result = b.String()

	return result
}
