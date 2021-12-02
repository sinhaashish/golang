//Given an array of integers nums and an integer target, return indices of
// the two numbers such that they add up to target.
//You may assume that each input would have exactly one solution,
//and you may not use the same element twice. You can return the answer in any order.

//Input: nums = [2,7,11,15], target = 9
//Output: [0,1]
//Output: Because nums[0] + nums[1] == 9, we return [0, 1].

package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	hashMap := make(map[int]int)

	for i, k := range nums {
		if containsKey(hashMap, target-k) {
			result = append(result, i)
			result = append(result, hashMap[target-k])
			return result
		}
		hashMap[k] = i
	}
	return result
}

func containsKey(m map[int]int, val int) bool {
	if _, ok := m[val]; ok {
		return true
	}
	return false
}
