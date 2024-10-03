/*
//Example for log use:

	log.Info("Info message", "some key", "some value", "some key2", "some value2")
	log.Debug("Debug message", "some key", "some value", "some key2", "some value2")
	log.Warn("Warn message", "some key", "some value", "some key2", "some value2")
	log.Error("Error message", "some key", "some value", "some key2", "some value2")
*/
package main

import (
	"fmt"
	"github.com/alexandear/truckgo/internal/pkg/config"
	"github.com/alexandear/truckgo/internal/pkg/logging"
	"github.com/alexandear/truckgo/migrations/auth"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"os"
)

func main() {
	if err := run(); err != nil {
		logging.InitLogger()
		logging.Log.Info("run error:", err.Error())
		os.Exit(1)
	}
}

func run() error {
	config.InitConfig()
	log, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	database := os.Getenv("database")

	//  use sqlite for weak computers where there is no way to run docker with postgresql
	if database != "postgresql" && database != "sqlite" {
		return fmt.Errorf("database type: %s is not valid. Expecting `mysql` or `postgresql`", database)
	}

	var db *gorm.DB
	if database == "sqlite" {
		dbFile := viper.GetString("sqlite_path")
		db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			return fmt.Errorf("failed to open sqlite: %s", dbFile)
		}
		log.Info("Successfully connected to sqlite")
	} else {
		user := viper.GetString("POSTGRES_USER")
		pass := viper.GetString("POSTGRES_PASSWORD")
		dbname := viper.GetString("POSTGRES_DB_AUTH")

		dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kyiv", user, pass, dbname, "8888")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to connect to postgresql %s", err)
		}
		log.Info("Successfully connected to postgresql")
	}

	err = auth.Migrate(db)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %s", err)
	}
	return nil
}
