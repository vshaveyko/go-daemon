package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbURL = os.Getenv("DATABASE_URL")

type queryFunction func(db gorm.DB)

func ExecQuery(q queryFunction) {
	db, err := gorm.Open("postgres", dbURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	q(*db)

}
