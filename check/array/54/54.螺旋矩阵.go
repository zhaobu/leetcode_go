package main

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start

/*
解法1: 模拟遍历方向
注意切换方向时,下标要往当前方向回退一格
*/

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	const (
		Left  = 1
		Right = 2
		Up    = 3
		Down  = 4
	)
	ret := make([]int, 0, m*n)
	const Fill = 101 //用来标记已访问
	direct := Right  //当前前进的方向
	count := m * n   //总共会遍历的总次数
	i, j := 0, 0     //下标索引
	for count > 0 {
		switch direct {
		case Right:
			for ; j < n && matrix[i][j] != Fill; j++ {
				ret = append(ret, matrix[i][j])
				matrix[i][j] = Fill
				count--
			}
			i++
			j-- //注意要回退一格,下同
			direct = Down
		case Down:
			for ; i < m && matrix[i][j] != Fill; i++ {
				ret = append(ret, matrix[i][j])
				matrix[i][j] = Fill
				count--
			}
			j--
			i--
			direct = Left
		case Left:
			for ; j >= 0 && matrix[i][j] != Fill; j-- {
				ret = append(ret, matrix[i][j])
				matrix[i][j] = Fill
				count--
			}
			i--
			j++
			direct = Up
		case Up:
			for ; i >= 0 && matrix[i][j] != Fill; i-- {
				ret = append(ret, matrix[i][j])
				matrix[i][j] = Fill
				count--
			}
			j++
			i++
			direct = Right
		}
	}
	// fmt.Printf("count=%d\n", count)
	return ret
}

// @lc code=end
