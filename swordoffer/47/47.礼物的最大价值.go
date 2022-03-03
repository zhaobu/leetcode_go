package main

import "fmt"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * 剑指 Offer 47. 礼物的最大价值
 */

/*
 解法2
 动态规划,优化空间复杂度
由于 dp[i][j] 只与 dp[i-1][j] , dp[i][j-1] ,
 grid[i][j]有关系，因此可以将原矩阵 grid用作 dp 矩阵，即直接在 grid 上修改即可

*/
// @lc code=start
func maxValue(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	/*
	   从坐上角开始,i表示行,j表示列
	   dp[i][j]表示走到第i行,第j列时能拿到的最大价值
	*/
	rowSum := 0  //用于累加第0列
	cellSum := 0 //用于累加第0行

	//初始化第0列
	for i := 0; i < len(grid); i++ {
		//初始化每一行为空的数组
		rowSum += grid[i][0]
		grid[i][0] = rowSum
	}
	//初始化第0行
	for j := 0; j < len(grid[0]); j++ {
		cellSum += grid[0][j]
		grid[0][j] = cellSum
	}

	fmt.Printf("**1**grid=%+v\n", grid)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			grid[i][j] = max(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	fmt.Printf("**2**grid=%+v \n", grid)

	return grid[len(grid)-1][len(grid[0])-1]
}

/*
 解法1
 动态规划
*/
// @lc code=start
func maxValue1(grid [][]int) int {
	if len(grid) < 1 || len(grid[0]) < 1 {
		return 0
	}
	/*
	   从坐上角开始,i表示行,j表示列
	   dp[i][j]表示走到第i行,第j列时能拿到的最大价值
	*/
	dp := make([][]int, len(grid))
	rowSum := 0  //用于累加第0列
	cellSum := 0 //用于累加第0行

	//初始化第0列
	for i := 0; i < len(grid); i++ {
		//初始化每一行为空的数组
		dp[i] = make([]int, len(grid[i]))
		rowSum += grid[i][0]
		dp[i][0] = rowSum
	}
	//初始化第0行
	for j := 0; j < len(grid[0]); j++ {
		cellSum += grid[0][j]
		dp[0][j] = cellSum
	}

	// fmt.Printf("**1**grid=%+v , dp=%+v\n", grid, dp)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	// fmt.Printf("**2**grid=%+v , dp=%+v\n", grid, dp)

	return dp[len(grid)-1][len(grid[0])-1]
}

// @lc code=end
