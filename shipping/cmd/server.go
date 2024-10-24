package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"

	grpcapiShipping "github.com/alexandear/truckgo/shipping/grpc/grpcapi"
	userpb "github.com/alexandear/truckgo/user/grpcapi"
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

	userClient userpb.UserServiceClient
}

func (s *server) GetCoordinatesByAddress(ctx context.Context,
	req *grpcapiShipping.LocationRequest,
) (*grpcapiShipping.LocationResponse, error) {
	coordinates, err := geocode(ctx, getAPIKey(), req.Address)
	if err != nil {
		return nil, fmt.Errorf("error during getting coordinates: %v", err)
	}

	return &grpcapiShipping.LocationResponse{
		Longitude: coordinates[0],
		Latitude:  coordinates[1],
	}, nil
}

func (s *server) CalculateRouteByCoordinates(ctx context.Context,
	req *grpcapiShipping.CoordinatesRouteRequest,
) (*grpcapiShipping.RouteResponse, error) {
	var steps []*grpcapiShipping.Step

	startPoint := []float64{req.OriginLongitude, req.OriginLatitude}
	endPoint := []float64{req.DestinationLongitude, req.DestinationLatitude}

	segment, err := calculateRouteByCoordinatesImpl(ctx, getAPIKey(), startPoint, endPoint)
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

func (s *server) CalculateRoute(ctx context.Context, req *grpcapiShipping.RouteRequest) (*grpcapiShipping.RouteResponse, error) {
	// log, err := logging.InitLogger(serviceName)
	// if err != nil {
	// 	return &grpcapiShipping.RouteResponse{}, err
	// }

	fmt.Printf("Origin: %v\n", req.Origin)
	fmt.Printf("Destination: %v\n", req.Destination)

	originResponse, err := s.GetCoordinatesByAddress(ctx, &grpcapiShipping.LocationRequest{Address: req.Origin})
	if err != nil {
		return nil, fmt.Errorf("error during getting origin coordinates: %v", err)
	}

	destinationResponse, err := s.GetCoordinatesByAddress(ctx, &grpcapiShipping.LocationRequest{Address: req.Destination})
	if err != nil {
		return nil, fmt.Errorf("error during getting destination coordinates: %v", err)
	}

	fmt.Printf("Origin coordinates: Longitude: %v, Latitude: %v\n", originResponse.Longitude, originResponse.Latitude)
	fmt.Printf("Destination coordinates: Longitude: %v, Latitude: %v\n", destinationResponse.Longitude, destinationResponse.Latitude)

	return s.CalculateRouteByCoordinates(ctx, &grpcapiShipping.CoordinatesRouteRequest{
		OriginLongitude:      originResponse.Longitude,
		OriginLatitude:       originResponse.Latitude,
		DestinationLongitude: destinationResponse.Longitude,
		DestinationLatitude:  destinationResponse.Latitude,
	})
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

func (s *server) FindTheNearestDrivers(ctx context.Context,
	req *grpcapiShipping.DriverRequest,
) (*grpcapiShipping.DriverResponse, error) {
	if req.DriversCount == 0 {
		return nil, errors.New("no drivers were provided")
	}

	if req.ClientLatitude == 0 || req.ClientLongitude == 0 {
		return nil, errors.New("client coordinates are incorrect")
	}

	resp, err := s.userClient.ListDrivers(ctx, &userpb.ListDriverRequest{})
	fmt.Printf("Drivers: %v\n", resp.Drivers)
	if err != nil {
		return nil, err
	}

	type DriverDistance struct {
		id       uint32
		distance float64
		time     float64
	}
	driverDistances := make([]*DriverDistance, 0)

	for _, driver := range resp.Drivers {
		route, err := s.CalculateRouteByCoordinates(ctx, &grpcapiShipping.CoordinatesRouteRequest{
			OriginLongitude:      driver.Longitude,
			OriginLatitude:       driver.Latitude,
			DestinationLongitude: req.ClientLongitude,
			DestinationLatitude:  req.ClientLatitude,
		})
		if err != nil {
			// TODO think what is better to do in this case
			continue
		}

		driverDistances = append(driverDistances, &DriverDistance{
			id:       driver.Id,
			distance: route.Distance,
			time:     route.Time,
		})
	}

	findNearests := func(drivers []*DriverDistance, count uint32) []uint32 {
		if len(drivers) == 0 {
			return nil
		}

		var driverCount uint32
		driverLen := len(drivers)
		if driverLen <= int(^uint32(0)) {
			driverCount = uint32(driverLen)
		} else {
			// impossible case, but added for linter
			driverCount = ^uint32(0)
		}
		if count > driverCount {
			count = driverCount
		}

		sort.Slice(drivers, func(i, j int) bool {
			return drivers[i].time < drivers[j].time
		})
		res := make([]uint32, 0, count)
		for i := uint32(0); i < count; i++ {
			res = append(res, drivers[i].id)
		}
		return res
	}

	nearestDrivers := findNearests(driverDistances, req.DriversCount)
	if len(nearestDrivers) == 0 {
		return nil, errors.New("cannot find drivers with min distance, addresses are incorrect")
	}

	return &grpcapiShipping.DriverResponse{
		DriverIds: nearestDrivers,
		Message:   "Drivers were found!",
	}, nil
}
