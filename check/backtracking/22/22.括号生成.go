package main

import "fmt"

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

/*
解法1 暴力法
1. 生成n对基本字符chars
2. 利用生成的基本字符chars生成所有不重复的全排列情况
3. 判断有效的排列
*/
func generateParenthesis1(n int) []string {
	ret := []string{}
	if n < 1 {
		return ret
	}

	//生成用来组合挂号的字符
	chars := make([]byte, 0, 2*n)
	for i := 0; i < n; i++ {
		chars = append(chars, []byte{'(', ')'}...)
	}

	// fmt.Printf("chars=%+v\n", string(chars))
	n = 2 * n

	//验证组合是否有效
	var valid func(picked []byte) bool
	valid = func(picked []byte) bool {
		count := 0
		for _, v := range picked {
			if v == '(' {
				count++
			} else {
				count--
			}
			if count < 0 {
				return false
			}
		}
		return count == 0
	}

	vis := make([]bool, n)

	var repeated func(i int) bool
	repeated = func(i int) bool {
		for j := 0; j <= i; j++ {
			if vis[j] && chars[j] == chars[i] {
				return true
			}
		}
		return false
	}

	//利用基本字符生成不重复的全排列
	var dfs func(picked []byte)
	dfs = func(picked []byte) {
		if len(picked) == n {
			fmt.Printf("picked=%+v\n", string(picked))
			if valid(picked) {
				// fmt.Printf("picked=%+v\n", string(picked))
				ret = append(ret, string(picked))
			}
			return
		}

		for i := 0; i < n; i++ {
			if repeated(i) {
				continue
			}
			vis[i] = true
			dfs(append(picked, chars[i]))
			vis[i] = false
		}
	}

	dfs(make([]byte, 0, n))

	return ret
}

/*
解法2 暴力法 递归版
*/
func generateParenthesis2(n int) []string {
	ret := []string{}
	if n < 1 {
		return ret
	}

	n = 2 * n

	//验证组合是否有效
	var valid func(picked []byte) bool
	valid = func(picked []byte) bool {
		count := 0
		for _, v := range picked {
			if v == '(' {
				count++
			} else {
				count--
			}
			if count < 0 {
				return false
			}
		}
		return count == 0
	}

	//递归生成所有可能的组合
	var dfs func(picked []byte)
	dfs = func(picked []byte) {
		if len(picked) == n {
			// fmt.Printf("picked=%+v\n", string(picked))
			if valid(picked) {
				// fmt.Printf("picked=%+v\n", string(picked))
				ret = append(ret, string(picked))
			}
			return
		}
		dfs(append(picked, '('))
		dfs(append(picked, ')'))
	}

	dfs(make([]byte, 0, n))

	return ret
}

/*
解法3 回溯法方法2的剪枝优化版
*/
func generateParenthesis(n int) []string {
	ret := []string{}
	if n < 1 {
		return ret
	}

	m := 2 * n
	left, right := 0, 0 //用来记录左右括号的数量
	//递归生成所有可能的组合
	var dfs func(picked []byte)
	dfs = func(picked []byte) {
		if len(picked) == m {
			ret = append(ret, string(picked))
			return
		}
		if left < n { //左括号的数量不能超过括号对数
			left++
			dfs(append(picked, '('))
			left--
		}
		if right < left { //右括号数量不能超过左括号
			right++
			dfs(append(picked, ')'))
			right--
		}
	}

	dfs(make([]byte, 0, n))

	return ret
}

// @lc code=end
