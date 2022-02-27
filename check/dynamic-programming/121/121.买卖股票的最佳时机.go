package main

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
// @lc code=start

/*
解法2
动态规划
1.dp[i][j]表示第i天买第j天卖的利润,也就是prices[j] - price[i]
2.计算dp[i][j]可以转换为计算dp[i][i+1]+dp[i+1][i+2]+ ... + dp[j-1][j]
3.通过2的转换后,就变成了求最大连续子序列和的问题
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	// 求出所有的dp[i][i+1],也就是第i天买,第i+1天卖的值
	for i := 1; i < len(prices); i++ {
		prices[i-1] = prices[i] - prices[i-1]
	}

	/*
		1.dp表示转换为最大连续子序列问题后的dp
		2.dp[i]表示以i结尾的最大连续子序列的值
		3.因为dp[i]只和 dp[i-1]+price[i]和price[i]相关,
		所以可以直接优化空间复杂度到1,用一个dp变量代替dp数组
	*/
	dp := 0
	ret := 0 //记录求得的最大利润
	for i := 0; i < len(prices)-1; i++ {
		sum := dp + prices[i]
		if sum > prices[i] { //包括之前的子序列
			dp = sum
		} else { //不包括之前的子序列
			dp = prices[i]
		}
		if dp > ret {
			ret = dp
		}
	}
	return ret
}

/*
解法1
遍历一遍,记录最低的股票价格, 每天都尝试卖.
如果能获取的利润更大就更新最大利润记录
即可求最大利润
*/
func maxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0] //记录已遍历过的最低价格
	ret := 0              //记录最大利润
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice { //更新历史最低股票价格
			minPrice = prices[i]
		} else {
			sale := prices[i] - minPrice
			if sale > ret { //把今天的股票卖掉的利润和历史最大利润比价
				ret = sale
			}
		}
	}
	return ret
}

// @lc code=end
