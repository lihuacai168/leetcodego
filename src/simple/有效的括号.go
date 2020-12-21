package simple

func isLeft(i byte) bool {
	if i == '(' || i == '[' || i == '{' {
		return true
	}
	return false
}

func isMatch(left, right byte) bool {
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
	var stack []byte
	for _, i := range s {
		if isLeft(byte(i)) {
			stack = append(stack, byte(i))
		} else {
			n := len(stack)
			// 栈为空，遇到右括号，肯定为false
			if n == 0 {
				return false
			}
			stackTop := stack[n-1]
			// 栈顶的左括号是否和当前的右括号匹配
			if isMatch(stackTop, byte(i)) {
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
