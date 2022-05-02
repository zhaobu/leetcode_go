package main

/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start

/*
解法1 每一行一次二分查找
*/
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		if target < matrix[i][0] || target > matrix[i][n-1] {
			continue
		}
		start, end := 0, n-1
		for start <= end {
			mid := start + (end-start)>>1
			if matrix[i][mid] == target {
				return true
			} else if matrix[i][mid] < target {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}

/*
解法2 从右上角开始z字形
每次判断都可以少一行,或者一列,最多可以执行 m+n 次
*/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	for i, j := 0, n-1; i < m && j >= 0; { //从右上角开始
		if target > matrix[i][j] {
			/*
				1. 如果第一次遍历时进到这个分支,
				说明target>第一行最后一个,说明第一行所有的都小于target,所以i++
				2. 如果第一次进到的是第二个分支,说明最后一列也都可以全部排除,
				所以以后如果再进到这个分支,j可以继续沿用,不需要用最后一列 matrix[i][n-1] 进行判断
			*/
			i++
		} else if target < matrix[i][j] {
			/*
				1. 如果第一次遍历时进到这个分支,
				说明target<第一列最后一个,说明第一列所有的都大于target,所以j--
				2. 如果第一次进到的是第一个分支,说明第一行也都可以全部排除,
				所以以后如果再进到这个分支,i可以继续沿用,不需要用第一行 matrix[0][j] 进行判断
			*/
			j--
		} else {
			return true
		}
	}

	return false
}

// @lc code=end
