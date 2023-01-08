package postgresql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(host string, port int, user string, password string, dbname string, timezone string) *gorm.DB {
	pConf := postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			host,
			port,
			user,
			password,
			dbname,
			timezone,
		),
	}

	gConf := &gorm.Config{
		SkipDefaultTransaction: false,
	}

	db, err := gorm.Open(postgres.New(pConf), gConf)
	if err != nil {
		log.Fatalf("[DB] Error connecting to database: %v", err)
	}

	log.Println("[DB] Connected to database")

	return db
}
