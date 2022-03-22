package main

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start

/*
解法1 dfs
*/
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 1 || n < 1 {
		return 0
	}

	ret := 0

	var dfs func(i, j int, count *int)
	dfs = func(i, j int, count *int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 {
			return
		}
		*count += 1
		grid[i][j] = 0
		dfs(i-1, j, count)
		dfs(i+1, j, count)
		dfs(i, j-1, count)
		dfs(i, j+1, count)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			dfs(i, j, &count)
			if count > ret {
				ret = count
			}
		}
	}

	return ret
}

// @lc code=end
