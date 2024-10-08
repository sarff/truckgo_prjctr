package main

import (
	"fmt"
	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/alexandear/truckgo/shipping-service/database"

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

	db, err := database.Initialize(log, "POSTGRES_DB_SHIPPINGS")
	if err != nil {
		return err
	}

	err = database.Migrate(db)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %s", err)
	}
	return nil
}
