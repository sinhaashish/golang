// Given an integer array nums and an integer k, return the number of non-empty subarrays that have a sum divisible by k.

// A subarray is a contiguous part of an array.

// Example 1:

// Input: nums = [4,5,0,-2,-3,1], k = 5
// Output: 7
// Explanation: There are 7 subarrays with a sum divisible by k = 5:

package main

func subarraysDivByK(nums []int, k int) int {
	hashMap := make(map[int]int)
	ans := 0
	sum := 0
	rem := 0
	hashMap[0] = 1
	for _, value := range nums {
		sum = sum + value
		rem = sum % k
		if rem < 0 {
			rem = rem + k
		}
		hashVal, ok := hashMap[rem]
		if ok {
			ans = ans + hashVal
			hashMap[rem] = hashVal + 1
		} else {
			hashMap[rem] = 1
		}
	}
	return ans
}
