package main

import "strconv"

/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start

/*
解法1 分治
*/
func diffWaysToCompute(expression string) []int {

	record := map[string][]int{}

	doDiffWaysToCompute := func(expression string) []int {
		if ret, ok := record[expression]; ok {
			return ret
		}

		ret := []int{}

		for i, ch := range expression {
			if ch == '+' || ch == '-' || ch == '*' {
				left := diffWaysToCompute(expression[:i])
				right := diffWaysToCompute(expression[i+1:])

				for _, num1 := range left {
					for _, num2 := range right {
						if ch == '+' {
							ret = append(ret, num1+num2)
						} else if ch == '-' {
							ret = append(ret, num1-num2)
						} else if ch == '*' {
							ret = append(ret, num1*num2)
						}
					}
				}
			}
		}
		if len(ret) == 0 {
			num, _ := strconv.Atoi(expression)
			ret = append(ret, num)
		}
		record[expression] = ret
		return ret
	}

	return doDiffWaysToCompute(expression)
}

// @lc code=end
