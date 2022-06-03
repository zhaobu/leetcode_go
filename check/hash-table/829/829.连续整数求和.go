package main

/*
 * @lc app=leetcode.cn id=829 lang=golang
 *
 * [829] 连续整数求和
 */

// @lc code=start

/*
解法1 暴力法,超时
*/
func consecutiveNumbersSum1(n int) int {
	if n < 3 {
		return 1
	}
	m := n>>1 + 1
	preSum := make([]int, m+1)
	for i := 0; i < m; i++ {
		preSum[i+1] = preSum[i] + i + 1
	}

	/*
	   1. 求[i, j) pre[j]-pre[i]
	   2. nums= [1,2,3), preSum=[0,1,3,6]
	*/
	ret := 1
	lastJ, nextJ := m, 0
	for i := m - 2; i >= 0; i-- {
		nextJ = 0
		for j := lastJ; j > 1; j-- {
			cur := preSum[j] - preSum[i]
			if cur == n {
				ret++
			} else if cur < n {
				if nextJ == 0 {
					nextJ = j
				}
				break
			}
		}
		lastJ = nextJ
	}

	return ret
}

/*
解法2 数学
如果正整数 n 可以表示成 k 个连续正整数之和，则由于 k 个连续正整数之和的最小值是
sum = (1+k)*k/2, 因此有 n>=(1+k)*k/2,也就是(k+1)*k<=2n.
​枚举每个符合(k+1)*k<=2n的正整数 k，判断正整数 n 是否可以表示成 k 个连续正整数之和。

1. 如果 k 是奇数，则当 n 可以被 k 整除时，正整数 n 可以表示成 k 个连续正整数之和；
2. 如果 k 是偶数，则当 n 不可以被 k 整除且 2n 可以被 k 整除时，正整数 n 可以表示成 k 个连续正整数之和。
*/
func consecutiveNumbersSum(n int) (ans int) {
	//判断n是否能被k整除
	kDivided := func(k int) bool {
		if k&1 > 0 {
			return n%k == 0
		}
		return n%k != 0 && (2*n)%k == 0
	}

	for k := 1; (k+1)*k <= 2*n; k++ {
		if kDivided(k) {
			ans++
		}
	}

	return ans
}

// @lc code=end
