// Find the Kth Largest Integer in the Array
// You are given an array of strings nums and an integer k. Each string in nums represents an integer without leading zeros.

// Return the string that represents the kth largest integer in nums.

// Note: Duplicate numbers should be counted distinctly. For example, if nums is ["1","2","2"], "2" is the first largest integer, "2" is the second-largest integer, and "1" is the third-largest integer.

package main

import (
	"fmt"
	"strconv"
)

func kthLargestNumber(nums []string, k int) string {
	num := make([]int, len(nums))
	for _, v := range nums {
		value, _ := strconv.Atoi(v)
		num = append(num, value)
	}
	fmt.Println()

	//Use heap
}
