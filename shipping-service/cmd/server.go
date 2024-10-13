package main

import (
	"context"
	"fmt"

	grpcapiShipping "github.com/alexandear/truckgo/shipping-service/grpc/grpcapi"
	"github.com/spf13/viper"
)

func getStartPrice() float64 {
	return viper.GetFloat64("START_PRICE")
}

func getKMPrice() float64 {
	return viper.GetFloat64("KM_PRICE")
}

func priceByFormula(distance float64) float64 {
	return getStartPrice() + distance*getKMPrice()
}

type server struct {
	grpcapiShipping.UnimplementedShippingServiceServer
}

func (s *server) CalculatePrice(ctx context.Context, req *grpcapiShipping.PriceRequest) (*grpcapiShipping.PriceResponse, error) {
	res, err := s.CalculateRoute(ctx, &grpcapiShipping.RouteRequest{
		Origin:      req.Origin,
		Destination: req.Destination,
	})

	if err != nil {
		return &grpcapiShipping.PriceResponse{}, err
	}

	// TODO some smart calculations
	price := priceByFormula(res.Distance)

	return &grpcapiShipping.PriceResponse{
		Message:  "The route calculated successfully!",
		Price:    price,
		Distance: res.Distance,
		Time:     res.Time,
	}, nil
}

func (s *server) CalculateRoute(ctx context.Context, req *grpcapiShipping.RouteRequest) (*grpcapiShipping.RouteResponse, error) {
	// log, err := logging.InitLogger(serviceName)
	// if err != nil {
	// 	return &grpcapiShipping.RouteResponse{}, err
	// }

	fmt.Printf("Origin: %v\n", req.Origin)
	fmt.Printf("Destination: %v\n", req.Destination)

	startPoint, err := geocode(req.Origin)
	if err != nil {
		return &grpcapiShipping.RouteResponse{}, fmt.Errorf("error during getting origin coordinates: %v", err)
	}

	endPoint, err := geocode(req.Destination)
	if err != nil {
		return &grpcapiShipping.RouteResponse{}, fmt.Errorf("error during getting destination coordinates: %v", err)
	}

	fmt.Printf("Origin coordinates: %v\n", startPoint)
	fmt.Printf("Destination coordinates: %v\n", endPoint)

	featureCollection, err := calculateRoute(startPoint, endPoint)
	if err != nil {
		return &grpcapiShipping.RouteResponse{}, fmt.Errorf("error during route calculation: %v", err)
	}

	var steps []*grpcapiShipping.Step

	//	NOTE only one feature, and only one properties inside feature
	var distance float64
	var duration float64

	if len(featureCollection.Features) != 1 || len(featureCollection.Features[0].Properties.Segments) != 1 {
		return &grpcapiShipping.RouteResponse{}, fmt.Errorf("incorrect output format")
	}

	segment := featureCollection.Features[0].Properties.Segments[0]
	distance = segment.Distance
	duration = segment.Duration

	fmt.Printf("Segment Distance: %.2f meters\n", segment.Distance)
	fmt.Printf("Segment Duration: %.2f seconds\n", segment.Duration)
	for _, step := range segment.Steps {
		fmt.Printf("Instruction: %s\n", step.Instruction)

		step := &grpcapiShipping.Step{
			Instruction: step.Instruction,
			Distance:    step.Distance / 1000,
			Duration:    step.Duration / 60,
		}
		steps = append(steps, step)
	}

	_ = distance
	_ = duration
	_ = steps

	return &grpcapiShipping.RouteResponse{
		Message: "The route calculated successfully!",
		Steps:    steps,
		Distance: distance / 1000,
		Time:     duration / 60,
	}, nil
}
func (s *server) TestFunc(ctx context.Context, req *grpcapiShipping.TestRequest) (*grpcapiShipping.TestResponse, error) {
	return &grpcapiShipping.TestResponse{
		Message: "Some testing!",
	}, nil
}
