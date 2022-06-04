package main

/*
 *
 *
 * 剑指 Offer 14- II. 剪绳子 II
 */

// @lc code=start

/*
 解法1 快速幂求余
核心思路是：尽可能把绳子分成长度为3的小段，这样乘积最大
1. 如果 n == 2，返回1，如果 n == 3，返回2
2. 如果 n == 4，返回4
3. 如果 n > 4，分成尽可能多的长度为3的小段，每次循环长度n减去3，乘积res乘以3；
最后返回时乘以小于等于4的最后一小段；每次乘法操作后记得取余就行

*/
func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	ret := 1
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}

	//出来循环只有三种情况，分别是n=2、3、4
	return ret * n % 1000000007
}

// @lc code=end
