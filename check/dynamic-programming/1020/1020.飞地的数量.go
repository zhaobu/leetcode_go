package main

/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */

// @lc code=start

/*
解法1
*/
func numEnclaves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 1 || n < 1 {
		return 0
	}

	ret := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		ret++
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	//上下边界陆地变成海洋
	for j := 0; j < n; j++ {
		dfs(0, j)
		// printGrid(grid)
		dfs(m-1, j)
		// printGrid(grid)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	ret = 0
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
		}
	}

	return ret
}

// @lc code=end
