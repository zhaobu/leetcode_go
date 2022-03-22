package main

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	if n < 1 {
		return nil
	}

	ret := [][]string{}

	//初始化棋盘
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}

	//判断当前位置是否可以放置
	isValid := func(row, col int) bool {
		//判断同一列
		for i := 0; i < n; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}
		//判断左上部分
		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		//判断右上部分
		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	//放置第i行
	var backTrack func(record []string, row int)
	backTrack = func(record []string, row int) {
		if row == n {
			ret = append(ret, append([]string{}, record...))
			// fmt.Printf("ret=%+v\n", ret)
			return
		}

		//每一行都有n列可选
		for col := 0; col < n; col++ {
			//如果当前放置方式不合适就继续下一个位置
			if !isValid(row, col) {
				continue
			}

			board[row][col] = 'Q'
			record[row] = string(board[row])
			backTrack(record, row+1)
			board[row][col] = '.'
		}
	}

	backTrack(make([]string, n), 0)

	return ret
}

// @lc code=end
