package simple

func firstUniqChar(s string) int {
	charCount := [26]int{}
	ans := -1
	for _, char := range s {
		charCount[char-'a']++
	}

	for idx, char := range s {
		if charCount[char-'a'] == 1 {
			return idx
		}
	}
	return ans
}
