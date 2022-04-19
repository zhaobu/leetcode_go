package main

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start

/*
解法1 二分查找
找到最后一个mid*mid <= x的mid
*/
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	ret := 0
	start, end := 0, x
	for start <= end {
		mid := start + (end-start)>>1
		cur := mid * mid
		if cur == x {
			return mid
		} else if cur > x {
			end = mid
		} else {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			start = mid
		}
	}

	return ret
}

// @lc code=end
