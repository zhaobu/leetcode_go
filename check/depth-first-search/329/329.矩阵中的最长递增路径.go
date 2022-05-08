package main

/*
 * @lc app=leetcode.cn id=329 lang=golang
 *
 * [329] 矩阵中的最长递增路径
 */

// @lc code=start

/*
解法1 暴力法(超时)
*/
func longestIncreasingPath1(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	ret := 0

	var dfs func(i, j, last, length int, visited [][]int)
	dfs = func(i, j, last, length int, visited [][]int) {
		if (i < 0 || i >= m) || (j < 0 || j >= n) {
			return
		}
		if visited[i][j] == 1 {
			return
		}
		if matrix[i][j] > last {
			length += 1
		} else {
			length = 1
		}
		if length > ret {
			ret = length
		}
		visited[i][j] = 1
		last = matrix[i][j]
		//往上
		dfs(i-1, j, last, length, visited)
		//往下
		dfs(i+1, j, last, length, visited)
		//往左
		dfs(i, j-1, last, length, visited)
		//往右
		dfs(i, j+1, last, length, visited)
		visited[i][j] = 0
	}

	for i := range matrix {
		for j := range matrix[i] {
			visited := make([][]int, m)
			for k := range visited {
				visited[k] = make([]int, n)
			}
			dfs(i, j, -1, 0, visited)
		}
	}

	return ret
}

/*
解法2
*/
var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func longestIncreasingPath2(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	ret := 0
	/*
		record[i][j]记录从matrix[i][j]作为起始点的最长递增路径的长度
	*/
	record := make([][]int, m)
	for k := range record {
		record[k] = make([]int, n)
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	/*
		计算从matrix[i][j]作为起始点的最长递增路径的长度
	*/
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if record[i][j] != 0 {
			return record[i][j]
		}
		/*
			如果matrix[i][j]下一步不能往4个方向中的任一个移动,则
			从matrix[i][j]作为起始点的最长递增路径的长度就为1
		*/
		record[i][j] = 1
		for _, dir := range dirs {
			i1, j1 := i+dir[0], j+dir[1]
			if (i1 >= 0 && i1 < m) && (j1 >= 0 && j1 < n) && matrix[i1][j1] > matrix[i][j] {
				record[i][j] = max(record[i][j], dfs(i1, j1)+1)
			}
		}
		return record[i][j]
	}

	for i := range matrix {
		for j := range matrix[i] {
			//每个点都作为起点计算一次
			ret = max(dfs(i, j), ret)
		}
	}

	return ret
}

/*
解法3 拓扑排序
*/

func longestIncreasingPath(matrix [][]int) int {
	return 0
}

// @lc code=end
