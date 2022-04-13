package main

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start

/*
解法1 两遍扫描
官方题解1 (c++的next_permutation库函数实现方式)
*/
func nextPermutation(nums []int) {
	m := len(nums)
	if m < 2 {
		return
	}
	//从后往前找到第一个出现升序序列时的下标
	left := m - 2 //第一个升序序列左边元素下标
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}

	if left != -1 {
		right := m - 1 //left右边第一个比nums[left]大的元素下标
		for right >= 0 && nums[right] <= nums[left] {
			right--
		}

		//交换nums[left],nums[right]
		nums[left], nums[right] = nums[right], nums[left]
	}

	//区间(left,m)此时必定是降序,所以通过双指针让区间[left+1,m)变为升序
	for i, j := left+1, m-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	return
}

// @lc code=end
