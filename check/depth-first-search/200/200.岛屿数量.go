package main

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start

/*
解法1 dfs
1. 把二维数组看做是有4条边的图进行dfs遍历
2. 把遍历过的所有地方都变为0可以避免重复达到,可以少用visit数组
*/
func numIslands1(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	if m < 1 || n < 1 {
		return 0
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ret++
			}
		}
	}

	return ret
}

/*
解法2 bfs
1. 使用队列进行bfs遍历
2. 把遍历过的所有地方都变为0可以避免重复达到,可以少用visit数组
*/
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	if m < 1 || n < 1 {
		return 0
	}

	type land struct {
		row int
		col int
	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ret++

				grid[i][j] = '0'
				queue := []*land{{i, j}}

				for len(queue) > 0 {
					front := queue[0]
					queue = queue[1:]
					row, col := front.row, front.col
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, &land{row: row - 1, col: col})
						grid[row-1][col] = '0'
					}
					if row+1 < m && grid[row+1][col] == '1' {
						queue = append(queue, &land{row: row + 1, col: col})
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, &land{row: row, col: col - 1})
						grid[row][col-1] = '0'
					}
					if col+1 < n && grid[row][col+1] == '1' {
						queue = append(queue, &land{row: row, col: col + 1})
						grid[row][col+1] = '0'
					}
				}

			}
		}
	}

	return ret
}

// @lc code=end
