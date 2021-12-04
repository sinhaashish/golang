// Given an integer n and an integer start.

// Define an array nums where nums[i] = start + 2*i (0-indexed) and n == nums.length.

// Return the bitwise XOR of all elements of nums.

package main

import "fmt"

func main() {
	n := 5
	start := 0
	fmt.Println(xorOperation(n, start))
}

func xorOperation(n int, start int) int {
	result := start
	for i := 1; i < n; i++ {
		start = start + 2*1
		result ^= start
	}
	return result
}
