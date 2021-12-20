package main

import (
	"reflect"
	"testing"
)

func TestSubArrayRanges(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 2, 3},
			output: 4,
		},
		{
			input:  []int{1, 3, 3},
			output: 4,
		},
		{
			input:  []int{4, -2, -3, 4, 1},
			output: 59,
		},
	}

	for i, testCase := range testCases {
		output := subArrayRanges(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
