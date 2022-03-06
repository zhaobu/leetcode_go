package main

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
// @lc code=start

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

/*
解法2
动态规划
1.dp[i][j]表示第i天买第j天卖的利润,也就是prices[j] - price[i]
2.计算dp[i][j]可以转换为计算dp[i][i+1]+dp[i+1][i+2]+ ... + dp[j-1][j]
3.通过2的转换后,就变成了求最大连续子序列和的问题
*/
func maxProfit2(prices []int) int {
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
解法3
解法2的优化写法:https://leetcode.com/problems/best-time-to-buy-and-sell-stock/discuss/39038/Kadane's-Algorithm-Since-no-one-has-mentioned-about-this-so-far-%3A)-(In-case-if-interviewer-twists-the-input)

*/
func maxProfit3(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	/*
		1. maxCur可以看成是以i结尾的最大连续子序列的值
	*/
	maxCur := 0
	maxSoFar := 0
	var max func(a, b int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		maxCur = max(0, maxCur+prices[i]-prices[i-1])
		maxSoFar = max(maxCur, maxSoFar)
	}
	return maxSoFar
}

/*
解法4:方法不通用,扩展思路
1. 使用一个单调递增栈,并且保证每个元素能够进栈并出栈一次
2. 当栈顶元素a遇到更小的元素b时就出栈,即使后面遇到更大的元素c, c-b > c-a,
所以弹出a也不会影响到求最大利润
3. 元素在出栈时表示卖掉该值股票,并且卖掉时,能获取的最大利润就是ascStack[栈顶]-ascStack[栈底]
这样就能统计所有卖掉股票的情况
*/
func maxProfit4(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	ret := 0
	ascStack := []int{}        //单调递增栈
	prices = append(prices, 0) //在尾部新增一个0,保证遍历到这个最后元素时,栈里面的元素能够全部弹出
	for i, v := range prices {
		// fmt.Printf("ascStack=%+v\n", ascStack)
		for len(ascStack) > 0 && ascStack[len(ascStack)-1] > v {
			/*
				栈顶元素遇到更小元素时说明
			*/
			top := ascStack[len(ascStack)-1]
			sale := top - ascStack[0]
			// fmt.Printf("top=%+v,sale=%+v\n", top, sale)
			if sale > ret {
				ret = sale
			}
			ascStack = ascStack[:len(ascStack)-1]
		}
		ascStack = append(ascStack, prices[i])
	}
	return ret
}

/*
解法5 动态规划 股票问题通用解法

状态定义:
dp[i][0]表示第i天,不持有股票时的最大利润
dp[i][1]表示第i天,持有股票时的最大利润

basecase:
dp[0][0] = 0
dp[0][1] = -prices[0]

状态转移方程:
第i天不持有: 1. 第i-1天也不持有,第i天没操作 2. 第i-1天持有,但第i天卖了
dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
第i天持有: 1. 第i-1天也持有,第i天没操作 2. 第i-1天没持有,第i天买了
因为只能进行一次操作,所以第i-1天没持有dp[i-1][0]=0
dp[i][1] = max(dp[i-1][1], -prices[i])

可以参考: https://labuladong.gitee.io/algo/3/27/100/
*/
func maxProfit5(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[len(prices)-1][0]
}

/*
解法6 动态规划 股票问题通用解法
解法5 的状态压缩版,优化空间复杂度
*/
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = -prices[0]

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(prices); i++ {
		dp[0] = max(dp[0], dp[1]+prices[i])
		dp[1] = max(dp[1], -prices[i])
	}

	return dp[0]
}

// @lc code=end
