// A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each word consists of lowercase and uppercase English letters.

// A sentence can be shuffled by appending the 1-indexed word position to each word then rearranging the words in the sentence.

// For example, the sentence "This is a sentence" can be shuffled as "sentence4 a3 is2 This1" or "is2 sentence4 This1 a3".
// Given a shuffled sentence s containing no more than 9 words, reconstruct and return the original sentence

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "is2 sentence4 This1 a3"
	fmt.Println(s)
	fmt.Println(sortSentence(s))
}

func sortSentence(s string) string {
	words := strings.Fields(s)
	sort.Slice(words, func(i, j int) bool {
		return words[i][len(words[i])-1] < words[j][len(words[j])-1]
	})
	for i, w := range words {
		words[i] = w[:len(w)-1]
	}
	return strings.Join(words, " ")
}
