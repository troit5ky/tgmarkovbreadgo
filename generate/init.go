package generate

import (
	"tgmarkovbreadgo/database"
)

var (
	dbApi *database.Api
)

func Init(_db *database.Api) {
	dbApi = _db
}
