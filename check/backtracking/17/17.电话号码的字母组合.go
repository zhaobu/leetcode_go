package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	ret := []string{}
	if len(digits) < 1 {
		return ret
	}
	charMap := []string{
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	var dfs func(i int, strByte []rune)
	dfs = func(i int, strByte []rune) {
		if i == len(digits) {
			ret = append(ret, string(strByte))
			return
		}
		dight, _ := strconv.Atoi(string(digits[i]))
		for _, v := range charMap[dight-2] {
			strByte[i] = v
			// fmt.Printf("v=%s\n", string(v))
			dfs(i+1, strByte)
		}
	}

	dfs(0, make([]rune, len(digits)))

	return ret
}

// @lc code=end
