package main

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start

/*
解法1 2次二分查找
1. 先找到最后一个matrix[i][0]<=target的那一行
2. 在第一步找到的哪一行再进行一次二分查找
*/
func searchMatrix1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m == 1 && n == 1 {
		return matrix[0][0] == target
	}

	search := func(a []int) bool {
		i, j := 0, n-1
		for i <= j {
			mid := i + (j-i)>>1
			if a[mid] == target {
				return true
			} else if a[mid] < target {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
		return false
	}

	//先找出matrix[i][0]<=target的最后一行
	i, j := 0, m-1
	for i <= j {
		mid := i + (j-i)>>1
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			if mid == m-1 || matrix[mid+1][0] > target {
				return search(matrix[mid])
			} else {
				i = mid + 1
			}
		} else {
			j = mid - 1
		}
	}

	return false
}

/*
解法2 一次二分查找
因为每一行开头元素都大于上一行最后一个元素,所以二维数组可以看成是一个升序数组
*/
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m == 1 && n == 1 {
		return matrix[0][0] == target
	}

	//一维数组的下标转换为二维数组的下标
	getIndex := func(k int) (a, b int) {
		return k / n, k % n
	}

	i, j := 0, m*n-1
	for i <= j {
		mid := i + (j-i)>>1
		a, b := getIndex(mid)
		if matrix[a][b] == target {
			return true
		} else if matrix[a][b] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return false
}

// @lc code=end
