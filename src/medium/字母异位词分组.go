package medium

import (
	"sort"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		k := SortString(str)
		if m[k] == nil {
			m[k] = []string{str}
		} else {
			m[k] = append(m[k], str)
		}
	}
	ans := make([][]string, len(m))
	i := 0
	for _, strings := range m {
		ans[i] = strings
		i++
	}
	return ans
}
