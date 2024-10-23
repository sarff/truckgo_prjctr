package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewOrder(t *testing.T) {
	price := 10.2
	userID := uint32(1)
	got := NewOrder(price, userID)

	want := Order{
		Status: StatusNew,
		Price:  price,
		UserID: userID,
	}

	if !cmp.Equal(want.Status, got.Status) || !cmp.Equal(want.Price, got.Price) || !cmp.Equal(want.UserID, got.UserID) {
		t.Errorf("NewOrder(): got = %v, want %v", got, want)
	}
}

func TestValidateStatus(t *testing.T) {
	testCases := []struct {
		name      string
		order     Order
		newStatus Status
		want      bool
	}{
		{
			name:      "Valid state change from new to accepted",
			order:     Order{Status: StatusNew},
			newStatus: StatusAccepted,
			want:      true,
		},
		{
			name:      "Invalid state change from new to done",
			order:     Order{Status: StatusNew},
			newStatus: StatusDone,
			want:      false,
		},
		{
			name:      "Valid state change from accepted to in progress",
			order:     Order{Status: StatusAccepted},
			newStatus: StatusInProgress,
			want:      true,
		},
		{
			name:      "Invalid state change from accepted to done",
			order:     Order{Status: StatusAccepted},
			newStatus: StatusDone,
			want:      false,
		},
		{
			name:      "Valid state change from accepted to cancelled",
			order:     Order{Status: StatusAccepted},
			newStatus: StatusCancelled,
			want:      true,
		},
		{
			name:      "Invalid state change from in progress to cancelled",
			order:     Order{Status: StatusInProgress},
			newStatus: StatusCancelled,
			want:      false,
		},
		{
			name:      "Valid state change from in progress to done",
			order:     Order{Status: StatusInProgress},
			newStatus: StatusDone,
			want:      true,
		},
		{
			name:      "Invalid state change from cancelled in progress to done",
			order:     Order{Status: StatusCancelled},
			newStatus: StatusDone,
			want:      false,
		},
		{
			name:      "Invalid state change from done to new",
			order:     Order{Status: StatusDone},
			newStatus: StatusNew,
			want:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateStatus(tc.order, tc.newStatus)

			if got != tc.want {
				t.Errorf("ValidateStatus(): got = %v, want %v", got, tc.want)
			}
		})
	}
}
