package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"net"
	"os"
	"os/signal"
	"syscall"

	authpb "github.com/alexandear/truckgo/auth/grpcapi"
	"github.com/alexandear/truckgo/auth/internal/database"
	"github.com/alexandear/truckgo/auth/internal/services"
	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
)

const serviceName = "AUTH"

func main() {
	config.InitConfig()
	log, err := logging.InitLogger(serviceName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, log); err != nil {
		log.Info("run error:", err.Error())
		os.Exit(1)
	}
	log.Info("Successfully connected to DB")
}

func run(ctx context.Context, log *logging.Logger) error {
	dbs, err := initDB(ctx)
	if err != nil {
		return err
	}
	defer func() {
		sqlDB, _ := dbs.DB()
		if err := sqlDB.Close(); err != nil {
			log.Error("Failed to close database connection", "error", err)
		}
	}()

	grpcErrChan := make(chan error, 1)
	go func() {
		grpcErrChan <- startGRPCServer(ctx, log, dbs)
	}()

	select {
	case <-ctx.Done():
		log.Info("Received shutdown signal, shutting down...")
		return nil
	case err := <-grpcErrChan:
		if err != nil {
			log.Error("gRPC server error:", err)
			return err
		}
	}

	return nil
}

func initDB(ctx context.Context) (*gorm.DB, error) {
	db, err := database.Initialize()
	if err != nil {
		return nil, err
	}

	err = database.Migrate(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %s", err)
	}
	return db, nil
}

func startGRPCServer(ctx context.Context, log *logging.Logger, db *gorm.DB) error {
	port := os.Getenv("GRPC_PORT_" + serviceName)
	if port == "" {
		return fmt.Errorf("GRPC_PORT_%s environment variable is not set", serviceName)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	authpb.RegisterAuthServiceServer(grpcServer, &services.AuthServiceServer{
		Logger: log,
		DB:     db,
	})

	go func() {
		<-ctx.Done()
		log.Info("Gracefully stopping gRPC server...")
		grpcServer.GracefulStop()
	}()

	log.Info("gRPC server is starting...", "port", port)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
