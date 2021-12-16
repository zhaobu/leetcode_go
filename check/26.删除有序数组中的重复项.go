/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
// func removeDuplicates(nums []int) int {
// 	if len(nums) < 2 {
// 		return len(nums)
// 	}
// 	lastIndex := 1     //上一个已经整理好的数组的后一个位置
// 	lastNum := nums[0] //上一个数字

// 	for i := 1; i < len(nums); i++ {
// 		if lastNum == nums[i] {
// 			continue
// 		}
// 		lastNum = nums[i]
// 		nums[lastIndex], nums[i] = nums[i], nums[lastIndex]
// 		lastIndex++
// 	}
// 	return lastIndex
// }

//双指针
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 1 //表示已经整理好的部分的下一个索引位置,
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// @lc code=end

