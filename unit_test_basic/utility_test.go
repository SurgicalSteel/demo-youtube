package unit_test_basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterOdd(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "empty slice input returns empty slice",
			input:          []int{},
			expectedOutput: []int{},
		},
		{
			name:           "even only slice input returns empty slice",
			input:          []int{32, 46, 76, 68, 100},
			expectedOutput: []int{},
		},
		{
			name:           "mixed slice (even and odd) input returns odd slice",
			input:          []int{33, 34, 35, 36, 37, 38, 39, 40},
			expectedOutput: []int{33, 35, 37, 39},
		},
		{
			name:           "odd only slice input returns same slice",
			input:          []int{33, 35, 37, 39},
			expectedOutput: []int{33, 35, 37, 39},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualOutput := FilterOdd(tc.input)
			assertionResult := assert.ElementsMatch(t, actualOutput, tc.expectedOutput)
			if !assertionResult {
				t.Error("actualOutput is not same with expectedOutput")
			}
		})
	}

}
