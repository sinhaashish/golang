// Given the array nums consisting of 2n elements in the form [x1,x2,...,xn,y1,y2,...,yn].

// Return the array in the form [x1,y1,x2,y2,...,xn,yn].

package main

import "fmt"

func main() {

	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	fmt.Println(shuffle(nums, n))

}
func shuffle(nums []int, n int) []int {
	result := make([]int, 2*n)
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		result[j] = nums[i]
		result[j+1] = nums[i+n]
	}
	return result
}
