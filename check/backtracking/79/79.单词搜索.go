package main

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start

/*
解法1 回溯法
*/
func exist1(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	strLen := len(word)
	if strLen > m*n {
		return false
	}

	ret := false
	visited := make([][]bool, m) //记录已访问位置
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	/*
		1. k表示当前正在匹配的word字符的下标
		2. board[i][j]表示board中用来进行匹配的字符
	*/
	strLen -= 1 //word最后一个字符的下标
	var dfs func(k, i, j int)
	dfs = func(k, i, j int) {
		if ret {
			return
		}
		if (i < 0 || i >= m) || (j < 0 || j >= n) {
			return
		}
		if board[i][j] != word[k] {
			return
		}
		if visited[i][j] {
			return
		}
		if k == strLen {
			ret = true
			return
		}
		visited[i][j] = true
		//往上
		dfs(k+1, i-1, j)
		//往下
		dfs(k+1, i+1, j)
		//往左
		dfs(k+1, i, j-1)
		//往右
		dfs(k+1, i, j+1)
		visited[i][j] = false //回溯
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				dfs(0, i, j)
			}
		}
	}

	return ret
}

/*
解法2 回溯法
用一个字母个数表进行优化
*/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	strLen := len(word)
	if strLen > m*n {
		return false
	}

	ret := false
	visited := make([][]bool, m) //记录已访问位置
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	/*
		1. k表示当前正在匹配的word字符的下标
		2. board[i][j]表示board中用来进行匹配的字符
	*/
	strLen -= 1 //word最后一个字符的下标
	var dfs func(k, i, j int)
	dfs = func(k, i, j int) {
		if ret {
			return
		}
		if (i < 0 || i >= m) || (j < 0 || j >= n) {
			return
		}
		if board[i][j] != word[k] {
			return
		}
		if visited[i][j] {
			return
		}
		if k == strLen {
			ret = true
			return
		}
		visited[i][j] = true
		//往上
		dfs(k+1, i-1, j)
		//往下
		dfs(k+1, i+1, j)
		//往左
		dfs(k+1, i, j-1)
		//往右
		dfs(k+1, i, j+1)
		visited[i][j] = false //回溯
	}
	/*
		1. A的ascii码为65,z的ascii码位122,所需需要一个容量为122-65+1的数组
		2. 先统计board中的字符个数,再通过word过滤掉那些个数不足以匹配的情况
	*/
	record := [58]byte{}
	for _, v := range board {
		for _, v1 := range v {
			record[v1-'A']++
		}
	}
	for _, v := range word {
		if record[v-'A'] < 1 {
			return false
		}
		record[v-'A']--
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				dfs(0, i, j)
			}
		}
	}

	return ret
}

// @lc code=end
