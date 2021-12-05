// 1614. Maximum Nesting Depth of the Parentheses
// Easy

// 622

// 142

// Add to List

// Share
// A string is a valid parentheses string (denoted VPS) if it meets one of the following:

// It is an empty string "", or a single character not equal to "(" or ")",
// It can be written as AB (A concatenated with B), where A and B are VPS's, or
// It can be written as (A), where A is a VPS.
// We can similarly define the nesting depth depth(S) of any VPS S as follows:

// depth("") = 0
// depth(C) = 0, where C is a string with a single character not equal to "(" or ")".
// depth(A + B) = max(depth(A), depth(B)), where A and B are VPS's.
// depth("(" + A + ")") = 1 + depth(A), where A is a VPS.
// For example, "", "()()", and "()(()())" are VPS's (with nesting depths 0, 1, and 2), and ")(" and "(()" are not VPS's.

// Given a VPS represented as string s, return the nesting depth of s.

package main

import "fmt"

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	fmt.Println(maxDepth(s))

}

func maxDepth(s string) int {
	count, output := 0, 0
	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--
		}
		if count > output {
			output = count
		}
	}
	return output
}
