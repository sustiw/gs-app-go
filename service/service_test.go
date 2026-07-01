package service

import (
	"reflect"
	"testing"
)

func TestCreatePackDetails(t *testing.T) {
	packSizes := []int{250, 500, 1000, 2000, 5000}

	tests := []struct {
		name           string
		targetQuantity int
		expectedPacks  string
	}{
		{
			name:           "Exact Match check for 500",
			targetQuantity: 500,
			expectedPacks:  "1x500",
		},
		{
			name:           "Check for 251 ",
			targetQuantity: 251,
			expectedPacks:  "1x500",
		},

		{
			name:           "check for 501",
			targetQuantity: 501,
			expectedPacks:  "1x500,1x250",
		},
		{
			name:           "check for 12001",
			targetQuantity: 12001,
			expectedPacks:  "2x5000,1x2000,1x250",
		},
		{
			name:           "Negative or Zero Quantity Check",
			targetQuantity: 0,
			expectedPacks:  "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CreatePackDetails(tc.targetQuantity, packSizes)

			if !reflect.DeepEqual(result.PackDetails, tc.expectedPacks) {
				t.Errorf("Test Failed: %s\nInput Qty: %d\nExpected: %v\nGot:      %v",
					tc.name, tc.targetQuantity, tc.expectedPacks, result.PackDetails)
			} else {
				t.Logf("Test Passed: %s (Qty %d -> %v)", tc.name, tc.targetQuantity, result.PackDetails)
			}
		})
	}
}
