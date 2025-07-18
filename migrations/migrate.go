package migrations

import (
	"log"

	"github.com/mdzahid786/golang-graphql/db"
	"github.com/mdzahid786/golang-graphql/graph/model"
)

func Migrate() {
	db.InitDB()
	err:= db.DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}
}