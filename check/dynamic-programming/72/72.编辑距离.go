package main

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 < 1 {
		return n2
	}
	if n2 < 1 {
		return n1
	}
	/*
	   dp[i][j] 表示 word1 的前 i 个字符和 word2 的前 j 个字符的编辑距离
	*/
	dp := make([][]int, n1+1)
	// fmt.Printf("0, dp=%+v\n", dp)

	//初始化二位数组的每一行,并且初始化dp[i][0],也就是从word1到空串word2的编辑距离
	for i := 0; i < n1+1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][0] = i
	}
	// fmt.Printf("1, dp=%+v\n", dp)

	//初始化dp[0][j],也就是从空串word1到word2的编辑距离
	for j := 0; j < n2+1; j++ {
		dp[0][j] = j
	}
	// fmt.Printf("2, dp=%+v\n", dp)

	var min func(a, b, c int) int
	min = func(a, b, c int) int {
		if a > b {
			if b > c {
				return c
			}
			return b
		}
		if a > c {
			return c
		}
		return a
	}

	for i := 1; i < n1+1; i++ {
		for j := 1; j < n2+1; j++ {
			/*
			   先删除 word1[0,i) 的最后一个字符得到 word1[0,i-1) ,
			   然后由 word1[0,i-1) 转换为 word2[0,j).
			*/
			dp1 := dp[i-1][j] + 1

			/*
			   先由word1[0,i)转换为word2[0,j-1),
			   然后在最后插入字符word2[j-1],得到word2[0,j)
			*/
			dp2 := dp[i][j-1] + 1

			/*
			   如果word1[i-1] == word2[j-1],由word1[0,i-1)转换为
			   word2[0,j-1)后不用再做任何操作
			*/
			dp3 := dp[i-1][j-1]
			/*
			   如果word1[i-1] != word2[j-1],先由word1[0,i-1)转换为word2[0,j-1)
			   然后再将word1[i-1]替换为word2[j-1]
			*/
			if word1[i-1] != word2[j-1] {
				dp3 += 1
			}
			dp[i][j] = min(dp1, dp2, dp3)
		}
	}

	// fmt.Printf("3, dp=%+v\n", dp)
	return dp[n1][n2]
}

// @lc code=end
