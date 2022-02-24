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

func TestDecodeAtIndex(t *testing.T) {
	testCases := []struct {
		difficulty []int
		profit     []int
		worker     []int
		output     int
	}{

		{
			difficulty: []int{85, 47, 57},
			profit:     []int{24, 66, 99},
			worker:     []int{40, 25, 25},
			output:     0,
		},
		{
			difficulty: []int{2, 4, 6, 8, 10},
			profit:     []int{10, 20, 30, 40, 50},
			worker:     []int{4, 5, 6, 7},
			output:     100,
		},
	}

	for i, testCase := range testCases {
		output := maxProfitAssignment(testCase.difficulty, testCase.profit, testCase.worker)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
