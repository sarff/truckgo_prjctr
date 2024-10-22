package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"

	grpcapiShipping "github.com/alexandear/truckgo/shipping/grpc/grpcapi"
)

func getStartPrice() float64 {
	price, err := strconv.ParseFloat(os.Getenv("START_PRICE"), 64)
	if err != nil {
		return 0.0
	}
	return price
}

func getKMPrice() float64 {
	price, err := strconv.ParseFloat(os.Getenv("KM_PRICE"), 64)
	if err != nil {
		return 0.0
	}
	return price
}

func priceByFormula(distance float64) float64 {
	return getStartPrice() + distance*getKMPrice()
}

type server struct {
	grpcapiShipping.UnimplementedShippingServiceServer
}

func (s *server) GetCoordinatesByAddress(ctx context.Context, req *grpcapiShipping.LocationRequest) (*grpcapiShipping.LocationResponse, error) {
	coordinates, err := geocode(ctx, req.Address)
	if err != nil {
		return nil, fmt.Errorf("error during getting coordinates: %v", err)
	}

	return &grpcapiShipping.LocationResponse{
		Longitude: coordinates[0],
		Latitude: coordinates[1],
	}, nil
}

func (s *server) CalculatePrice(ctx context.Context, req *grpcapiShipping.PriceRequest) (*grpcapiShipping.PriceResponse, error) {
	res, err := s.CalculateRoute(ctx, &grpcapiShipping.RouteRequest{
		Origin:      req.Origin,
		Destination: req.Destination,
	})
	if err != nil {
		return nil, err
	}

	// TODO some smart calculations
	price := priceByFormula(res.Distance)

	return &grpcapiShipping.PriceResponse{
		Message:  "The price calculated successfully!",
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

	startPoint, err := geocode(ctx, req.Origin)
	if err != nil {
		return nil, fmt.Errorf("error during getting origin coordinates: %v", err)
	}

	endPoint, err := geocode(ctx, req.Destination)
	if err != nil {
		return nil, fmt.Errorf("error during getting destination coordinates: %v", err)
	}

	fmt.Printf("Origin coordinates: %v\n", startPoint)
	fmt.Printf("Destination coordinates: %v\n", endPoint)

	var steps []*grpcapiShipping.Step

	segment, err := calculateRouteByCoordinates(ctx, startPoint, endPoint)
	if err != nil {
		return nil, fmt.Errorf("error during route calculation: %v", err)
	}

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

	return &grpcapiShipping.RouteResponse{
		Message:  "The route calculated successfully!",
		Steps:    steps,
		Distance: segment.Distance / 1000,
		Time:     segment.Duration / 60,
	}, nil
}

func (s *server) FindTheNearestDriver(ctx context.Context,
	req *grpcapiShipping.DriverRequest,
) (*grpcapiShipping.DriverResponse, error) {
	if len(req.Drivers) == 0 {
		return nil, errors.New("no drivers were provided")
	}

	type DriverDistance struct {
		id       uint32
		distance float64
		time     float64
	}
	var minDriverDistance *DriverDistance = nil

	for _, driver := range req.Drivers {
		route, err := s.CalculateRoute(ctx, &grpcapiShipping.RouteRequest{})
		if err == nil {
			if minDriverDistance == nil {
				minDriverDistance = &DriverDistance{
					id:       driver.Id,
					distance: route.Distance,
					time:     route.Time,
				}
			} else if route.Time < minDriverDistance.time {
				minDriverDistance.id = driver.Id
				minDriverDistance.distance = route.Distance
				minDriverDistance.time = route.Time
			}
		}
	}

	if minDriverDistance == nil {
		return nil, errors.New("cannot find driver with min distance, addresses are incorrect")
	}

	return &grpcapiShipping.DriverResponse{
		Id:       minDriverDistance.id,
		Distance: minDriverDistance.distance,
		Duration: minDriverDistance.time,
		Message:  "Driver was found!",
	}, nil
}

func (s *server) TestFunc(ctx context.Context, req *grpcapiShipping.TestRequest) (*grpcapiShipping.TestResponse, error) {
	return &grpcapiShipping.TestResponse{
		Message: "Some testing!",
	}, nil
}
