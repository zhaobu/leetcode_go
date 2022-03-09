package main

/*
 * @lc app=leetcode.cn id=516 lang=golang
 *
 * [516] 最长回文子序列
 */

// @lc code=start

/*
解法1 动态规划
dp[i][j] 表示s[i,j]字符串之间的最长回文子序列长度

basecase:
dp[i][i] = 1

转移方程:
如果s[i] == s[j]
dp[i][j] = dp[i+1][j-1] + 2
如果s[i] != s[j]
dp[i][j] = max(dp[i][j-1], dp[i+1][j])

注意i和j的遍历顺序由状态转移方程决定
*/
func longestPalindromeSubseq1(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i][j] s[i,j]之间的字符串的最长回文子序列长度
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- { //由转移方程决定i要从大到小求解
		for j := i + 1; j < n; j++ { //由转移方程决定j要从小到大求解
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
		// fmt.Printf("dp[%d]=%+v\n", i, dp[i])
	}

	return dp[0][n-1]
}

/*
解法2 解法1的状态压缩
dp[4]=[0 0 0 0 1]
dp[3]=[0 0 0 1 1]
dp[2]=[0 0 1 1 3]
dp[1]=[0 1 2 2 3]
dp[0]=[1 2 3 3 4]

类似上图的二维数组,从上到下,从左到右
dp[i][j] 只和dp[i+1][j-1](左上),dp[i][j-1](左)
dp[i+1][j](上)的值相关

所以可以优化为一维数组

*/
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// dp[i][j] s[i,j]之间的字符串的最长回文子序列长度
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- { //由转移方程决定i要从大到小求解
		dp[i] = 1                    //默认初始化dp[i][i]为1
		nextLeftTop := 0             //用来记录下一次计算时的左上元素
		for j := i + 1; j < n; j++ { //由转移方程决定j要从小到大求解
			curLeftTop := nextLeftTop //从上一次计算中拿到本次计算的leftTop
			nextLeftTop = dp[j]       //此时的dp[j]还是上一轮循环的值,记录下来,作为下一次计算的左上
			if s[i] == s[j] {
				dp[j] = curLeftTop + 2
			} else {
				dp[j] = max(dp[j-1], dp[j])
			}
		}
		// fmt.Printf("i=%d,dp=%+v\n", i, dp)
	}

	return dp[n-1]
}

// @lc code=end
