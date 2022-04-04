package generate

import (
	"errors"
	db "tgmarkovbreadgo/database"

	markov "cpl.li/go/markov"
)

func Generate(db *db.Api, id int64) (string, error) {
	var result string
	var err error

	defer func() {
		if rec := recover(); rec != nil {
			err = errors.New("🧐 Мало данных для генерации")
			result = err.Error()
		}
	}()

	ch := markov.NewChain(2)
	ch.Add(db.GetMessages(id))

	b := ch.NewBuilder(nil)
	b.Generate(100 - ch.PairSize)
	result = b.String()

	return result, err
}
