package generate

import (
	"strings"

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

	for _, sentence := range dbApi.GetMessages(id) {
		ch.Add(strings.Fields(sentence))
	}

	b := ch.NewBuilder(nil)
	b.Generate(13 - ch.PairSize)
	result = b.String()

	return result
}
