package main

import (
	"reflect"
	"testing"
)

func TestComplexNumberMultiply(t *testing.T) {
	testCases := []struct {
		input  []string
		output string
	}{
		{
			input:  []string{"1+-1i", "1+-1i"},
			output: "0+2i",
		},
		{
			input:  []string{"1+1i", "1+1i"},
			output: "0+2i",
		},
	}

	for i, testCase := range testCases {
		output := complexNumberMultiply(testCase.input[0], testCase.input[1])
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
