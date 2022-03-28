package main

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start

/*
解法1 递归, 二分
*/
func myPow1(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n&1 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}

/*
解法2 迭代
原理参考官方解法2
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	ret := float64(1)

	pow := func(x float64, n int) {
		cur := x
		for n > 0 {
			if n&1 == 1 {
				ret *= cur
			}
			cur *= cur
			n >>= 1
		}
	}

	if n < 0 {
		pow(x, -n)
		ret = 1 / ret
	} else {
		pow(x, n)
	}
	return ret
}

// @lc code=end
