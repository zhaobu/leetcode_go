package main

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start

/*
解法1
递归
解法超时,只为展示该种思路
*/
func longestCommonSubsequence1(text1 string, text2 string) int {
	if len(text1) < 1 || len(text2) < 1 {
		return 0
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var lcs func(s1, s2 []byte) int
	lcs = func(s1, s2 []byte) int {
		if len(s1) < 1 || len(s2) < 1 {
			return 0
		}
		l1, l2 := len(s1)-1, len(s2)-1
		if s1[l1] == s2[l2] {
			return lcs(s1[:l1], s2[:l2]) + 1
		}
		return max(lcs(s1[:l1], s2), lcs(s1, s2[:l2]))
	}

	return lcs([]byte(text1), []byte(text2))
}

/*
解法2
动态规划

状态定义:
dp[j][j] 表示 text1 的前 i 个元素和 text2 的前 j 个元素的最长公共子序列长度

初始状态:
当 i=0 或者 j=0 时,表示一个字符串和空串的最长公共子序列为 0, 也就是dp[i][j]=0

状态转移方程:
1. 当 text1[i] == text2[j] 时,dp[i][j]=dp[i-1][j-1]+1
比如 text1=abcd 和 text2=bfd时 dp[4][3]对应的最长公共子序列是 bd
dp[3][2]对应的最长公共子序列是 b, 加上最后一个公共字符 d 就是dp[4][3]
2. 当 text1[i] != text2[j] 时, dp[i][j]=max(dp[i-1][j],dp[i][j-1])
比如 text1=abcd 和 text2=agc 时 dp[4][3]对应的最长公共子序列是 ac
dp[4][2]对应的最长公共子序列是 a
dp[3][3]对应的最长公共子序列是 ac

*/
func longestCommonSubsequence2(text1 string, text2 string) int {
	if len(text1) < 1 || len(text2) < 1 {
		return 0
	}

	//dp[i][j] 表示 text1 的前 i 个元素和 text2 的前 j 个元素的最长公共子序列长度
	dp := make([][]int, len(text1)+1)

	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//dp[i][0] 和dp[0][j] 都为0,不用初始化
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

/*
解法3
动态规划
观察发现dp[i][j]的结果计算只需要上一行
所以二维数组可以减少的到2行

*/
func longestCommonSubsequence3(text1 string, text2 string) int {
	if len(text1) < 1 || len(text2) < 1 {
		return 0
	}

	/*
		dp[i][j] 表示 text1 的前 i 个元素和 text2 的前 j 个元素的最长公共子序列长度
		用滚动二维数组做dp table
	*/
	dp := make([][]int, 2)

	for i, _ := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	//dp[i][0] 和dp[0][j] 都为0,不用初始化
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			curRow := i & 1
			preRow := (i - 1) & 1
			if text1[i-1] == text2[j-1] {
				dp[curRow][j] = dp[preRow][j-1] + 1
			} else {
				dp[curRow][j] = max(dp[preRow][j], dp[curRow][j-1])
			}
		}
	}
	return dp[len(text1)&1][len(text2)]
}

/*
解法4
动态规划
继续观察发现 计算时,text1和text2的长度谁做行谁做列都可以计算出结果,
所以可以考虑长度更小的做列,这样空间复杂度继续降低

继续观察发现计算dp[i][j]时,假设dp二维数组 i 向下增大, j 向右增大
dp[i][j] = dp[i-1][j-1] + 1 和当前元素的左上角有关,
dp[i][j] = max(dp[i-1][j], dp[i][j-1]) 和当前元素的左方元素和上方元素有关.

申请一维数组dp[i]代替二位数组 ,在计算dp[i][j]时按如下步骤:
1. 用变量 leftTop 记录dp[i],用作下次计算时的左上方
2. 根据条件计算出 dp[i][j] 放在 dp[i],这样 leftTop,dp[i],dp[i+1] 就分别
是下一位待计算的dp[i][j+1]的左上,左,上的三个值

*/
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < 1 || len(text2) < 1 {
		return 0
	}

	// 求较小的长度作为列
	row, col := text1, text2
	if len(text1) < len(text2) {
		row, col = text2, text1
	}
	/*
		dp[i][j] 表示 text1 的前 i 个元素和 text2 的前 j 个元素的最长公共子序列长度
		用一维数组代替, dp[0,i] 存储的是当前行的最长公共子序列长度,
		而dp[i+1,len]还是上一行的最长公共子序列长度
	*/
	dp := make([]int, len(col)+1)

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i <= len(row); i++ {
		nextLeftTop := 0 //每次换到新行,都要进行重置
		for j := 1; j <= len(col); j++ {
			leftTop := nextLeftTop //当前位置的左上角元素是上一次记录的
			nextLeftTop = dp[j]    //记录下一次的左上角元素
			if row[i-1] == col[j-1] {
				dp[j] = leftTop + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
		}
	}
	return dp[len(col)]
}

// @lc code=end
