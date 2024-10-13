package main

import (
	"fmt"
	"net"

	"github.com/alexandear/truckgo/order-service/database"
	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
	grpcapiOrder "github.com/alexandear/truckgo/order-service/grpc/grpcapi"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"os"
)

const serviceName = "ORDER"

func main() {
	config.InitConfig()
	log, _ := logging.InitLogger(serviceName)
	if err := run(log); err != nil {
		log.Info("run error:", err.Error())
		os.Exit(1)
	}
	log.Info("Successfully connected to DB")
}

func run(log *logging.Logger) error {
	// temporary disabled
	// err := initDB()
	// if err != nil {
	// 	return err
	// }

	err := initGRPCServer(log)
	if err != nil {
		return err
	}

	return nil
}

func initDB() error {
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

func initGRPCServer(log *logging.Logger) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcapiOrder.RegisterOrderServiceServer(grpcServer, &server{})

	log.Info("gRPC server is running on port :50052")
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
