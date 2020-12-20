package simple

func isLeft(i int32) bool {
	if i == '(' || i == '[' || i == '{' {
		return true
	}
	return false
}

func isMatch(left, right int32) bool {
	if left == '(' && right == ')' {
		return true
	}
	if left == '[' && right == ']' {
		return true
	}
	if left == '{' && right == '}' {
		return true
	}
	return false
}

func isValid(s string) bool {
	// 只存左括号的栈
	var stack []int32
	for _, i := range s {
		if isLeft(i) {
			stack = append(stack, i)
		} else {
			n := len(stack)
			// 栈为空，遇到右括号，肯定为false
			if n == 0 {
				return false
			}
			stackTop := stack[n-1]
			// 栈顶的左括号是否和当前的右括号匹配
			if isMatch(stackTop, i) {
				// 左右括号匹配，然后弹栈
				stack = stack[:n-1]
			} else {
				return false
			}
		}
	}
	// 栈不为空，也就是存在左括号没有被匹配
	if len(stack) != 0 {
		return false
	}
	return true
}
