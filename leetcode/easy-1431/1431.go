// There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

// Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

// Note that multiple kids can have the greatest number of candies.

package main

import (
	"fmt"
	"math"
)

func main() {

	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))

}
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := candies[0]
	for _, val := range candies {
		maxCandies = int(math.Max(float64(maxCandies), float64(val)))
	}
	result := make([]bool, len(candies))
	for i, v := range candies {
		result[i] = v+extraCandies >= maxCandies
	}
	return result
}
