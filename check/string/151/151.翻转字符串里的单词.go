package main

import "fmt"

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
/*
思路:
1.先去除多余的空格
2.翻转整个字符串,然后翻转每个单词
*/
func reverseWords(s string) string {
	//去除多余的空格
	b := make([]byte, 0, len(s))
	preSpace := true //前一个字符是否空格
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			b = append(b, s[i])
			preSpace = false
		} else if !preSpace {
			b = append(b, ' ')
			preSpace = true
		}
	}
	if preSpace {
		b = b[:len(b)-1]
	}
	fmt.Printf("string(b)=[%s]\n", string(b))
	// 翻转字符串
	var swapStr func(b []byte, left, right int)
	swapStr = func(b []byte, left, right int) {
		for i, j := left, right; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return
	}

	swapStr(b, 0, len(b)-1)

	fmt.Printf("string(b)=[%s]\n", string(b))

	begin := 0
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			swapStr(b, begin, i-1)
			begin = i + 1
		}
	}
	//翻转最后一个单词
	swapStr(b, begin, len(b)-1)
	return string(b)
}

// @lc code=end
