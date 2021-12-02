// Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i]
// and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).

// Specifically, ans is the concatenation of two nums arrays.

// Return the array ans.

package main

import "fmt"

func main() {
	a := []int{1, 2, 4}
	fmt.Println(getConcatenation(a))

}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
