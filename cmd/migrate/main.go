package main

import (
	"log"
	"time"

	"github.com/suryaadi44/iris-playground/app/entity"
	"github.com/suryaadi44/iris-playground/utils/config"
	"github.com/suryaadi44/iris-playground/utils/database/postgresql"
)

func main() {
	config, errr := config.Load("config")
	if errr != nil {
		panic(errr)
	}

	local, err := time.LoadLocation(config.GetString("timezone"))
	if err != nil {
		log.Fatalf("[Config] Error loading timezone: %v", err)
	}
	time.Local = local

	db := postgresql.InitDatabase(
		config.GetString("database.postgres.host"),
		config.GetInt("database.postgres.port"),
		config.GetString("database.postgres.user"),
		config.GetString("database.postgres.password"),
		config.GetString("database.postgres.database"),
		config.GetString("timezone"),
	)

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
}
