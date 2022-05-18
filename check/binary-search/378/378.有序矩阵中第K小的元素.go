package main

/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第 K 小的元素
 */

// @lc code=start

/*
解法1 二分查找
此题思想类似 https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/
*/
func kthSmallest(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])

	//找出第row行最后一个<=target的下标
	search := func(row, target int) int {
		i, j := 0, n-1
		for i <= j {
			mid := i + (j-i)>>1
			if matrix[row][mid] > target {
				j = mid - 1
			} else {
				if mid == n-1 || matrix[row][mid+1] > target {
					return mid
				}
				i = mid + 1
			}
		}

		return -1
	}

	//统计<=m的个数
	getCount := func(num int) (count int) {
		//找出第一个
		for i := 0; i < m; i++ {
			index := search(i, num)
			if index == -1 {
				break
			}
			count += index + 1
		}
		return count
	}

	//二分查找第一个数mid,在矩阵中<=mid的数的个数>=k
	i, j := matrix[0][0], matrix[m-1][n-1]
	for i <= j {
		mid := i + (j-i)>>1
		count := getCount(mid)
		if count < k {
			i = mid + 1
		} else {
			if mid == 1 || getCount(mid-1) < k {
				return mid
			}
			j = mid - 1
		}
	}

	return -1
}

// @lc code=end
