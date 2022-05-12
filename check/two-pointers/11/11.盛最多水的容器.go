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
/*
解法1 双指针
*/
func maxArea1(height []int) int {
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

/*
解法2 双指针优化
*/
func maxArea2(height []int) int {
	m := len(height)
	ret := 0

	for i, j := 0, m-1; i < j; {
		cur := 0
		if height[i] <= height[j] {
			cur = (j - i) * height[i]
			k := i + 1
			/*
				1. 因为height[i] <= height[j],所以对于右边界j来说j往左走时,不存在j2使由i,j2构成的面积比由i,j构成的面积
				更大的情况,也就是说就算后面j再怎么变,也用不上当前的i了,所以应该让i往右走
				2. i往右走时应该寻找比height[i]更大的i,但同时i应该小于j
			*/
			for k < j && height[k] <= height[i] {
				k++
			}
			i = k
		} else if height[i] > height[j] {
			cur = (j - i) * height[j]
			k := j - 1
			/*
				和height[i] <= height[j]时同理
			*/
			for k > i && height[k] <= height[j] {
				k--
			}
			j = k
		}
		if cur > ret {
			ret = cur
		}
	}

	return ret
}

/*
解法3 单调栈
1. i从左往右遍历每一个高度,用一个单调递增栈保存遍历过的高度
2. 对于每一个i,求出由每一个栈元素j和i组成的面积,并更新最大值
3. 单调栈其实过滤的是那些在栈顶元素j和i之间的那些情况,因为那些情况height比
栈顶的height小,同时下标也比栈顶的大,自然面积不可能更大
*/
func maxArea(height []int) int {
	m := len(height)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	stack := []int{0}
	ret := 0
	for i := 1; i < m; i++ {
		if i+1 == m || height[i+1] < height[i] {
			for j := len(stack) - 1; j >= 0; j-- {
				cur := (i - stack[j]) * min(height[i], height[stack[j]])
				if cur > ret {
					ret = cur
				}
			}
		}
		if height[i] > height[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}

	return ret
}

// @lc code=end
