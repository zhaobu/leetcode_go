package main

/*
 *
 *
 * 剑指 Offer 17. 打印从1到最大的n位数
 */

// @lc code=start

/*
 解法1
*/
func printNumbers(n int) []int {
	//求出最大数
	max := 1
	for i := 0; i < n; i++ {
		max *= 10
	}
	ret := make([]int, 0, max)
	for i := 1; i < max; i++ {
		ret = append(ret, i)
	}
	return ret
}

// @lc code=end
