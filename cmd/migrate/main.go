package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/viper"
	"github.com/suryaadi44/iris-playground/app/entity"
	"github.com/suryaadi44/iris-playground/utils/config"
	"github.com/suryaadi44/iris-playground/utils/database/postgresql"
	"gorm.io/gorm"
)

const (
	TargetDB             = "target"
	DevDB                = "target_dev"
	SchemaFilePermission = 0644
)

func main() {
	name := flag.String("name", "", "Migration name")
	flag.Parse()

	if *name == "" {
		log.Fatalf("[Migrate] Migration name is required")
	}

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

	err = createEmptyDatabase(db, TargetDB, DevDB)
	if err != nil {
		log.Fatalf("[Migrate] Error creating empty database: %v", err)
	}

	err = migrateDatabase(TargetDB, config)
	if err != nil {
		log.Fatalf("[Migrate] Error migrating database: %v", err)
	}

	err = inspectSchema(TargetDB, config)
	if err != nil {
		log.Fatalf("[Migrate] Error inspecting database: %v", err)
	}

	err = generateMigrationFile(DevDB, *name, config)
	if err != nil {
		log.Fatalf("[Migrate] Error generating migration file: %v", err)
	}

	err = dropDatabase(db, TargetDB, DevDB)
	if err != nil {
		log.Fatalf("[Migrate] Error dropping database: %v", err)
	}
}

func createEmptyDatabase(db *gorm.DB, names ...string) error {
	for _, n := range names {
		err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", n)).Error
		if err != nil {
			return err
		}

		err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", n)).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func dropDatabase(db *gorm.DB, names ...string) error {
	for _, n := range names {
		err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", n)).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func migrateDatabase(targetDb string, conf *viper.Viper) error {
	db := postgresql.InitDatabase(
		conf.GetString("database.postgres.host"),
		conf.GetInt("database.postgres.port"),
		conf.GetString("database.postgres.user"),
		conf.GetString("database.postgres.password"),
		targetDb,
		conf.GetString("timezone"),
	)

	// TODO: Add model to migrate here
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func inspectSchema(targetDb string, conf *viper.Viper) error {
	// call atlas binary to get schema
	// atlas schema inspect -u  "postgres://user:password@host:port/target_db?sslmode=disable" --exclude '*.schema_migrations' > schema.hcl

	dbStr := fmt.Sprintf(
		"\"postgres://%s:%s@%s:%d/%s?sslmode=disable\"",
		conf.GetString("database.postgres.user"),
		conf.GetString("database.postgres.password"),
		conf.GetString("database.postgres.host"),
		conf.GetInt("database.postgres.port"),
		targetDb,
	)
	inspectStr := fmt.Sprintf("atlas schema inspect -u %s --exclude \"*.schema_migrations\"", dbStr)
	output, err := exec.Command("sh", "-c", inspectStr).Output()
	if err != nil {
		return err
	}

	err = os.WriteFile("schema.hcl", output, SchemaFilePermission)
	if err != nil {
		return err
	}

	return nil
}

func generateMigrationFile(devDb string, name string, conf *viper.Viper) error {
	// call atlas binary to generate migration file
	// atlas migrate diff name --dir "file://migrations?format=golang-migrate"  --dev-url "postgres://user:password@host:port/target_db?sslmode=disable" --to "file://schema.hcl"

	devStr := fmt.Sprintf(
		"\"postgres://%s:%s@%s:%d/%s?sslmode=disable\"",
		conf.GetString("database.postgres.user"),
		conf.GetString("database.postgres.password"),
		conf.GetString("database.postgres.host"),
		conf.GetInt("database.postgres.port"),
		devDb,
	)
	diffStr := fmt.Sprintf(
		"atlas migrate diff %s --dir \"file://migrations?format=golang-migrate\" --dev-url %s --to \"file://schema.hcl\"",
		name,
		devStr,
	)
	output, err := exec.Command("sh", "-c", diffStr).Output()
	if err != nil {
		log.Printf("[Migrate] Error generating migration file: %s", output)
		return err
	}

	return nil
}
