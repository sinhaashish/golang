// You are given a 0-indexed string s and a 0-indexed integer array spaces that describes the indices in the original string where spaces will be added. Each space should be inserted before the character at the given index.

// For example, given s = "EnjoyYourCoffee" and spaces = [5, 9], we place spaces before 'Y' and 'C', which are at indices 5 and 9 respectively. Thus, we obtain "Enjoy Your Coffee".
// Return the modified string after the spaces have been added.

package main

import "strings"

func addSpaces(s string, spaces []int) string {
	var ans []string
	ans = append(ans, s[:spaces[0]])
	for i := 0; i < len(spaces)-1; i++ {
		ans = append(ans, s[spaces[i]:spaces[i+1]])
	}
	ans = append(ans, s[spaces[len(spaces)-1]:])
	return strings.Join(ans, " ")
}
