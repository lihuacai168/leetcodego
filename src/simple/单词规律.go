package simple

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[uint8]int)
	m1 := make(map[string]int)

	n := len(pattern)
	split := strings.Split(s, " ")
	if n != len(split) {
		return false
	}
	for i := 0; i < n; i++ {
		m[pattern[i]] += i
		m1[split[i]] += i
	}

	for i := 0; i < n; i++ {
		if m[pattern[i]] != m1[split[i]] {
			return false
		}
	}
	return true
}
