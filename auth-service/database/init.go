package database

import (
	"fmt"
	"github.com/alexandear/truckgo/shared/logging"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TODO maybe not needed to pass 'log' as a parameter
func Initialize(log *logging.Logger, dbVarName string) (*gorm.DB, error) {
	databaseType := os.Getenv("database")

	//  use sqlite for weak computers where there is no way to run docker with postgresql
	if databaseType != "postgresql" && databaseType != "sqlite" {
		return nil, fmt.Errorf("database type: %s is not valid. Expecting `mysql` or `postgresql`", databaseType)
	}

	var db *gorm.DB
	var err error

	if databaseType == "sqlite" {
		dbFile := viper.GetString("sqlite_path")
		db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: false,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to open sqlite: %s", dbFile)
		}
		log.Info("Successfully connected to sqlite")
	} else {
		user := viper.GetString("POSTGRES_USER")
		pass := viper.GetString("POSTGRES_PASSWORD")
		dbname := viper.GetString(dbVarName)

		// TODO: перенести порт в конфіг
		dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Kyiv", user, pass, dbname, "8888")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to postgresql %s", err)
		}
		log.Info("Successfully connected to postgresql")
	}

	return db, nil
}
