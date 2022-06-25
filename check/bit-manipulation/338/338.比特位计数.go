package main

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start

/*
解法1 动态规划（考虑最低有效位）

计算每个数字i二进制的时候，最低位上是否为1只需要看i % 2的值, 然后左边更高位上的1的个数等于dp[i/2]
1. 当i为偶数时，最低位为0，dp[i] = dp[i/2]
2. 当i为奇数时，最低位为1，dp[i] = dp[i/2]+1
*/
func countBits1(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = i&1 + dp[i/2]
	}
	return dp
}

/*
解法2 动态规划（考虑最高有效位）
0 --> 0
1 --> 1

2 --> 10
3 --> 11

4 --> 100
5 --> 101
6 --> 110
7 --> 111

8 -->  1000
9 -->  1001
10 --> 1010
11 --> 1011
12 --> 1100
13 --> 1101
14 --> 1110
15 --> 1111

16 --> 10000

1. 观察可以发现i从1遍历到n的过程中，每次当i等于2的整数次幂时，二进制数的最高位1就会左移1位，
二进制数会变多1位，同时二进制数除最高位1外，其余数字的1的个数是已知的.
2. 如何求出最高位1右边的数字呢，显然就是i-higBit

*/

func countBits2(n int) []int {
	dp := make([]int, n+1)
	higBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			//如果消掉最低位的1后变成了0，说明i为2的整数次幂
			higBit = i
		}
		dp[i] = dp[i-higBit] + 1
	}
	return dp
}

/*
解法3 动态规划（考虑最低位的1）
1. i&(i-1)可以消掉最低位的1, 除了0之外，每个数的二进制肯定至少包含1个1
2. 消掉最低位的1后的数肯定比i小，所以dp[i] = dp[i&(i-1)] + 1
*/

func countBits3(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i&(i-1)] + 1
	}
	return dp
}

/*
解法4 Brian Kernighan 算法
*/

func countBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		count := 0
		for j := i; j > 0; j = j & (j - 1) {
			count++
		}
		dp[i] = count
	}
	return dp
}

// @lc code=end
