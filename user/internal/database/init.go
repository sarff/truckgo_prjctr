package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	var err error
	databaseType := viper.GetString("db_type")
	err = godotenv.Load(viper.GetString("env_path"))
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	//  use sqlite for weak computers where there is no way to run docker with postgresql
	if databaseType != "postgresql" && databaseType != "sqlite" {
		return nil, fmt.Errorf("database type: %s is not valid. Expecting `sqlite` or `postgresql`", databaseType)
	}

	var db *gorm.DB

	if databaseType == "sqlite" {
		dbFile := viper.GetString("sqlite_path")
		db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to open sqlite: %s", dbFile)
		}
	} else {
		user := os.Getenv("POSTGRES_USER")
		pass := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB_USER")

		// NOTE inside docker always 5432
		dsn := fmt.Sprintf("host=db_user user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Kyiv",
			user, pass, dbname)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to postgresql %s", err)
		}
	}

	return db, nil
}
