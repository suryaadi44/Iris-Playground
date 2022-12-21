package main

import (
	"log"
	"suryaadi44/iris-playground/app/entity"
	"suryaadi44/iris-playground/utils/config"
	"suryaadi44/iris-playground/utils/database/sql"
	"time"
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

	db := sql.InitDatabase(
		config.GetString("database.host"),
		config.GetInt("database.port"),
		config.GetString("database.user"),
		config.GetString("database.password"),
		config.GetString("database.database"),
		config.GetString("timezone"),
	)

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
}
