package main

/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start

/*
解法1 位操作
异或计算
1. a^a=0
2. a^0=a
3. a^b^c = (a^b)^c = a^(b^c) = (a^c)^b
*/
func singleNumber(nums []int) int {
	ret := 0
	for i := range nums {
		ret ^= nums[i]
	}
	return ret
}

// @lc code=end
