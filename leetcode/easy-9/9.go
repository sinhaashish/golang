// Given an integer x, return true if x is palindrome integer.

// An integer is a palindrome when it reads the same backward as forward. For example, 121 is palindrome while 123 is not.

// Example 1:

// Input: x = 121
// Output: true

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(" Is 121 palindrome ", isPalindrome(121))
	fmt.Println(" Is -121 palindrome ", isPalindrome(-121))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var buf bytes.Buffer
	temp := x
	for temp > 0 {
		buf.WriteRune(rune(temp%10 + '0'))
		temp /= 10
	}

	v, _ := strconv.Atoi(buf.String())
	return v == x
}
