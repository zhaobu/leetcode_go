package main

/*
 * @lc app=leetcode.cn id=1905 lang=golang
 *
 * [1905] 统计子岛屿
 */

// @lc code=start
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m, n := len(grid2), len(grid2[0])
	if m < 1 || n < 1 {
		return 0
	}

	ret := 0

	var dfs func(i, j int, isSub *bool)
	dfs = func(i, j int, isSub *bool) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid2[i][j] == 0 {
			return
		}
		if grid1[i][j] != 1 {
			*isSub = false
		}
		grid2[i][j] = 0
		dfs(i-1, j, isSub)
		dfs(i+1, j, isSub)
		dfs(i, j-1, isSub)
		dfs(i, j+1, isSub)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				isSub := true
				dfs(i, j, &isSub)
				if isSub {
					ret++
				}
			}
		}
	}

	return ret
}

// @lc code=end
