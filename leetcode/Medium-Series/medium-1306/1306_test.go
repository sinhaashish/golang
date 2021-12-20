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

func TestCanReach(t *testing.T) {
	testCases := []struct {
		input  []int
		start  int
		output bool
	}{
		{
			input:  []int{4, 2, 3, 0, 3, 1, 2},
			start:  5,
			output: true,
		},
		{
			input:  []int{4, 2, 3, 0, 3, 1, 2},
			start:  0,
			output: true,
		},
		{
			input:  []int{3, 0, 2, 1, 2},
			start:  2,
			output: false,
		},
	}

	for i, testCase := range testCases {
		output := canReach(testCase.input, testCase.start)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
