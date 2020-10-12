package Solution

/*
[小浩算法 ：动态规划)](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/204.html#_03%E3%80%81go%E8%AF%AD%E8%A8%80%E7%A4%BA%E4%BE%8B)
*/
func MinPathSum1(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}
	m := len(grid)
	n := len(grid[m-1])
	return grid[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
