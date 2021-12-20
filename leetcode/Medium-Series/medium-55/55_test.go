package main

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{2, 3, 1, 1, 4},
			output: true,
		},
		{
			input:  []int{3, 2, 1, 0, 4},
			output: false,
		},
	}

	for i, testCase := range testCases {
		output := canJump(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
