package main

import (
	"reflect"
	"testing"
)

func TestMinSubarray(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output int64
	}{
		{
			input:  [][]int{{1, -1}, {-1, 1}},
			output: 4,
		},
		{
			input:  [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}},
			output: 16,
		},
	}

	for i, testCase := range testCases {
		output := maxMatrixSum(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
