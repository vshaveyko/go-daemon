package process

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/vshaveyko/test-go-daemon/db"
)

func getRequiredProcessing() (result []Pipeline) {

	queryFunction := func(conn gorm.DB) {
		currentHour := time.Now().Hour()

		conn.Where(
			"hour = ?", currentHour,
		).Preload(
			"Connectors",
		).Preload(
			"Connectors.Source",
		).Preload(
			"Connectors.Target",
		).Find(&result)
	}

	db.ExecQuery(queryFunction)

	return

}
