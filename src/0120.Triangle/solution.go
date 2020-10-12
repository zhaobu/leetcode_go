package Solution

/*
[小浩算法 ：动态规划)](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/204.html#_03%E3%80%81go%E8%AF%AD%E8%A8%80%E7%A4%BA%E4%BE%8B)
*/
func MinimumTotal1(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}

	dp := make([][]int, len(triangle)) //包含第i行第j列元素的最小路径和
	dp[0] = []int{triangle[0][0]}

	for i := 0; i < len(triangle)-1; i++ {
		//初始化下一行的dp数组
		dp[i+1] = make([]int, 0, len(triangle[i+1]))
		for j := 0; j <= i; j++ {
			//如果dp[i+1][j]不存在就追加,如果存在就和dp[i][j]+triangle[i+1][j]进行比较,取最小值
			if len(dp[i+1]) <= j {
				dp[i+1] = append(dp[i+1], dp[i][j]+triangle[i+1][j])
			} else {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+triangle[i+1][j])
			}
			//如果dp[i+1][j+1]不存在就追加,如果存在就和dp[i][j]+triangle[i+1][j+1]进行比较,取最小值
			if len(dp[i+1]) <= j+1 {
				dp[i+1] = append(dp[i+1], dp[i][j]+triangle[i+1][j+1])
			} else {
				dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j]+triangle[i+1][j+1])
			}

		}
	}
	//计算dp中最后一行里的最小值
	minAns := dp[len(dp)-1][0]
	for _, v := range dp[len(dp)-1] {
		minAns = min(minAns, v)
	}
	return minAns
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
[小浩算法 ：动态规划](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/204.html#_03%E3%80%81go%E8%AF%AD%E8%A8%80%E7%A4%BA%E4%BE%8B)
*/
func MinimumTotal2(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}

	dp := make([][]int, len(triangle)) //包含第i行第j列元素的最小路径和
	dp[0] = []int{triangle[0][0]}

	for i := 1; i < len(triangle); i++ {
		//初始化下一行的dp数组
		dp[i] = make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				//如果是第一列,直接用dp[i-1][j]+元素的值
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				//如果是最后一列,直接用dp[i-1][j-1]+元素的值
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				//如果是中间的情况,可以求上方和左上方的dp值加上元素的值
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			}
		}
	}
	//计算dp中最后一行里的最小值
	minAns := dp[len(dp)-1][0]
	for _, v := range dp[len(dp)-1] {
		minAns = min(minAns, v)
	}
	return minAns
}

/*
[小浩算法 ：动态规划(优化空间复杂度)](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/204.html#_03%E3%80%81go%E8%AF%AD%E8%A8%80%E7%A4%BA%E4%BE%8B)
*/
func MinimumTotal3(triangle [][]int) int {
	n := len(triangle)
	if n < 1 {
		return 0
	}

	dp := make([]int, len(triangle[n-1])) //上一层累计的最小路径
	dp[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		//存储dp[i-1][j-1]的值
		pre := 0
		//初始化下一行的dp数组
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				pre = dp[j]
				//如果是第一列,直接用dp[i-1][j]+元素的值
				dp[j] = dp[j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				//如果是最后一列,直接用dp[i-1][j-1]+元素的值
				dp[j] = pre + triangle[i][j]
			} else {
				//如果是中间的情况,可以求上方和左上方的dp值加上元素的值
				tmp := min(pre, dp[j])
				pre = dp[j]
				dp[j] = tmp + triangle[i][j]
			}
		}
	}
	//计算dp中最后一行里的最小值
	minAns := dp[0]
	for _, v := range dp {
		minAns = min(minAns, v)
	}
	return minAns
}

/*
[小浩算法 ：动态规划(直接使用triangle存储dp的值,优化空间复杂度)](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/204.html#_03%E3%80%81go%E8%AF%AD%E8%A8%80%E7%A4%BA%E4%BE%8B)
*/
func MinimumTotal4(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}

	for i := 1; i < len(triangle); i++ {
		//初始化下一行的triangle数组
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				//如果是第一列,直接用triangle[i-1][j]+元素的值
				triangle[i][j] = triangle[i-1][j] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				//如果是最后一列,直接用triangle[i-1][j-1]+元素的值
				triangle[i][j] = triangle[i-1][j-1] + triangle[i][j]
			} else {
				//如果是中间的情况,可以求上方和左上方的triangle值加上元素的值
				triangle[i][j] = min(triangle[i-1][j-1], triangle[i-1][j]) + triangle[i][j]
			}
		}
	}
	//计算triangle中最后一行里的最小值
	minAns := triangle[len(triangle)-1][0]
	for _, v := range triangle[len(triangle)-1] {
		minAns = min(minAns, v)
	}
	return minAns
}

// 自底向上解法
func MinimumTotal5(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}
	//从三角形倒数第2行开始
	for i := len(triangle) - 2; i >= 0; i-- {
		//计算第i行,第j列元素的最小路径和
		for j := 0; j < len(triangle[i]); j++ {
			//取下方元素和右下方元素的较小值与元素值相加
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}
