// 1528. Shuffle String
// Easy

// 1004

// 219

// Add to List

// Share
// Given a string s and an integer array indices of the same length.

// The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

// Return the shuffled string.

// Example 1:

// Input: s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"
// Explanation: As shown, "codeleet" becomes "leetcode" after shuffling.

package main

import "fmt"

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}

func restoreString(s string, indices []int) string {
	ans := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		ans[indices[i]] = rune(s[i])
	}
	return string(ans)
}
