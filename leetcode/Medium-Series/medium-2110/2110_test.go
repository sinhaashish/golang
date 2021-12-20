// 880. Decoded String at Index

// You are given an encoded string s. To decode the string to a tape, the encoded string is read one character at a time and the following steps are taken:

// If the character read is a letter, that letter is written onto the tape.
// If the character read is a digit d, the entire current tape is repeatedly written d - 1 more times in total.
// Given an integer k, return the kth letter (1-indexed) in the decoded string.

package main

import (
	"reflect"
	"testing"
)

func TestMinSubarray(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 1, 4},
			output: 7,
		},
		{
			input: []int{8, 6, 7, 7},

			output: 4,
		},
		{
			input:  []int{1},
			output: 1,
		},
	}

	for i, testCase := range testCases {
		output := getDescentPeriods(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
