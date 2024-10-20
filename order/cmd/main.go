package main

import (
	"fmt"
	grpcapiOrder "github.com/alexandear/truckgo/order/grpcapi"
	"github.com/alexandear/truckgo/order/internal/database"
	"github.com/alexandear/truckgo/order/internal/repository"
	"gorm.io/gorm"
	"net"

	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
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
}

func run(log *logging.Logger) error {
	db, err := initDB()
	if err != nil {
		return err
	}
	log.Info("Successfully connected to DB")

	done := make(chan bool, 1)
	go func() {
		if err := initGRPCServer(db, log); err != nil {
			log.Error("gRPC server error:", "GRPC", err)
			os.Exit(1)
		}
		done <- true
	}()
	<-done

	return nil
}

func initDB() (*gorm.DB, error) {
	db, err := database.Initialize()
	if err != nil {
		return nil, err
	}

	err = database.Migrate(db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err)
	}
	return db, nil
}

func initGRPCServer(db *gorm.DB, log *logging.Logger) error {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT_"+serviceName))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderRepository := repository.NewOrderRepository(db)
	grpcapiOrder.RegisterOrderServer(grpcServer, &server{log: log, orderRepository: orderRepository})

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	log.Info("gRPC server is starting...")

	return nil
}
