/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
// 解法1:使用map记录另一个数
func twoSum(nums []int, target int) []int {
	mapIndex := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if a, ok := mapIndex[target-nums[i]]; ok {
			return []int{i, a}
		}
		mapIndex[nums[i]] = i
	}
	return nil
}

// @lc code=end

