package main

/*
 * @lc app=leetcode.cn id=1175 lang=golang
 *
 * [1175] 质数排列
 */

// @lc code=start

/*
解法1
*/
func numPrimeArrangements(n int) int {
	//计算多少个质数
	primeCount := 0
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		if isPrime[i-1] {
			primeCount++
			for j := i; j*i <= n; j++ {
				isPrime[j*i-1] = false
			}
		}
	}

	const mod = 1e9 + 7
	// 计算阶乘
	getN := func(a int) (b int) {
		b = 1
		for i := a; i > 1; i-- {
			b = b * i % mod
		}
		return b
	}
	// fmt.Printf("primeCount=%d, c=%d\n", primeCount, getN(primeCount))
	return getN(primeCount) * getN(n-primeCount) % mod
}

// @lc code=end
