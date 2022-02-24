package main

import (
	"reflect"
	"testing"
)

func TestDecodeAtIndex(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}},
			output: [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}},
		},
	}

	for i, testCase := range testCases {
		output := gameOfLife(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
