package main

import (
	"context"
	"fmt"
	"github.com/alexandear/truckgo/shared/config"
	"github.com/alexandear/truckgo/shared/logging"
	grpcapi2 "github.com/alexandear/truckgo/shipping-service/grpc/grpcapi"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	grpcapi2.UnimplementedRouteServiceServer
}

func (s *server) CalculateRoute(ctx context.Context, req *grpcapi2.RouteRequest) (*grpcapi2.RouteResponse, error) {
	fmt.Printf("Origin: %v\n", req.Origin)
	fmt.Printf("Destination: %v\n", req.Destination)

	startPoint, err := geocode(req.Origin)
	if err != nil {
		log.Fatalf("Error during getting origin coordinates: %v", err)
	}

	endPoint, err := geocode(req.Destination)
	if err != nil {
		log.Fatalf("Error during getting destination coordinates: %v", err)
	}

	fmt.Printf("Origin coordinates: %v\n", startPoint)
	fmt.Printf("Destination coordinates: %v\n", endPoint)

	featureCollection, err := calculateRoute(startPoint, endPoint)
	if err != nil {
		log.Fatalf("Error during route calculation: %v", err)
	}

	var steps []*grpcapi2.Step

	for _, feature := range featureCollection.Features {
		for _, segment := range feature.Properties.Segments {
			fmt.Printf("Segment Distance: %.2f meters\n", segment.Distance)
			fmt.Printf("Segment Duration: %.2f seconds\n", segment.Duration)
			for _, step := range segment.Steps {
				fmt.Printf("Instruction: %s\n", step.Instruction)

				step := &grpcapi2.Step{
					Distance:    float32(step.Distance / 1000),
					Duration:    int32(step.Duration / 60),
					Instruction: step.Instruction,
				}
				steps = append(steps, step)
			}
		}
	}

	return &grpcapi2.RouteResponse{
		Message: "The route calculated successfully!",
		Steps:   steps,
	}, nil
}
func (s *server) TestFunc(ctx context.Context, req *grpcapi2.TestRequest) (*grpcapi2.TestResponse, error) {
	return &grpcapi2.TestResponse{
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
	grpcapi2.RegisterRouteServiceServer(grpcServer, &server{})
	log.Info("gRPC server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Error("failed to serve: %v", err)
	}
}
