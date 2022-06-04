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
参考注释: https://www.yuque.com/docs/share/b6e2bd85-31a7-41ab-8e68-9bf04bb728b7?# 《大数问题（循环求余/快速幂求余/快速幂问题）》
比如求 2^7:
2^6*2 = (2^2 * 2^2 * 2^2) *2
算法计算步骤
1. n=7,  ret-> 2,                     x-> 2^2,               n->3
2. n=3,  ret-> 2* 2^2,                x-> 2^2*2^2=2^4,       n->1
3. n=1,  ret-> 2* 2^2 * (2^2*2^2),    x-> 2^8,               n->0
*/
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	pow := func(x float64, n int) (ret float64) {
		ret = 1.0
		cur := x    // 贡献的初始值为 x
		for n > 0 { // 在对 n 进行二进制拆分的同时计算答案
			if n&1 == 1 { //如果n的二进制位最低位为1,需要把当前的cur计入贡献值
				ret *= cur
			}
			cur *= cur // 将贡献不断地平方
			n >>= 1    // 去掉n的二进制数的最低位，这样我们每次只要判断最低位即可
		}
		return ret
	}

	if n < 0 {
		return 1 / pow(x, -n)
	} else {
		return pow(x, n)
	}
}

// @lc code=end
