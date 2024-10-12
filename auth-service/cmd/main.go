package main

import (
	"fmt"
	"github.com/alexandear/truckgo/auth-service/database"
	pb "github.com/alexandear/truckgo/auth-service/generated"
	"github.com/alexandear/truckgo/auth-service/internal/services"
	"gorm.io/gorm"

	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
)

const serviceName = "AUTH"

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

	done := make(chan bool, 1)
	//TODO: add ctx
	go func() {
		if err = startGRPCServer(log, db); err != nil {
			log.Error("gRPC server error:", "GRPC", err)
			os.Exit(1)
		}
		done <- true
	}()
	<-done

	return nil
}

func startGRPCServer(log *logging.Logger, db *gorm.DB) error {
	lis, err := net.Listen("tcp", ":"+viper.GetString("GRPC_PORT_"+serviceName))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, &services.AuthServiceServer{
		Logger: log,
		DB:     db,
	})

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	log.Info("gRPC server is starting...")

	return nil
}
