package main

import (
	"fmt"
	authpb "github.com/alexandear/truckgo/auth/grpcapi"
	"github.com/alexandear/truckgo/auth/internal/database"
	"github.com/alexandear/truckgo/auth/internal/services"

	"gorm.io/gorm"

	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
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
	dbs, err := initDB()
	if err != nil {
		return err
	}

	done := make(chan bool, 1)
	//TODO: add ctx.
	go func() {
		if err = startGRPCServer(log, dbs); err != nil {
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

func startGRPCServer(log *logging.Logger, db *gorm.DB) error {
	lis, err := net.Listen("tcp", ":"+os.Getenv("GRPC_PORT_"+serviceName))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, &services.AuthServiceServer{
		Logger: log,
		DB:     db,
	})

	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	log.Info("gRPC server is starting...")

	return nil
}
