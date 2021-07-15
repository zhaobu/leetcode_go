/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
// package check

// @lc code=start
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			stack = append(stack, s[i])
		} else {
			if len(stack) < 1 { //栈为空,没有左括号
				return false
			}
			//出栈
			left := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if !isMatch(left, s[i]) {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isMatch(left, right byte) bool {
	switch right {
	case ')':
		return left == '('
	case '}':
		return left == '{'
	case ']':
		return left == '['
	}
	return false
}

func isLeft(ch byte) bool {
	return (ch == '(') || (ch == '{') || (ch == '[')
}

// @lc code=end
