package main

import (
	"math"
	"sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	hs := make(map[int]int)
	for i := 0; i < len(difficulty); i++ {
		hs[difficulty[i]] = profit[i]
	}
	sort.Ints(difficulty)
	sum, bestProfit := 0, 0
	for _, val := range worker {
		for j, diff := range difficulty {
			if diff <= val {
				bestProfit = int(math.Max(float64(bestProfit), float64(hs[difficulty[j]])))
			}
			sum += bestProfit
		}
	}
	return sum
}
