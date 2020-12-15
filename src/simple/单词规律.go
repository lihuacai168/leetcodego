package simple

import (
	"strconv"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	m := make(map[uint8]string)
	m1 := make(map[string]string)

	n := len(pattern)
	split := strings.Split(s, " ")
	if n != len(split) {
		return false
	}
	for i := 0; i < n; i++ {
		m[pattern[i]] += strconv.Itoa(i)
		m1[split[i]] += strconv.Itoa(i)
	}

	for i := 0; i < n; i++ {
		if m[pattern[i]] != m1[split[i]] {
			return false
		}
	}
	return true
}
