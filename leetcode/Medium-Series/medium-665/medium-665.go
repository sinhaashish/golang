// 665. Non-decreasing Array
// Medium

// Share
// Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.

// We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).

// Example 1:

// Input: nums = [4,2,3]
// Output: true
// Explanation: You could modify the first 4 to 1 to get a non-decreasing array.

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 4, 2, 3}
	fmt.Println(CheckPossibility(nums))
}

func CheckPossibility(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	modified := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		if modified {
			return false
		}
		if i > 0 && nums[i-1] > nums[i+1] {
			nums[i+1] = nums[i]
		}
		modified = true
		fmt.Println(nums)
	}
	return true
}
