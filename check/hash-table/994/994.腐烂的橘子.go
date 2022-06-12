package main

/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] 腐烂的橘子
 */

// @lc code=start

/*
解法1 dfs
*/
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	/*
		1. 当grid[i][j]为1时,timeRecord[i][j]表示的就是该新鲜橘子被腐烂必须经过的最小分钟数
		2. 遍历结束时,timeRecord[i][j]最大值就是所有新鲜橘子被腐烂必须经过的最小分钟数
	*/
	timeRecord := make([][]int, m)
	/*
		1. goodCnt表示新鲜橘子的个数,在遍历过程中,新鲜橘子第一次能被腐烂时就减1
		2. 遍历结束时如果>0,则表示不是所有橘子都能被腐烂
	*/
	goodCnt := 0
	for i := range timeRecord {
		timeRecord[i] = make([]int, n)
		for j := range timeRecord[i] {
			if grid[i][j] == 1 {
				goodCnt++
			}
		}
	}
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	/*
		1. 函数的意义表示从第grid[i][j]这个腐烂橘子开始,往四周继续腐烂其他橘子
		计算过程中不断更新timeRecord
		2. k表示grid[i][j]被腐烂时经过的时间,如果是第一个腐烂橘子, k=0
		3. 从每个第一个腐烂橘子开始,计算一次
	*/
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		for _, dir := range dirs {
			newI, newJ := i+dir[0], j+dir[1]
			if newI >= 0 && newI < m && newJ >= 0 && newJ < n && grid[newI][newJ] == 1 {
				grid[newI][newJ] = 2
				k++

				if timeRecord[newI][newJ] == 0 {
					//该新鲜橘子第一次能被腐烂
					goodCnt--
					timeRecord[newI][newJ] = k
				} else if timeRecord[newI][newJ] > k {
					/*
						1. 如果该新鲜橘子不是第一次能被腐烂,说明之前肯定有其他腐烂源能传染到这个橘子
						2. 记录2种传播途径中的较小值
					*/
					timeRecord[newI][newJ] = k
				}
				dfs(newI, newJ, k)

				//回溯
				k--
				grid[newI][newJ] = 1
			}
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				//每个腐烂橘子源头都计算一次
				dfs(i, j, 0)
			}
		}
	}

	if goodCnt > 0 {
		return -1
	}

	ret := 0
	//计算所有橘子腐烂需要的时间的最大值
	for i := range timeRecord {
		for j := range timeRecord[i] {
			if timeRecord[i][j] > ret {
				ret = timeRecord[i][j]
			}
		}
	}

	return ret
}

// @lc code=end
