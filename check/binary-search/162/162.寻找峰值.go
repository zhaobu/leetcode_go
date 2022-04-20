package main

/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start

/*
解法1 直接求得第一个波峰
求第一个nums[i] > nums[i+1]的值,所以在区间[0,i]上必然是递增的
*/
func findPeakElement1(nums []int) int {
	n := len(nums)
	if n < 2 { //n必定>=1
		return 0
	}

	for i := range nums {
		if i+1 < n && nums[i] > nums[i+1] {
			return i
		}
	}
	return n - 1
}

/*
解法2 二分查找
*/
func findPeakElement(nums []int) int {
	n := len(nums)
	if n < 2 { //n必定>=1
		return 0
	}

	//判断nums[idx1] < nums[idx2]
	less := func(idx1, idx2 int) bool {
		if idx1 == -1 || idx1 == n {
			return true
		}
		if idx2 == n || idx2 == -1 {
			return false
		}
		return nums[idx1] < nums[idx2]
	}

	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if less(mid-1, mid) && less(mid+1, mid) {
			return mid
		}
		if less(mid, mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// @lc code=end
