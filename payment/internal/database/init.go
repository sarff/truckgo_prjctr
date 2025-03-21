package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	databaseType := viper.GetString("db_type")

	//  use sqlite for weak computers where there is no way to run docker with postgresql
	if databaseType != "postgresql" && databaseType != "sqlite" {
		return nil, fmt.Errorf("database type: %s is not valid. Expecting `sqlite` or `postgresql`", databaseType)
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
	} else {
		user := viper.GetString("POSTGRES_USER")
		pass := viper.GetString("POSTGRES_PASSWORD")
		dbname := "truckgodb_payment"

		dsnTemplate := "host=db_payment user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Kyiv"
		dsn := fmt.Sprintf(dsnTemplate, user, pass, dbname)

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to postgresql %s", err)
		}
	}

	return db, nil
}
