package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/alexandear/truckgo/payment-service/database"
	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
)

const serviceName = "PAYMENT"

func main() {
	config.InitConfig()
	log, _ := logging.InitLogger(serviceName)
	if err := run(); err != nil {
		log.Info("run error:", err.Error())
		os.Exit(1)
	}
	log.Info("Successfully connected to DB")
}

func run() error {
	dbVarName := "POSTGRES_DB_" + serviceName
	port := viper.GetString("POSTGRES_PORT_" + serviceName)
	db, err := database.Initialize(dbVarName, port)
	if err != nil {
		return err
	}

	err = database.Migrate(db)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %s", err)
	}
	return nil
}
