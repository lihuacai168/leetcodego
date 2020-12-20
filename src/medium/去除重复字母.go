package medium

func removeDuplicateLetters(s string) string {
	left := [26]int{}
	// 记录s字符串中每个字符的个数
	for _, ch := range s {
		left[ch-'a']++
	}
	var stack []byte
	// 记录s字符串中的字符是否已经存在stack中
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]
		// 元素已经存在stack中，取反
		// 元素已经存在，结果为false
		if !inStack[ch-'a'] {
			// stack不为空，且stack的最后一个元素大于当前的元素
			// 需要抛弃stack最后的元素，并且记录抛弃掉的元素不在stack中
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				// 当元素只有一个时，并且已经在stack中，不能将它弹出去
				// 因为要保持原本的字符顺序
				if left[last] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[last] = false
			}
			// 当stack的ch大于stack的最后一个元素
			// 把ch加入到stack中
			// 并且记录ch已经在stack中
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		// 已经遍历过的字符串，需要标记一次
		left[ch-'a']--
	}
	return string(stack)
}
