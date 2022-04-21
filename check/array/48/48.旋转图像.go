package main

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	m := len(matrix)
	if m == 1 {
		return
	}
	// for i := range matrix {
	// 	fmt.Printf("1matrix[%d]=%+v\n", i, matrix[i])
	// }
	//先沿着左上右下对角线对称翻转
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// for i := range matrix {
	// 	fmt.Printf("2matrix[%d]=%+v\n", i, matrix[i])
	// }
	//然后再反转每一行
	for i := 0; i < m; i++ {
		for j, k := 0, m-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	return
}

// @lc code=end
