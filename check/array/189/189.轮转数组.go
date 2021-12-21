package main

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */

// @lc code=start
//解法1:暴力解法,超时
func rotate1(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	k = k % n

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
func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newIndex := (i + k) % len(nums)
		newNums[newIndex] = nums[i]
	}
	copy(nums, newNums)
}

// 解法3 反转
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	k %= len(nums)
	reverse(nums[:])
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums) >> 1
	for i, j := 0, len(nums)-1; i < n; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}

// @lc code=end
