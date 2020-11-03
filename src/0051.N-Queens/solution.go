package Solution

// [labuladong 回溯法](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/hui-su-suan-fa-xiang-jie-xiu-ding-ban)
/* 输入棋盘边长 n，返回所有合法的放置 */
func solveNQueens1(n int) [][]string {
	res := [][]string{}
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘。
	board := make([][]rune, n)
	v := make([]rune, n)
	for i := 0; i < n; i++ {
		v[i] = '.'
	}
	for i := 0; i < n; i++ {
		board[i] = append([]rune{}, v...)
	}
	backtrack1(board, 0, &res)
	return res
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func backtrack1(board [][]rune, row int, res *[][]string) {
	// 触发结束条件
	if row == len(board) {
		str := make([]string, 0, len(board[0]))
		for _, v := range board {
			str = append(str, string(v))
		}
		*res = append(*res, str)
		return
	}

	/*
		看成是从第0行,0列开始的从上往下,row递增,从左往右,col递增
		或者是从下往上,row递增,从左往右col递增的n*n的棋盘
	*/
	for col := 0; col < len(board[row]); col++ {
		if !isValid1(board, row, col) {
			continue
		}
		board[row][col] = rune('Q')
		backtrack1(board, row+1, res)
		board[row][col] = rune('.')
	}
}

/* 是否可以在 board[row][col] 放置皇后？ */
func isValid1(board [][]rune, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == rune('Q') {
			return false
		}
	}
	// 检查右下方对角线
	for i, j := row-1, col+1; i >= 0 && j < n; i-- {
		if board[i][j] == rune('Q') {
			return false
		}
		j++
	}
	// 检查左下方对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i-- {
		if board[i][j] == rune('Q') {
			return false
		}
		j--
	}
	return true
}

/*
[TODO leetcode基于位运算的回溯](https://leetcode-cn.com/problems/n-queens/solution/nhuang-hou-by-leetcode-solution/)
*/
func solveNQueens2(n int) [][]string {
	return nil
}
