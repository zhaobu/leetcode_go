package main

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start

/*
解法1 暴力法
*/
func countPrimes1(n int) int {

	isPrime := func(x int) bool {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	ret := 0
	for i := 2; i < n; i++ {
		if isPrime(i) {
			// fmt.Printf("i=%d\n", i)
			ret++
		}
	}

	return ret
}

/*
解法2 埃氏筛
*/
func countPrimes2(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	ret := 0

	for i := 2; i < n; i++ {
		if isPrime[i] {
			ret++
			/*
				对于小于 i*i的那些数，比如i*（i-1）一定在i = i-1时，处理过
			*/
			for j := i; j*i < n; j++ {
				isPrime[j*i] = false
			}
		}
	}

	return ret
}

/*
解法3 线性筛
每个合数只会被其「最小的质因数」筛去，即每个合数被标记一次
*/
func countPrimes(n int) int {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	primes := []int{}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
		for _, j := range primes {
			if i*j >= n {
				break
			}
			isPrime[i*j] = false
			if i%j == 0 {
				break
			}
		}
	}

	return len(primes)
}

// @lc code=end
