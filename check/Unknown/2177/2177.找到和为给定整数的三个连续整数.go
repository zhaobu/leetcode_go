package main

/*
 * @lc app=leetcode.cn id=2177 lang=golang
 *
 * [2177] 找到和为给定整数的三个连续整数
 */

// @lc code=start
/*
解法1 数学
3+4+5=12
*/
func sumOfThree(num int64) []int64 {
	ret := []int64{}
	if num%3 == 0 {
		ret = []int64{num/3 - 1, num / 3, num/3 + 1}
	}
	return ret
}

// @lc code=end
