// 1711. Count Good Meals
// A good meal is a meal that contains exactly two different food items with a sum of deliciousness equal to a power of two.

// You can pick any two different foods to make a good meal.

// Given an array of integers deliciousness where deliciousness[i] is the deliciousness of the i​​​​​​th​​​​​​​​ item of food, return the number of different good meals you can make from this list modulo 109 + 7.

// Note that items with different indices are considered different even if they have the same deliciousness value.

// Example 1:

// Input: deliciousness = [1,3,5,7,9]
// Output: 4
// Explanation: The good meals are (1,3), (1,7), (3,5) and, (7,9).
// Their respective sums are 4, 8, 8, and 16, all of which are powers of 2.

package main

import (
	"strconv"
)

func countPairs(deliciousness []int) int {
	ans := 0
	for i := 0; i < len(deliciousness)-1; i++ {
		for j := i + 1; j < len(deliciousness); j++ {
			val := deliciousness[j] - deliciousness[i]
			// if val && (!(strconv.FormatInt(val, 2) & (strconv.FormatInt(int64(val)-1, 2)))) {
			if (strconv.FormatInt(int64(val), 2) & (strconv.FormatInt(int64(val), 2) - strconv.FormatInt(1, 2))) == 0 {

			}
			// if (val && val-1) == 0 {
			// 	ans++
			// }
		}
	}
	return ans
}
