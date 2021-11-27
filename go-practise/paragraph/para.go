package main

import (
	"fmt"
	"sort"
	"strings"
)

// paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
// banned = ["hit"]
// Output: "ball"
// Explanation:
// "hit" occurs 3 times, but it is a banned word.
// "ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph.

var m map[string]int

func main() {
	banned := []string{"hit"}
	m := make(map[string]int)
	para := "Bob hit a ball, the hit BALL flew far after it was it hit."
	for _, val := range strings.Split(para, ",") {
		val = strings.Trim(strings.Trim(val, "."), " ")
		for _, word := range strings.Split(val, " ") {
			val, ok := m[word]
			if ok {
				m[word] = val + 1
			} else {
				m[word] = 1
			}
		}
	}
	keys := make([]string, 0, len(m))
	for name := range m {
		keys = append(keys, name)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	fmt.Print(m)

	for _, key := range keys {
		_, ok := contains(banned, key)
		if ok {
			continue
		} else {
			fmt.Println("\n Answer : \t ", key, "\t count : \t ", m[key])
			break
		}
	}
}

func contains(sts []string, s string) (int, bool) {
	for i, v := range sts {
		if v == s {
			return i, true
		}
	}
	return -1, false
}
