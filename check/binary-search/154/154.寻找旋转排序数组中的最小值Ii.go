package main

/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */

// @lc code=start

/*
解法1 二分查找
参考153题
*/
func findMin(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			/*
				1. mid一定在属于左半部分,最小值在[mid+1:right]范围内
				2. [left,mid]范围内的值都可以排除掉
			*/
			left = mid + 1
		} else if nums[mid] < nums[right] {
			/*
				1. mid一定属于右半部分,最小指在[left,mid]范围内
				2. [mid+1:right]范围内的元素都可以排除掉
			*/
			right = mid
		} else {
			/*
				唯一能确定的是可以把右边界点right排除掉,让范围缩小到[left,right-1]
			*/
			right -= 1
		}
	}

	return nums[left]
}

// @lc code=end
