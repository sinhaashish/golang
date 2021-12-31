package main

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{2, 3, 1, 1, 4},
			output: 2,
		},
		{
			input:  []int{2, 3, 0, 1, 4},
			output: 2,
		},
	}

	for i, testCase := range testCases {
		output := jump(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected : %v, got: %v", i+1, testCase.output, output)
		}
	}

}
