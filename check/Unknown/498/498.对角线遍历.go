package main

/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start

/*
解法1 直接遍历
1. 不要舍不得变量,以为变量少代码更简洁,反而不好控制循环
2. 此类问题一般都可以类似图的深度遍历一样设置一个方向数组和当前方向变量
用类似状态机的思想模拟行走方向
*/
func findDiagonalOrder(mat [][]int) []int {
	/*
		1. i,j分别表示行和列,行向下变大,列向右变大
		2. dirs[0]表示向右上的方向,dirs[1]表示向左下的方向
	*/
	dirs := [][]int{{-1, 1}, {1, -1}}
	m, n := len(mat), len(mat[0])

	i, j := 0, 0
	curDir := 0 //初始时,方向是向右上

	ret := make([]int, 0, m*n)
	for i < m && j < n { //也可以用len(ret) <= m*n来控制
		dir := dirs[curDir]
		for {
			ret = append(ret, mat[i][j])
			newI, newJ := i+dir[0], j+dir[1]
			//如果沿当前方向的下一个点不在范围内就结束循环
			if newI < 0 || newI >= m || newJ < 0 || newJ >= n {
				break
			}
			i, j = newI, newJ
		}
		if curDir == 0 {
			//向右上方向时,如果已经是最后一列,则反向后的起始点应该是行+1
			if j == n-1 {
				i++
			} else {
				j++
			}
			curDir = 1
		} else {
			//向左下方向时,如果已经是最后一行,则反向后的起始点应该是列+1
			if i == m-1 {
				j++
			} else {
				i++
			}
			curDir = 0
		}
	}

	return ret
}

// @lc code=end
