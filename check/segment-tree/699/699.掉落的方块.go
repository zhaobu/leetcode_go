package main

/*
 * @lc app=leetcode.cn id=699 lang=golang
 *
 * [699] 掉落的方块
 */

// @lc code=start

/*
解法1 暴力
*/
func fallingSquares(positions [][]int) []int {
	//判断positions[i]和positions[j]底边是否有重合部分
	check := func(i, j int) bool {
		p1, p2 := positions[i], positions[j]
		// fmt.Printf("p[%d]=%+v,p[%d]=%+v\n",i , p1, j, p2)
		if p2[0] > p1[0]+p1[1]-1 { //p2正方形的左边>p1正方形的右边
			return false
		}
		if p2[0]+p2[1]-1 < p1[0] { //p2正方形的右边<p1正方形的左边
			return false
		}
		return true
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	height := make([]int, len(positions)) //height[i]表示遍历完第i个正方形后其堆叠高度
	for i := 0; i < len(positions); i++ {
		height[i] = positions[i][1]
		for j := 0; j < i; j++ {
			if check(i, j) {
				// fmt.Printf("i=%d,j=%d,%d有重合\n",i , j, positions[j][1]+positions[i][1])
				height[i] = max(height[i], height[j]+positions[i][1])
			}
		}
	}

	//求出遍历完第i个正方形后,当前所有遍历过的正方形所能到达的最高高度
	for i := 1; i < len(height); i++ {
		height[i] = max(height[i], height[i-1])
	}
	return height
}

// @lc code=end
