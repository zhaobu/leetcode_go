package main

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start

/*
解法1
*/
func totalNQueens(n int) int {
	if n < 1 {
		return 0
	}

	ret := 0

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
	var backTrack func(row int)
	backTrack = func(row int) {
		//第n行放完就代表找到一个可行解
		if row == n {
			ret++
			return
		}

		//每一行都有n列可选
		for col := 0; col < n; col++ {
			//如果当前放置方式不合适就继续下一个位置
			if !isValid(row, col) {
				continue
			}

			colCheck[col] = 1
			leftTopCheck[row-col+n] = 1
			rightTopCheck[row+col] = 1
			backTrack(row + 1)
			colCheck[col] = 0
			leftTopCheck[row-col+n] = 0
			rightTopCheck[row+col] = 0
		}
	}

	backTrack(0)

	return ret
}

// @lc code=end
