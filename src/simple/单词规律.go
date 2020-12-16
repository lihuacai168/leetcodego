package simple

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	patterMap := make(map[uint8]int)
	wordsMap := make(map[string]int)

	n := len(pattern)
	words := strings.Split(s, " ")
	if n != len(words) {
		return false
	}
	for i := 0; i < n; i++ {
		patterMap[pattern[i]] += i
		wordsMap[words[i]] += i
	}

	for i := 0; i < n; i++ {
		if patterMap[pattern[i]] != wordsMap[words[i]] {
			return false
		}
	}
	return true
}
