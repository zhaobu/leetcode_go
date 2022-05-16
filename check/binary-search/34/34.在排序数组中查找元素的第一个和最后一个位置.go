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
	if n == 0 {
		return []int{-1, -1}
	}

	ret := []int{-1, -1}
	//二分查找第一个为target的下标
	i, j := 0, n-1
	for i <= j {
		mid := i + (j-i)>>1
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			if mid > 0 && nums[mid-1] == target { //如果mid左边还有等于target的继续在左区间二分
				if mid > ret[1] { //尽量让二分查找右边界时的左起始位置靠右,减小查找区间
					ret[1] = mid
				}
				j = mid - 1
			} else {
				ret[0] = mid
				break
			}
		}
	}
	if ret[0] == -1 {
		return ret
	}

	//二分查找最后一个为target的下标
	i, j = ret[0], n-1
	if ret[1] > ret[0] { //第一次二分查找时,可能不会更新ret[1]
		i = ret[1]
	}
	// fmt.Printf("i=%d,j=%d\n", i, j)
	for i <= j {
		mid := i + (j-i)>>1
		if nums[mid] < target {
			i = mid + 1
		} else if nums[mid] > target {
			j = mid - 1
		} else {
			if mid < n-1 && nums[mid+1] == target { //如果mid右边还有等于target的继续在右区间二分
				i = mid + 1
			} else {
				ret[1] = mid
				break
			}
		}
	}

	return ret
}

// @lc code=end
