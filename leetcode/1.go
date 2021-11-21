package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	map1 := make(map[int]int)
	for i, k := range nums {
		if containsKey(map1, target-k) {
			result = append(result, i)
			result = append(result, map1[target-k])
			return result
		}
		map1[k] = i
	}
	return result
}

func containsKey(m map[int]int, val int) bool {
	if _, ok := m[val]; ok {
		return true
	}
	return false
}
