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
	userpb "github.com/alexandear/truckgo/user/grpcapi"
	"github.com/alexandear/truckgo/user/internal/database"
	"github.com/alexandear/truckgo/user/internal/services"
	"gorm.io/gorm"
	"net"

	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"os"
)

const serviceName = "USER"

func main() {
	config.InitConfig()
	log, _ := logging.InitLogger(serviceName)
	if err := run(log); err != nil {
		log.Info("run error:", err.Error())
		os.Exit(1)
	}
}

func run(log *logging.Logger) error {
	db, err := initDB()
	if err != nil {
		return err
	}
	log.Info("Successfully connected to DB")

	done := make(chan bool, 1)
	go func() {
		if err := initGRPCServer(log, db); err != nil {
			log.Error("gRPC server error:", "GRPC", err)
			os.Exit(1)
		}
		done <- true
	}()
	<-done

	return nil
}

func initDB() (*gorm.DB, error) {
	dbVarName := "POSTGRES_DB_" + serviceName
	port := viper.GetString("POSTGRES_PORT_" + serviceName)
	db, err := database.Initialize(dbVarName, port)
	if err != nil {
		return nil, err
	}

	err = database.Migrate(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err)
	}
	return db, nil
}

func initGRPCServer(log *logging.Logger, db *gorm.DB) error {
	port := viper.GetString("GRPC_PORT_" + serviceName)

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &services.UserServiceServer{
		DB:     db,
		Logger: log,
	})

	log.Info("gRPC server is running on", "port", port)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
