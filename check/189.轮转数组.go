/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
//解法1:暴力解法,超时
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	k = k % n
	for i := 0; i < k; i++ {
		rotateOnce(nums)
	}
	return
}

func rotateOnce(nums []int) {
	var (
		n    = len(nums)
		last = nums[n-1]
	)
	for i := n - 2; i >= 0; i-- {
		nums[i+1] = nums[i]
	}
	nums[0] = last
	return
}

//解法2:复制
// func rotate(nums []int, k int) {
// 	newNums := make([]int, len(nums))
// 	for i := 0; i < len(nums); i++ {
// 		newIndex := (i + k) % len(nums)
// 		newNums[newIndex] = nums[i]
// 	}
// 	copy(nums, newNums)
// }

// @lc code=end