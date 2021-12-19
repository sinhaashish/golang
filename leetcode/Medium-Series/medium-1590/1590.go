// 1590. Make Sum Divisible by P
// Medium

// Share
// Given an array of positive integers nums, remove the smallest subarray (possibly empty) such that the sum of the remaining elements is divisible by p. It is not allowed to remove the whole array.

// Return the length of the smallest subarray that you need to remove, or -1 if it's impossible.

// A subarray is defined as a contiguous block of elements in the array.

// Example 1:

// Input: nums = [3,1,4,2], p = 6
// Output: 1
// Explanation: The sum of the elements in nums is 10, which is not divisible by 6. We can remove the subarray [4], and the sum of the remaining elements is 6, which is divisible by 6.

package main

func minSubarray(nums []int, p int) int {
	sum := 0
	var rem int
	for _, val := range nums {
		sum += val
	}
	rem = sum % p
	if rem == 0 {
		return 0
	}

	if sum < p {
		return -1
	}
	return subarraysDivByK(nums, rem)
}

func subarraysDivByK(nums []int, k int) int {
	hashMap := make(map[int][]int)
	min := -1
	sum := 0
	rem := 0
	hashMap[0] = []int{1, 0}
	for index, value := range nums {
		sum = sum + value
		rem = sum % k
		if rem < 0 {
			rem = rem + k
		}
		hashVal, ok := hashMap[rem]
		if ok {
			if min < index-hashVal[1] {
				min = index - hashVal[1]
			}
			hashMap[rem] = []int{hashVal[0] + 1, index}
		} else {
			hashMap[rem] = []int{1, index}
		}
	}
	return min
}
