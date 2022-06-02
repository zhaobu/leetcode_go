package main

import (
	. "leetcode/check"
)

/*
 *
 *
 * 剑指 Offer 04. 二维数组中的查找
 */

// @lc code=start

/*
 解法1 双指针
*/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	for i, j := 0, n-1; i < m && j >= 0; {
		// fmt.Printf("i=%d,j=%d,matrix=%d\n",i, j, matrix[i][j])
		if target > matrix[i][j] {
			i++
		} else if target < matrix[i][j] {
			j--
		} else {
			return true
		}
	}

	return false
}

// @lc code=end
