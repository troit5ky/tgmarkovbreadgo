package generate

import (
	"math/rand"

	markov "cpl.li/go/markov"
)

func Generate(id int64) string {
	var result string

	defer func() {
		if rec := recover(); rec != nil {
			return
		}
	}()

	ch := markov.NewChain(2)

	arr := dbApi.GetMessages(id)
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	ch.Add(arr)

	b := ch.NewBuilder(nil)
	b.Generate(30 - ch.PairSize)
	result = b.String()

	return result
}
