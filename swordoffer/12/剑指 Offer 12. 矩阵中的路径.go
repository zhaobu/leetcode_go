package main

/*
 *
 *
 * 剑指 Offer 12. 矩阵中的路径
 */

// @lc code=start

/*
 解法1 回溯
*/
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	if m*n < len(word) || len(word) == 0 {
		return false
	}
	//dfs回溯前进行字符统计
	cnt := [58]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ch := board[i][j]
			cnt[ch-'A']++
		}
	}
	for _, ch := range word {
		if cnt[ch-'A'] < 1 {
			return false
		}
		cnt[ch-'A']--
	}

	// fmt.Printf("%d", 'z'-'A')
	ret := false
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

	var dfs func(i, j, index int, visited [][]bool)
	dfs = func(i, j, index int, visited [][]bool) {
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || ret {
			return
		}
		if board[i][j] != word[index] {
			return
		}
		visited[i][j] = true
		if index == len(word)-1 {
			ret = true
			return
		}
		for _, dir := range dirs {
			i1, j1 := i+dir[0], j+dir[1]
			dfs(i1, j1, index+1, visited)
		}
		visited[i][j] = false
	}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] {
				dfs(i, j, 0, visited)
			}
		}
	}

	return ret
}

// @lc code=end
