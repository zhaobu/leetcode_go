package Solution

/*
[小浩算法 ：动态规划)](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/205.html#_02%E3%80%81%E9%A2%98%E7%9B%AE%E5%9B%BE%E8%A7%A3)
*/
func MinPathSum1(grid [][]int) int {
	m := len(grid)
	n := len(grid[m-1])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i != 0 && j != 0 {
				grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
			} else if j == 0 && i != 0 {
				grid[i][j] += grid[i-1][j]
			} else if i == 0 && j != 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}
	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {

		return a
	}
	return b
}
