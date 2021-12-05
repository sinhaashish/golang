// Given an integer array nums and an integer k, return the number of pairs (i, j) where i < j such that |nums[i] - nums[j]| == k.

// The value of |x| is defined as:

// x if x >= 0.
// -x if x < 0.

// Example 1:

// Input: nums = [1,2,2,1], k = 1
// Output: 4
// Explanation: The pairs with an absolute difference of 1 are:
// - [1,2,2,1]
// - [1,2,2,1]
// - [1,2,2,1]
// - [1,2,2,1]

package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))

}

func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if int(math.Abs(float64(nums[i])-float64(nums[j]))) == k {
				count = count + 1
			}
		}
	}
	return count
}
