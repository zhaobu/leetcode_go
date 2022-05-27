package main

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2 的幂
 */

// @lc code=start

/*
解法1 位运算
1. 2的n次幂二进制数只有1个位置上为1
2. 使用n&(n-1)不断消掉最后一位置上的1
*/
func isPowerOfTwo(n int) bool {
	ret := 0
	for n > 0 {
		ret++
		n &= n - 1
	}
	return ret == 1
}

// @lc code=end
