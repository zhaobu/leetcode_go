package check

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start

// @lc code=start
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}
	i := 0
	j := len(height) - 1
	max := 0
	for i < j {
		curArea := 0 //i和j之间的最大容器的值
		if height[i] < height[j] {
			curArea = height[i] * (j - i)
			i++ //如果左边垂直线较小,就往右移
		} else {
			curArea = height[j] * (j - i)
			j-- //如果右边垂直线较小,就往左移
		}
		if curArea > max {
			max = curArea
		}
	}
	return max
}

// @lc code=end
