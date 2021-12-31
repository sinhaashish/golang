package main

import "fmt"

func nextLargerNodes(vals []int) []int {
	result := make([]int, 0)
	stack := make([]int, 0)
	for i := len(vals) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			result = append(result, 0)
		} else if len(stack) > 0 && stack[len(stack)-1] > vals[i] {
			result = append(result, stack[len(stack)-1])
		} else if len(stack) > 0 && stack[len(stack)-1] <= vals[i] {
			for len(stack) > 0 && stack[len(stack)-1] <= vals[i] {
				stack = stack[0 : len(stack)-1]
				fmt.Println(stack)
			}
			if len(stack) == 0 {
				result = append(result, 0)
			} else {
				result = append(result, stack[len(stack)-1])
			}
		}
		stack = append(stack, vals[i])
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result

}
