package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start

/*
解法1 栈
1. 如果当前的字符为字母或者左括号，或者数字, 直接进栈
2. 如果当前的字符为右括号，开始出栈，一直到左括号出栈，出栈序列反转后拼接成一个字符串，
这个字符串就是要拼接的字符串，继续出栈,出栈序列反转后拼接成一个字符串, 直到栈为空或者栈顶元素不为数字,
然后把拼接出的全是数字的字符串转换成整数,就是这个字符串应该出现的次数，
我们根据这个次数和字符串构造出新的字符串并进栈
3. 最后整个栈就是解析出来的字符串
*/
func decodeString1(s string) string {
	m := len(s)
	if m < 2 {
		return s
	}

	getType := func(ch byte) int {
		if ch >= 'a' && ch <= 'z' {
			return 0
		} else if ch >= '0' && ch <= '9' {
			return 1
		} else if ch == '[' {
			return 2
		} else {
			return 3
		}
	}

	stack := []byte{}

	for i := 0; i < m; i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {
			curStr := "" //当前要重复的字符串
			numStr := "" //要重复的次数
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				t := getType(top)
				if t == 0 { //字母
					curStr = string(top) + curStr
				} else if t == 1 { //数字
					numStr = string(top) + numStr
					/*
						判断下一个栈顶元素是否还是数字,如果不是则
						已经计算出要重复的字符串和要重复的次数
					*/
					if len(stack) == 0 || getType(stack[len(stack)-1]) != 1 {
						n, _ := strconv.Atoi(numStr)
						for i := 0; i < n; i++ {
							stack = append(stack, []byte(curStr)...)
						}
						//解析出一部分后,继续下一部分
						break
					}
				}
			}
		}
	}

	return string(stack)
}

/*
解法2 递归
*/
func decodeString(s string) string {
	m := len(s)
	if m < 2 {
		return s
	}

	getType := func(ch byte) int {
		if ch >= 'a' && ch <= 'z' {
			return 0
		} else if ch >= '0' && ch <= '9' {
			return 1
		} else if ch == '[' {
			return 2
		} else {
			return 3
		}
	}

	i := 0
	var getString func() string
	getString = func() string {
		if i >= m || s[i] == ']' {
			return ""
		}

		ret := ""
		ch := s[i]
		t := getType(ch)
		i++
		if t == 0 { //字符
			ret = string(ch)
		} else if t == 1 { //数字
			n := int(ch - '0')
			for ; i < m && getType(s[i]) == 1; i++ {
				n = n*10 + int(s[i]-'0')
			}
			i++ //过滤掉[
			str := getString()
			i++ //过滤掉]
			for j := 0; j < n; j++ {
				ret += str
			}
		}

		return ret + getString()
	}

	return getString()
}

// @lc code=end
