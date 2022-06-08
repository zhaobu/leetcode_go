package main

/*
 * @lc app=leetcode.cn id=1037 lang=golang
 *
 * [1037] 有效的回旋镖
 */

// @lc code=start
func isBoomerang(points [][]int) bool {
	/*
	   y1-y2     y3-y2
	   x1-x2     x3-x2
	*/
	p1, p2, p3 := points[0], points[1], points[2]
	return (p1[1]-p2[1])*(p3[0]-p2[0]) != (p1[0]-p2[0])*(p3[1]-p2[1])
}

// @lc code=end
