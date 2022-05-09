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
解法2 dfs
*/
func longestIncreasingPath2(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
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
		/*
			当matrix[i][j]往4个方向都计算一遍最长递增路径后,取4个方向上的最大值就是结果
		*/
		return record[i][j]
	}

	for i := range matrix {
		for j := range matrix[i] {
			//每个点都作为起点计算一次
			if record[i][j] == 0 {
				ret = max(dfs(i, j), ret)
			}
		}
	}

	return ret
}

/*
解法3 拓扑排序
*/
func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	outDegs := make([][]int, m) //每个节点的出度
	queue := [][2]int{}
	//初始化每个节点的出度
	for i := range matrix {
		outDegs[i] = make([]int, n)
		for j := range matrix[i] {
			for _, dir := range dirs {
				i1, j1 := i+dir[0], j+dir[1]
				if (i1 >= 0 && i1 < m) && (j1 >= 0 && j1 < n) && matrix[i1][j1] > matrix[i][j] {
					outDegs[i][j]++
				}
			}
			//4个方向上的节点都比当前节点小
			if outDegs[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	ret := 0
	/*
		1. 记录当前层的节点个数
		2. 每出队一个元素,就--,当变为0时就表示当前层遍历完成
	*/
	curLevelNum := len(queue)
	// 广度优先搜索
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		curLevelNum--
		//更新head节点四周的节点的出度
		i, j := head[0], head[1]
		for _, dir := range dirs {
			i1, j1 := i+dir[0], j+dir[1]
			//如果某一个方向上有一个节点比当前节点小,那这个节点一定有一条往当前节点的边
			if (i1 >= 0 && i1 < m) && (j1 >= 0 && j1 < n) && matrix[i][j] > matrix[i1][j1] {
				outDegs[i1][j1]--
				if outDegs[i1][j1] == 0 {
					queue = append(queue, [2]int{i1, j1})
				}
			}
		}
		if curLevelNum == 0 {
			ret++
			curLevelNum = len(queue)
		}
	}
	return ret
}

// @lc code=end
