package main

import (
	"reflect"
	"testing"
)

func TestMinSubarray(t *testing.T) {
	testCases := []struct {
		input  string
		spaces []int
		output string
	}{
		{
			input:  "LeetcodeHelpsMeLearn",
			spaces: []int{8, 13, 15},
			output: "Leetcode Helps Me Learn",
		},
	}

	for i, testCase := range testCases {
		output := addSpaces(testCase.input, testCase.spaces)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
