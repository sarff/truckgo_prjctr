package main

import (
	"fmt"
	"github.com/alexandear/truckgo/migrations"
	"github.com/alexandear/truckgo/pkg/config"
	"github.com/alexandear/truckgo/pkg/logging"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"os"
)

func main() {
	config.InitConfig()
	log, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	//Example for log use
	log.Info("Info message", "some key", "some value", "some key2", "some value2")
	log.Debug("Debug message", "some key", "some value", "some key2", "some value2")
	log.Warn("Warn message", "some key", "some value", "some key2", "some value2")
	log.Error("Error message", "some key", "some value", "some key2", "some value2")
	//log.Fatal("Fatal message", "some key", "some value", "some key2", "some value2") // завершає виконання програми

	database := os.Getenv("database")

	//  use sqlite for weak computers where there is no way to run docker with postgresql
	if database != "postgresql" && database != "sqlite" {
		log.Fatal("Database type: " + database + " is not valid. Expecting `mysql` or `postgresql`")
	}

	var db *gorm.DB
	if database == "sqlite" {
		dbFile := viper.GetString("sqlite_path")
		db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			log.Fatal("Failed to open sqlite: ", dbFile)
		}
		log.Info("Successfully connected to sqlite")
	} else {
		user := viper.GetString("POSTGRES_USER")
		pass := viper.GetString("POSTGRES_PASSWORD")
		dbname := viper.GetString("POSTGRES_DB")

		dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kyiv", user, pass, dbname, "8888")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to postgresql", err)
		}
		log.Info("Successfully connected to postgresql")
	}

	err = migrations.Migrate(db)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

}
