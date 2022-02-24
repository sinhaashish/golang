package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	var res []int
	var stack []Day
	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			res = append(res, -1)
		} else if len(stack) > 0 && stack[len(stack)-1].val > temperatures[i] {
			res = append(res, stack[len(stack)-1].index)
		} else if len(stack) > 0 && stack[len(stack)-1].val <= temperatures[i] {
			for len(stack) > 0 && stack[len(stack)-1].val <= temperatures[i] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				res = append(res, -1)
			} else {
				res = append(res, stack[len(stack)-1].index)
			}
		}
		stack = append(stack, Day{temperatures[i], i})
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	fmt.Println(res)
	for i := 0; i < len(res); i++ {
		if res[i] == -1 {
			res[i] = 0
		} else {
			res[i] = res[i] - i
		}
	}
	fmt.Println(res)
	return res
}

type Day struct {
	val   int
	index int
}
