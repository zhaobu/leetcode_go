package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (44.54%)
 * Likes:    2840
 * Dislikes: 0
 * Total Accepted:    869.4K
 * Total Submissions: 2M
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "()"
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "()[]{}"
 * 输出：true
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "(]"
 * 输出：false
 *
 *
 * 示例 4：
 *
 *
 * 输入：s = "([)]"
 * 输出：false
 *
 *
 * 示例 5：
 *
 *
 * 输入：s = "{[]}"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * s 仅由括号 '()[]{}' 组成
 *
 *
 */

// @lc code=start
// 解法1:
func isValid1(s string) bool {
	if len(s)&1 == 1 { //奇数个元素肯定不匹配
		return false
	}
	var (
		stack = make([]byte, 0, len(s)>>1) //默认分配s一半长的空间
	)

	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) { //左括号入栈
			stack = append(stack, s[i])
		} else { //右括号,出栈,并与当前右括号匹配
			if len(stack) < 1 {
				return false
			}
			left := stack[len(stack)-1]
			if !isMatch(left, s[i]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func isLeft(ch byte) bool {
	return ch == '(' || ch == '{' || ch == '['
}

func isMatch(left, right byte) bool {
	return left == '(' && right == ')' || left == '{' && right == '}' || left == '[' && right == ']'
}

// 解法2
func isValid(s string) bool {
	if len(s)&1 == 1 { //奇数个元素肯定不匹配
		return false
	}
	var (
		stack = make([]byte, 0, len(s)>>1) //默认分配s一半长的空间
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
		} else if s[i] == '{' {
			stack = append(stack, '}')
		} else if s[i] == '[' {
			stack = append(stack, ']')
			// 拿出当前的右括号和栈顶的元素比较
		} else if len(stack) < 1 || stack[len(stack)-1] != s[i] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// @lc code=end
