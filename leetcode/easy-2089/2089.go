// You are given a 0-indexed integer array nums and a target element target.

// A target index is an index i such that nums[i] == target.

// Return a list of the target indices of nums after sorting nums in non-decreasing order. If there are no target indices, return an empty list. The returned list must be sorted in increasing order.

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 5, 2, 3}
	target := 2
	fmt.Println(targetIndices(nums, target))
}

func targetIndices(nums []int, target int) []int {
	var result []int
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i, v := range nums {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}
