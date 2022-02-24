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

func TestDailyTemperatures(t *testing.T) {
	testCases := []struct {
		temperatures []int
		output       []int
	}{

		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			output:       []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
	}

	for i, testCase := range testCases {
		output := dailyTemperatures(testCase.temperatures)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
