package main

import "fmt"

/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start

/*
解法1 前缀和
*/
type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	preSum := make([][]int, m+1)

	//初始化第二维
	preSum[0] = make([]int, n+1)
	for k1 := range matrix {
		preSum[k1+1] = make([]int, n+1)
	}

	//计算前缀和
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	// printPreSum(preSum)
	return NumMatrix{
		preSum: preSum,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	preSum := this.preSum
	row2, col2 = row2+1, col2+1
	return preSum[row2][col2] - preSum[row1][col2] - preSum[row2][col1] + preSum[row1][col1]
}

func printPreSum(preSum [][]int) {
	for i := range preSum {
		fmt.Printf("preSum[%d]=%+v\n", i, preSum[i])
	}
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end
