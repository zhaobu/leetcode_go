package Solution

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
[abuladong:暴力递归](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func coinChange1(coins []int, amount int) int {
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	}
	//initVal 为true时表示res为初始值
	res, initVal := 0, true
	for _, v := range coins {
		sub := coinChange1(coins, amount-v)
		if sub == -1 {
			continue
		}
		if initVal {
			res = sub + 1
			initVal = false
		} else {
			res = min(res, sub+1)
		}
	}
	if initVal {
		return -1
	}
	return res
}

/*
[abuladong:带备忘录的递归](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func coinChange2(coins []int, amount int) int {
	//创建备忘录
	record := make([]int, amount+1)
	return _coinChange2(record, coins, amount)
}

func _coinChange2(record, coins []int, amount int) int {
	if amount < 0 {
		return -1
	} else if amount == 0 {
		return 0
	} else if record[amount] != 0 {
		return record[amount]
	}
	//initVal 为true时表示res为初始值
	res, initVal := 0, true
	for _, v := range coins {
		sub := _coinChange2(record, coins, amount-v)
		if sub == -1 {
			continue
		}
		if initVal {
			res = sub + 1
			initVal = false
		} else {
			res = min(res, sub+1)
		}
	}
	if initVal {
		res = -1
	}
	//将结果加入备忘录
	record[amount] = res
	return res
}

/*
[abuladong:dp 数组的迭代解法](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func coinChange3(coins []int, amount int) int {
	/*
		数组大小为 amount + 1，初始值也为 amount + 1
		因为凑成 amount 金额的硬币数最多只可能等于 amount（全用 1 元面值的硬币），所以初始化为 amount + 1 就相当于初始化为正无穷，便于后续取最小值。
	*/
	dp := make([]int, amount+1)
	dp[0] = 0
	// 自底向上计算出dp[i]
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

/*
Todo [leetcode :dp 贪心 + dfs](https://leetcode-cn.com/problems/coin-change/solution/322-by-ikaruga/)
*/
func coinChange4(coins []int, amount int) int {
	return 0
}
