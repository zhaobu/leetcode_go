package main

/*
 * @lc app=leetcode.cn id=81 lang=golang
 *
 * [81] 搜索旋转排序数组 II
 */

// @lc code=start

/*
解法1 二分查找
因为数组中存在重复元素,所以可能重复元素分布在数组的两端
比如数组 [2,3,4,5,0,1,2,2] 这种,要先去除2端元素,
保证数组具有二段性,变成[3,4,5,0,1,2]之后就和
33. 搜索旋转排序数组是一样的了

*/
func search(nums []int, target int) bool {
	n := len(nums)
	if n == 1 {
		return nums[0] == target
	}

	//二分查找
	left, right := 0, n-1
	for left < right && nums[left] == nums[right] {
		right--
	}
	// 第1步: 先保证数组具有二段性
	for left <= right && nums[left] == nums[right] {
		if nums[left] == target {
			return true
		}
		left++
		right--
	}
	// 第2步: 二分查找带有重复元素的旋转数组
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return true
		}
		if nums[mid] >= nums[left] {
			/*
				1. [left, mid]上面升序
				2. nums[mid] != target,应该考虑在范围[low,mid-1]范围查找
			*/
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			/*
				1. 在[mid,right]上升序
				2. nums[mid] != target,应该考虑在范围[mid+1,right]范围继续查找
			*/
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

// @lc code=end
