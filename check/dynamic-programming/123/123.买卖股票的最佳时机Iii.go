package main

/*
 * @lc app=leetcode.cn id=123 lang=golang
 *
 * [123] 买卖股票的最佳时机 III
 */

// @lc code=start

/*
解法1 动态规划
[官方题解解法1的写法](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-iii-by-wrnt/)

状态定义:
在每天结束时,会处于以下5种状态中的其中一种
1. 未进行过任何操作；
2. 只进行过一次买操作；
3. 进行了一次买操作和一次卖操作，即完成了一笔交易；
4. 在完成了一笔交易的前提下，进行了第二次买操作；
5. 完成了全部两笔交易。
第一种状态没利润.所以可以用4个变量分别记录后4种状态

basecase:

状态转移:
关键要理解动态规划的每一个状态,在进行转移时,
当前状态的定义不会改变,改变的只是该种状态下的值

*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buy1 := -prices[0] // 只进行过一次买操作状态的最高利润
	seller1 := 0       //进行了一次买操作和一次卖操作，即完成了一笔交易状态的最高利润
	buy2 := -prices[0] //在完成了一笔交易的前提下，进行了第二次买操作状态的最高利润
	seller2 := 0       //成了全部两笔交易状态的最高利润

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		seller1 = max(seller1, buy1+prices[i])
		buy2 = max(buy2, seller1-prices[i])
		seller2 = max(seller2, buy2+prices[i])
	}

	return seller2
}

// @lc code=end
