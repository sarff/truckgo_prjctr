package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func getAPIKey() string {
	return os.Getenv("MAP_API_KEY")
}

type GeocodeResponse struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

func geocode(ctx context.Context, address string) ([]float64, error) {
	escapedAddress := url.QueryEscape(address)
	url := fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%s&text=%s", getAPIKey(), escapedAddress)

	fmt.Printf("URL = %v\n", url)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot perform the request: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var geoResp GeocodeResponse
	if err := json.Unmarshal(body, &geoResp); err != nil {
		return nil, err
	}

	if len(geoResp.Features) == 0 {
		return nil, errors.New("coordinates are not found")
	}

	return geoResp.Features[0].Geometry.Coordinates, nil
}

type RouteRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
}

type RouteResponse struct {
	Features []struct {
		Properties struct {
			Segments []Segment `json:"segments"`
		} `json:"properties"`
	} `json:"features"`
}

type Segment struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Steps    []Step  `json:"steps"`
}

type Step struct {
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
	Instruction string  `json:"instruction"`
}

func calculateRouteByCoordinatesImpl(ctx context.Context, start, end []float64) (*Segment, error) {
	url := "https://api.openrouteservice.org/v2/directions/driving-car/geojson"

	reqBody := RouteRequest{
		Coordinates: [][]float64{
			start,
			end,
		},
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", getAPIKey())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("cannot perform the request: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var routeResp RouteResponse
	if err := json.Unmarshal(body, &routeResp); err != nil {
		return nil, err
	}

	if len(routeResp.Features) == 0 || len(routeResp.Features[0].Properties.Segments) == 0 {
		return nil, errors.New("segments are not found")
	}

	return &routeResp.Features[0].Properties.Segments[0], nil
}
