package main

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
/*
解法1 2次二分查找
*/
func searchRange(nums []int, target int) []int {
	n := len(nums)
	ret := []int{-1, -1}
	if n == 0 {
		return ret
	}
	//找到一个target的下标
	left, right, mid := 0, n-1, 0
	for left <= right {
		mid = left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			break
		}
	}
	if nums[mid] != target {
		return []int{-1, -1}
	}
	//   fmt.Printf("left=%d, right=%d\n", left, right)
	//找到第一个target的下标
	for i, j := left, right; i <= j; {
		mid = i + (j-i)>>1
		if nums[mid] < target {
			i = mid + 1
		} else {
			if mid == left || nums[mid-1] != target {
				ret[0] = mid
				break
			}
			j = mid - 1
		}
	}
	//   fmt.Printf("left=%d, right=%d\n", left, right)

	//找到最后一个target的下标
	for i, j := left, right; i <= j; {
		mid = i + (j-i)>>1
		if nums[mid] > target {
			j = mid - 1
		} else {
			if mid == right || nums[mid+1] != target {
				ret[1] = mid
				break
			}
			i = mid + 1
		}
	}

	return ret
}

// @lc code=end
