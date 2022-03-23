package main

import (
	"math/bits"
)

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start

/*
解法1 backtrack 常规解法
注意判断是否可以放置皇后的isValid函数
*/
func solveNQueens1(n int) [][]string {
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

/*
解法2 backtrack 用hashmap优化isValid
*/
func solveNQueens2(n int) [][]string {
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

	colCheck := make([]int, n)        //记录列上皇后放置情况,索引为col
	leftTopCheck := make([]int, 2*n)  //记录左上皇后放置情况,索引为row-col+n
	rightTopCheck := make([]int, 2*n) //记录右上皇后放置情况,索引为row+col
	//判断当前位置是否可以放置
	isValid := func(row, col int) bool {
		//判断同一列
		if colCheck[col] == 1 {
			return false
		}
		//判断左上部分
		if leftTopCheck[row-col+n] == 1 {
			return false
		}
		//判断右上部分
		if rightTopCheck[row+col] == 1 {
			return false
		}
		return true
	}

	//放置第i行
	var backTrack func(record []string, row int)
	backTrack = func(record []string, row int) {
		//第n行放完就代表找到一个可行解
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
			colCheck[col] = 1
			leftTopCheck[row-col+n] = 1
			rightTopCheck[row+col] = 1
			backTrack(record, row+1)
			board[row][col] = '.'
			colCheck[col] = 0
			leftTopCheck[row-col+n] = 0
			rightTopCheck[row+col] = 0
		}
	}

	backTrack(make([]string, n), 0)

	return ret
}

/*
解法3 backtrack 位操作优化
*/
func solveNQueens(n int) [][]string {
	if n < 1 {
		return nil
	}

	ret := [][]string{}

	generateBoard := func(record []int, n int) []string {
		board := []string{}
		for i := 0; i < n; i++ {
			row := make([]byte, n)
			for j := 0; j < n; j++ {
				row[j] = '.'
			}
			row[record[i]] = 'Q'
			board = append(board, string(row))
		}
		return board
	}

	//放置第i行
	var backTrack func(record []int, row, colCheck, leftTopCheck, rightTopCheck int)
	backTrack = func(record []int, row, colCheck, leftTopCheck, rightTopCheck int) {
		//第n行放完就代表找到一个可行解
		if row == n {
			ret = append(ret, generateBoard(record, n))
			// fmt.Printf("ret=%+v\n", ret)
			return
		}
		//可以放置房后的位置
		place := ((1 << n) - 1) & (^(colCheck | leftTopCheck | rightTopCheck))
		//每一行都有n列可选
		for ; place != 0; place &= (place - 1) {
			col := place & (-place)
			colCount := bits.OnesCount(uint(col - 1))
			record[row] = colCount
			backTrack(record, row+1, colCheck|col, (leftTopCheck|col)>>1, (rightTopCheck|col)<<1)
			record[row] = -1
		}
	}

	//初始化棋盘
	board := make([]int, n)
	for i := 0; i < n; i++ {
		board[i] = -1
	}
	backTrack(board, 0, 0, 0, 0)

	return ret
}

// @lc code=end
