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
			expectedPacks:  "1x1000",
		},
		{
			name:           "check for 12001",
			targetQuantity: 12001,
			expectedPacks:  "2x5000,1x2000,1x250",
		},
		{
			name:           "Negative or Zero Quantity Boundary Check",
			targetQuantity: 0,
			expectedPacks:  "",
		},
	}

	// Loop through and run each scenario
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Call your actual service class code function
			result := CreatePackDetails(tc.targetQuantity, packSizes)

			// Deep equal checks arrays content values safely
			if !reflect.DeepEqual(result.PackDetails, tc.expectedPacks) {
				t.Errorf("Test Failed: %s\nInput Qty: %d\nExpected: %v\nGot:      %v",
					tc.name, tc.targetQuantity, tc.expectedPacks, result.PackDetails)
			} else {
				t.Logf("Test Passed: %s (Qty %d -> %v)", tc.name, tc.targetQuantity, result.PackDetails)
			}
		})
	}
}

// TestDynamicSizesConfig confirms that adding or removing custom pack configurations works dynamically.
func TestDynamicSizesConfig(t *testing.T) {
	// Custom array without 400 or 500, but adding a unique 1500 pack size config
	customSizes := []int{250, 1000, 1500, 2000}

	// Testing input 1301 with our alternative array layout
	result := CreatePackDetails(1301, customSizes)
	expected := []string{"1x1500"} // 1x1500 is now closer than 1x2000 for covering 1301

	if !reflect.DeepEqual(result.PackDetails, expected) {
		t.Errorf("❌ Dynamic sizes failed! Got %v, expected %v", result.PackDetails, expected)
	}
}
