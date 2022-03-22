package main

/*
 * @lc app=leetcode.cn id=1254 lang=golang
 *
 * [1254] 统计封闭岛屿的数目
 */

// @lc code=start

/*
解法1 dfs
注意淹没边上的陆地时也要用到dfs函数,不能只是grid[i][j]=1
*/
func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 1 || n < 1 {
		return 0
	}

	// 从 (i, j) 开始，将与之相邻的陆地都变成海水
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 1 {
			// 已经是海水了
			return
		}

		// 将 (i, j) 变成海水
		grid[i][j] = 1

		// 淹没上下左右的陆地
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	//将上下靠边的陆地淹掉
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	//将左右靠边的陆地淹掉
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	ret := 0
	/*
	   1. 剩下求得的陆地都是封闭岛屿
	   2. 可以不计算靠边的,i,j都从1开始,到m-1和n-1结束
	*/
	m, n = m-1, n-1
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 0 {
				ret++
				dfs(i, j)
			}
		}
	}

	return ret
}

// @lc code=end
