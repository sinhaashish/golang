// Given an array of non-negative integers arr, you are initially positioned at start index of the array.
// When you are at index i, you can jump to i + arr[i] or i - arr[i], check if you can reach to any index with value 0.

// Notice that you can not jump outside of the array at any time.
// Input: arr = [4,2,3,0,3,1,2], start = 5
// Output: true
// Explanation:
// All possible ways to reach at index 3 with value 0 are:
// index 5 -> index 4 -> index 1 -> index 3
// index 5 -> index 6 -> index 4 -> index 1 -> index 3

package main

import "container/list"

func canReach(arr []int, start int) bool {
	queue := list.New()
	queue.PushBack(start)
	for queue.Len() > 0 {
		i := queue.Front()
		queue.Remove(i)
		index, _ := i.Value.(int)
		if index-arr[index] >= 0 {
			if arr[index-arr[index]] == 0 {
				return true
			} else {
				if arr[index-arr[index]] > 0 {
					queue.PushBack(index - arr[index])
				}
			}
		}
		if index+arr[index] <= len(arr)-1 {
			if arr[index+arr[index]] == 0 {
				return true
			} else {
				if arr[index+arr[index]] > 0 {
					queue.PushBack(index + arr[index])
				}
			}
		}
		arr[index] = -1
	}
	return false
}
