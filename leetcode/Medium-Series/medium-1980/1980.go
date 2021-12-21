// Find Unique Binary String
// Given an array of strings nums containing n unique binary strings each of length n, return a binary string of length n that does not appear in nums. If there are multiple answers, you may return any of them.

// Example 1:

// Input: nums = ["01","10"]
// Output: "11"
// Explanation: "11" does not appear in nums. "00" would also be correct.

package main

import "strings"

func findDifferentBinaryString(nums []string) string {
	var ans []string
	for i, val := range nums {
		t := []rune(val)
		if t[i] == '1' {
			ans = append(ans, "0")
		} else {
			ans = append(ans, "1")
		}
	}
	return strings.Join(ans, "")
}
