// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
package main

import "math"

func canJump(nums []int) bool {
	maxReach := 0
	for i, val := range nums {
		if maxReach < i {
			return false
		}
		maxReach = int(math.Max(float64(maxReach), float64(i)+float64(val)))
	}
	return true
}
