package main

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start

/*
解法1 二分查找
*/
func search1(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] {
			/*
				1. [0,mid]区间升序
				2. 已经确定nums[mid] != target,应该把范围缩小到[left,mid-1]
			*/
			if target >= nums[low] && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			/*
				1. [mid, n)区间升序
				2. 已经确定nums[mid] != target,应该把范围缩小到[mid+1,right]
			*/
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}

	return -1
}

/*
解法1 二分查找(官方写法)
*/

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[0] { //[0,mid]是升序的
			if target >= nums[0] && target < nums[mid] { //target在[0,mid)
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { //[mid,len(nums)-1]是升序的
			if target <= nums[len(nums)-1] && target > nums[mid] { //target在(mid,len(nums)-1]
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}

// @lc code=end
