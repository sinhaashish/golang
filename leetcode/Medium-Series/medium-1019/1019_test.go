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
		output []int
	}{
		{
			input:  []int{2, 7, 4, 3, 5},
			output: []int{7, 0, 5, 5, 0},
		},
		{
			input:  []int{2, 1, 5},
			output: []int{5, 5, 0},
		},
	}

	for i, testCase := range testCases {
		output := nextLargerNodes(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
