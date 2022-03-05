package main

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start

/*
解法1 动态规划

[官方题解地址](https://leetcode-cn.com/problems/unique-paths-ii/solution/bu-tong-lu-jing-ii-by-leetcode-solution-2/)
动态规划的题目分为两大类，一种是求最优解类，典型问题是背包问题，另一种就是计数类，比如这里的统计方案数的问题，
它们都存在一定的递推性质。前者的递推性质还有一个名字，叫做 「最优子结构」 ——即当前问题的最优解取决于子问题的最优解，
后者类似，当前问题的方案数取决于子问题的方案数。所以在遇到求方案数的问题时，可以往动态规划的方向考虑

状态定义
dp[i][j] 表示到达obstacleGrid[i][j]的路径条数

basecase
第一行dp[0][j]只能从左往右到达: 如果有一个obstacleGrid[0][k]=1 那从它开始后面的所有的位置都不可达
同理第一列dp[i][0]只能从上往下到达:

转移方程
dp[i][j] = dp[i-1][j] + dp[i][j-1]
一个位置只能从上一行同一列或者同一行前一列位置到达
*/
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 {
		return 0
	}
	//dp[i][j] 表示到达obstacleGrid[i][j]的路径条数
	dp := make([][]int, len(obstacleGrid))
	colLen := len(obstacleGrid[0])

	//初始化列
	for i, _ := range dp {
		dp[i] = make([]int, colLen)
		//初始化dp[i][0]
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else if i > 0 { //如果dp[i-1][0]无法到达,dp[i][0]也无法到达
			dp[i][0] = dp[i-1][0]
		} else { //dp[0][0]
			dp[i][0] = 1
		}
	}

	//初始化dp[0][j]
	for j := 1; j < colLen; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		} else { //如果dp[0][j-1]无法到达,dp[0][j]也无法到达
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < colLen; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[len(obstacleGrid)-1][colLen-1]
}

/*
解法2 动态规划 优化空间复杂度
由状态转移方程 dp[i][j] = dp[i-1][j] + dp[i][j-1] 分析:
1. 首先用滚动数组优化肯定可行,但还不是最优
2. dp[i-1][j] 表示上一行同一列位置, dp[i][j-1] 表示同一行前一列位置
因此可以优化dp为长度为colLen的1维数组,并且必须 i 要从小到达才能保证dp[i-1][j]正确,每次求dp[i][j]时
用的是dp[i-1]的结果. j要从小到达遍历才能保证dp[i][j-1]正确,每次用的是dp[i]的结果
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 {
		return 0
	}
	colLen := len(obstacleGrid[0])

	//dp[i][j] 表示到达obstacleGrid[i][j]的路径条数
	dp := make([]int, colLen)
	if obstacleGrid[0][0] == 0 {
		dp[0] = 1
	}

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < colLen; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else if j > 0 {
				dp[j] = dp[j] + dp[j-1]
			}
			/*
				当j==0时,根据未优化的dp来看dp[i][0]=dp[i-1][0]
				但是dp经过状态压缩后,dp保存的就是上一行的状态,所以
				就会是 dp = dp, 不需要执行
			*/
		}
	}

	return dp[colLen-1]
}

// @lc code=end
