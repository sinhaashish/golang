// Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

// Each element in the array represents your maximum jump length at that position.

// Your goal is to reach the last index in the minimum number of jumps.

// You can assume that you can always reach the last index.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

package main

import "math"

func jump(nums []int) int {
	var end, farthest, jump int
	for i := 0; i < len(nums)-1; i++ {
		farthest = int(math.Max(float64(farthest), float64(i)+float64(nums[i])))
		if i == end {
			jump++
			end = farthest
		}
	}
	return jump
}
