// 1909. Remove One Element to Make the Array Strictly Increasing
// Given a 0-indexed integer array nums, return true if it can be made strictly increasing after removing exactly one element, or false otherwise. If the array is already strictly increasing, return true.

// The array nums is strictly increasing if nums[i - 1] < nums[i] for each index (1 <= i < nums.length).

package main

import "fmt"

func main() {

	nums := []int{1, 2, 10, 5, 7}
	fmt.Println(canBeIncreasing(nums))
}

func canBeIncreasing(nums []int) bool {
	var modified bool
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			continue
		}
		if modified {
			return false
		}
		if i > 0 && nums[i-1] >= nums[i+1] {
			nums[i+1] = nums[i]
		}
		modified = true
	}
	return true
}
