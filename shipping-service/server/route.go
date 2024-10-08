package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// TODO read from .env
const API_KEY = "my_api_key"

type GeocodeResponse struct {
	Features []Feature `json:"features"`
}

type Feature struct {
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Coordinates []float64 `json:"coordinates"`
}

type FeatureCollection struct {
	Type     string     `json:"type"`
	BBox     []float64  `json:"bbox"`
	Features []Feature2 `json:"features"`
}

type Feature2 struct {
	Type       string     `json:"type"`
	Geometry   Geometry2  `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry2 struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type Properties struct {
	Segments  []Segment `json:"segments"`
	Summary   Summary   `json:"summary"`
	WayPoints []int     `json:"way_points"`
}

type Segment struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
	Steps    []Step  `json:"steps"`
}

type Step struct {
	Distance    float64 `json:"distance"`
	Duration    float64 `json:"duration"`
	Type        int     `json:"type"`
	Instruction string  `json:"instruction"`
	Name        string  `json:"name"`
	WayPoints   []int   `json:"way_points"`
}

type Summary struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

func geocode(address string) ([]float64, error) {
	escapedAddress := url.QueryEscape(address)
	url := fmt.Sprintf("https://api.openrouteservice.org/geocode/search?api_key=%s&text=%s", API_KEY, escapedAddress)

	fmt.Printf("URL = %v\n", url)

	resp, err := http.Get(url)
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
		return nil, fmt.Errorf("coordinates are not found")
	}

	return geoResp.Features[0].Geometry.Coordinates, nil
}

type RouteRequest struct {
	Coordinates [][]float64 `json:"coordinates"`
}

func calculateRoute(start, end []float64) (*FeatureCollection, error) {
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

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", API_KEY)
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

	var featureCollection FeatureCollection
	if err := json.Unmarshal(body, &featureCollection); err != nil {
		return nil, err
	}

	return &featureCollection, nil
}
