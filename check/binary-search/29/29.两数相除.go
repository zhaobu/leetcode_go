package main

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start

/*
解法1 二分法
*/
func divide(dividend int, divisor int) int {
	if dividend == -1<<31 {
		if divisor == 1 {
			return dividend
		}
		if divisor == -1 {
			return 1<<31 - 1
		}
	}
	if divisor == -1<<31 {
		if dividend == divisor {
			return 1
		}
		return 0
	}
	if dividend == 0 {
		return 0
	}

	flag := false //如果为true，表示最终结果要取反
	if dividend > 0 {
		dividend = -dividend
		flag = !flag
	}

	if divisor > 0 {
		divisor = -divisor
		flag = !flag
	}

	/*
		1. 快速乘,判断z*y >= x 是否成立
		2. x<0, y<0, z>0

	*/
	quickAdd := func(y, z, x int) bool {
		for cur, base := 0, y; z > 0; z >>= 1 {
			if z&1 > 0 {
				if cur < x-base {
					return false
				}
				cur += base
			}
			if z != 1 {
				if base < x-base {
					return false
				}
				base += base
			}
		}
		return true
	}

	ret := 0
	left, right := 1, 1<<31-1
	for left <= right {
		mid := left + (right-left)>>1
		if quickAdd(divisor, mid, dividend) {
			ret = mid
			if mid == 1<<31-1 {
				break
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if flag {
		return -ret
	}
	return ret
}

// @lc code=end
