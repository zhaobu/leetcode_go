package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (62.34%)
 * Likes:    2996
 * Dislikes: 0
 * Total Accepted:    581.2K
 * Total Submissions: 932.3K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 * 解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 * 示例 2：
 *
 *
 * 输入：height = [1,1]
 * 输出：1
 *
 *
 * 示例 3：
 *
 *
 * 输入：height = [4,3,2,1,4]
 * 输出：16
 *
 *
 * 示例 4：
 *
 *
 * 输入：height = [1,2,1]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * n == height.length
 * 2 <= n <= 10^5
 * 0 <= height[i] <= 10^4
 *
 *
 */

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
