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

	"github.com/alexandear/truckgo/cmd/db"
	"github.com/alexandear/truckgo/internal/pkg/config"
	"github.com/alexandear/truckgo/internal/pkg/logging"
	"github.com/alexandear/truckgo/migrations/auth"

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

	db, err := db.Initialize(log, "POSTGRES_DB_AUTH")
	if err != nil {
		return err
	}

	err = auth.Migrate(db)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %s", err)
	}
	return nil
}
