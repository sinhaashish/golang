// You are given an integer array prices representing the daily price history of a stock, where prices[i] is the stock price on the ith day.

// A smooth descent period of a stock consists of one or more contiguous days such that the price on each day is lower than the price on the preceding day by exactly 1. The first day of the period is exempted from this rule.

// Return the number of smooth descent periods.

package main

func getDescentPeriods(prices []int) int {
	ans := 0
	curr := 1
	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1]-1 {
			curr += 1
		} else {
			ans += (curr * (curr + 1)) / 2
			curr = 1
		}
	}
	return ans + (curr*(curr+1))/2
}
