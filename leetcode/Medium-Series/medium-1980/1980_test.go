package main

import (
	"reflect"
	"testing"
)

func TestSubArrayRanges(t *testing.T) {
	testCases := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"01", "10"},
			output: "11",
		},
	}

	for i, testCase := range testCases {
		output := findDifferentBinaryString(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
