package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// Change working directory to the root of the project
	err := os.Chdir("../../")
	if err != nil {
		fmt.Printf("Error changing directory: %v\n", err)
		os.Exit(1)
	}

	err = godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func TestGeocode(t *testing.T) {
	tests := []struct {
		name           string
		address        string
		apiKey         string
		mockResponse   string
		mockStatusCode int
		expectedCoords []float64
		expectError    bool
	}{
		{
			name:           "valid address",
			address:        "парк Межигірʼя",
			apiKey:         os.Getenv("MAP_API_KEY"),
			mockResponse:   `{"features":[{"geometry":{"coordinates":[30.456913, 50.62605]}}]}`,
			mockStatusCode: http.StatusOK,
			expectedCoords: []float64{30.456913, 50.62605},
			expectError:    false,
		},
		{
			name:           "address not found",
			address:        "JustSomeMockAddress",
			apiKey:         os.Getenv("MAP_API_KEY"),
			mockResponse:   `{"features":[]}`,
			mockStatusCode: http.StatusOK,
			expectedCoords: nil,
			expectError:    true,
		},
		{
			name:           "invalid API key",
			address:        "площа Ринок, Львів",
			apiKey:         "test-api-key",
			mockResponse:   `{"error":"invalid API key"}`,
			mockStatusCode: http.StatusUnauthorized,
			expectedCoords: nil,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				if _, err := w.Write([]byte(tt.mockResponse)); err != nil {
					t.Fatalf("failed to write mock response: %v", err)
				}
			}))
			defer server.Close()

			ctx := context.Background()
			coords, err := geocode(ctx, tt.apiKey, tt.address)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError && !equalFloat64Slices(coords, tt.expectedCoords) {
				t.Errorf("expected coordinates: %v, got: %v", tt.expectedCoords, coords)
			}
		})
	}
}

func TestCalculateRouteByCoordinatesImpl(t *testing.T) {
	tests := []struct {
		name            string
		start           []float64
		end             []float64
		apiKey          string
		mockResponse    string
		mockStatusCode  int
		expectedSegment *Segment
		expectError     bool
	}{
		{
			name:           "valid coordinates",
			start:          []float64{30.519094, 50.444875},
			end:            []float64{30.519349, 50.443732},
			apiKey:         os.Getenv("MAP_API_KEY"),
			mockResponse:   `{"features":[{"properties":{"segments":[{"distance":1000,"duration":600,"steps":[{"distance":500,"duration":300,"instruction":"Head north"}]}]}}]}`,
			mockStatusCode: http.StatusOK,
			expectedSegment: &Segment{
				Distance: 359.9,
				Duration: 69.1,
				Steps: []Step{
					{
						Distance:    136.2,
						Duration:    32,
						Instruction: "Head east on вулиця Богдана Хмельницького",
					},
					{
						Distance:    142.5,
						Duration:    12.8,
						Instruction: "Turn right onto вулиця Хрещатик",
					},
					{
						Distance:    81.3,
						Duration:    24.2,
						Instruction: "Turn right",
					},
					{
						Distance:    0,
						Duration:    0,
						Instruction: "Arrive at your destination, on the right",
					},
				},
			},
			expectError: false,
		},
		{
			name:            "invalid API key",
			start:           []float64{30.456913, 50.62605},
			end:             []float64{30.5234, 50.4501},
			apiKey:          "invalid-api-key",
			mockResponse:    `{"error":"invalid API key"}`,
			mockStatusCode:  http.StatusUnauthorized,
			expectedSegment: nil,
			expectError:     true,
		},
		{
			name:            "no segments found",
			start:           []float64{0.0, 0.0},
			end:             []float64{45.0, 45.0},
			apiKey:          os.Getenv("MAP_API_KEY"),
			mockResponse:    `{"features":[{"properties":{"segments":[]}}]}`,
			mockStatusCode:  http.StatusOK,
			expectedSegment: nil,
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				if _, err := w.Write([]byte(tt.mockResponse)); err != nil {
					t.Fatalf("failed to write mock response: %v", err)
				}
			}))
			defer server.Close()

			ctx := context.Background()
			segment, err := calculateRouteByCoordinatesImpl(ctx, tt.apiKey, tt.start, tt.end)

			if (err != nil) != tt.expectError {
				t.Errorf("expected error: %v, got: %v", tt.expectError, err)
			}
			if !tt.expectError && !equalSegments(segment, tt.expectedSegment) {
				t.Errorf("expected segment: %v, got: %v", tt.expectedSegment, segment)
			}
		})
	}
}

func equalSegments(a, b *Segment) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Distance != b.Distance || a.Duration != b.Duration {
		return false
	}
	if len(a.Steps) != len(b.Steps) {
		return false
	}
	for i := range a.Steps {
		if a.Steps[i] != b.Steps[i] {
			return false
		}
	}
	return true
}

func equalFloat64Slices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
