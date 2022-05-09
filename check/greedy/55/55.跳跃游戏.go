package main

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
/*
解法1
*/
func canJump1(nums []int) bool {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m := len(nums) - 1
	curMax := 0 //最远能达到的下标
	for i := 0; i <= m; i++ {
		curMax = max(curMax, nums[i]+i)
		// fmt.Printf("curMax=%d, m=%d\n", curMax, m)
		if curMax >= m { //如果最后一个下标已经能到达,就提前返回
			return true
		}
		/*
			1. 运行到这里说明curMax < m
			2. 如果遍历到curMax,说明没有找到比curMax更大的下标
		*/
		if i == curMax {
			return false
		}
	}

	return false
}

/*
解法2
*/
func canJump(nums []int) bool {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m := len(nums)
	curMax := 0
	for i := 0; i < m; i++ {
		if i > curMax {
			return false
		}
		curMax = max(curMax, nums[i]+i)
	}

	return true
}

// @lc code=end
