package main

import (
	"context"
	"fmt"
	"net"

	"github.com/alexandear/truckgo/internal/grpc/grpcapi"
	"github.com/alexandear/truckgo/internal/pkg/config"
	"github.com/alexandear/truckgo/internal/pkg/logging"
	"google.golang.org/grpc"
)

type server struct {
	grpcapi.UnimplementedRouteServiceServer
}

func (s *server) CalculateRoute(ctx context.Context, req *grpcapi.RouteRequest) (*grpcapi.RouteResponse, error) {
	return &grpcapi.RouteResponse{
		Message: "Stub calculate route",
	}, nil
}
func (s *server) TestFunc(ctx context.Context, req *grpcapi.TestRequest) (*grpcapi.TestResponse, error) {
	return &grpcapi.TestResponse{
		Message: "Some testing!",
	}, nil
}

func main() {
	config.InitConfig()
	log, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	// FIXME logs don't work now
	fmt.Println("Just a test")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Error("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcapi.RegisterRouteServiceServer(grpcServer, &server{})
	log.Info("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
	}
}
