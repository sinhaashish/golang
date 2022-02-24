// There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

// A province is a group of directly or indirectly connected cities and no other cities outside of the group.

// You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

// Return the total number of provinces.

package main

import (
	"reflect"
	"testing"
)

func TestDecodeAtIndex(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output int
	}{
		{
			input:  [][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {1, 0, 1, 1}},
			output: 3,
		},
		{
			input:  [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			output: 3,
		},
		{
			input:  [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			output: 2,
		},
	}

	for i, testCase := range testCases {
		output := findCircleNum(testCase.input)
		if !reflect.DeepEqual(output, testCase.output) {
			t.Fatalf("case %v: result: expected: %v, got: %v", i+1, testCase.output, output)
		}
	}

}
