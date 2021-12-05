// You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. You want to know how many of the stones you have are also jewels.

// Letters are case sensitive, so "a" is considered a different type of stone from "A".

// Example 1:

// Input: jewels = "aA", stones = "aAAbbbb"
// Output: 3

package main

import "fmt"

func main() {
	jewels, stones := "aA", "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels, stones))
	fmt.Println(numJewelsInStones1(jewels, stones))

}

func numJewelsInStones(jewels string, stones string) int {
	tmp := make([]int, 52)
	ans := 0
	for _, j := range jewels {
		if j >= 'a' && j <= 'z' {
			tmp[j-'a']++
		}
		if j >= 'A' && j <= 'Z' {
			tmp[j-'A'+26]++
		}
	}
	for _, s := range stones {
		if s >= 'a' && s <= 'z' && tmp[s-'a'] == 1 {
			ans++
		}
		if s >= 'A' && s <= 'Z' && tmp[s-'A'+26] == 1 {
			ans++
		}
	}
	return ans
}

func numJewelsInStones1(J string, S string) int {
	cache := make(map[rune]bool)
	for _, r := range J {
		cache[r] = true
	}
	result := 0
	for _, r := range S {
		if _, ok := cache[r]; ok {
			result++
		}
	}
	return result
}
