package main

import (
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
	"github.com/alexandear/truckgo/shipping/database"
	grpcapiShipping "github.com/alexandear/truckgo/shipping/grpc/grpcapi"
)

const serviceName = "SHIPPING"

func main() {
	config.InitConfig()
	log, _ := logging.InitLogger(serviceName)
	if err := run(log); err != nil {
		log.Info("run error:", err.Error())
		os.Exit(1)
	}
}

func run(log *logging.Logger) error {
	err := initDB()
	if err != nil {
		return err
	}
	log.Info("Successfully connected to DB")

	done := make(chan bool, 1)
	go func() {
		if err := startGRPCServer(log); err != nil {
			log.Error("gRPC server error:", "GRPC", err)
			os.Exit(1)
		}
		done <- true
	}()
	<-done

	return nil
}

func initDB() error {
	db, err := database.Initialize()
	if err != nil {
		return err
	}

	err = database.Migrate(db)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %s", err)
	}
	return nil
}

func startGRPCServer(log *logging.Logger) error {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT_"+serviceName))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpcapiShipping.RegisterShippingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	log.Info("gRPC server is starting...")

	return nil
}
