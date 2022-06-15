package main

import "sort"

/*
 * @lc app=leetcode.cn id=719 lang=golang
 *
 * [719] 找出第 k 小的距离对
 */

// @lc code=start

/*
解法1 二分查找
*/
func smallestDistancePair1(nums []int, k int) int {
	//排序
	sort.Ints(nums)
	//二分查找
	i, j := 0, nums[len(nums)-1]-nums[0]

	/*
		1. 求出距离对<=pivot的个数
	*/
	getCnt := func(pivot int) (cnt int) {
		for i := range nums {
			//二分查找最后一个nums[j]-nums[i] <= pivot 的下标j
			j1, j2 := i+1, len(nums)-1
			last := -1
			for j1 <= j2 {
				mid := j1 + (j2-j1)>>1
				if nums[mid]-nums[i] > pivot {
					j2 = mid - 1
				} else {
					if mid == len(nums)-1 || nums[mid+1]-nums[i] > pivot {
						last = mid
						break
					}
					j1 = mid + 1
				}
			}
			if last != -1 {
				cnt += last - i
			}
		}
		return cnt
	}

	for i < j {
		mid := i + (j-i)>>1
		curCnt := getCnt(mid)
		if curCnt < k {
			//如果<=mid的距离对个数小于k,则应该让mid变大,继续在右区间查找
			i = mid + 1
		} else {
			/*
				1. 如果<=mid的距离对个数>k,则有可能是因为有多个距离对都等于mid.所以应该在左区间查找
				2. 如果<=mid的距离对个数等于k,则此时这个mid不一定存在,所以此时不能直接返回mid,只能继续减小二分范围
			*/
			j = mid
		}
	}

	return i
}

/*
解法2 库函数二分查找
*/
func smallestDistancePair2(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt := 0
		for j, num := range nums {
			i := sort.SearchInts(nums[:j], num-mid)
			cnt += j - i
		}
		return cnt >= k
	})
}

/*
解法3 二分查找+滑动窗口
*/
func smallestDistancePair(nums []int, k int) int {
	//排序
	sort.Ints(nums)
	//二分查找
	i, j := 0, nums[len(nums)-1]-nums[0]

	/*
		1. 求出距离对<=pivot的个数
	*/
	getCnt := func(pivot int) (cnt int) {
		left, right := 0, 0
		for left < len(nums) {
			for right < len(nums) && nums[right]-nums[left] <= pivot {
				right++
			}
			/*
				1. [left+1,right-1]范围内的下标都满足 nums[right]-nums[left] <= pivot
				2. left变大后,nums[right]-nums[left] 会变小. 所以循环时,right直接从本次循环结束时开始
				往后查找
			*/
			cnt += right - left - 1
			left++
		}
		return cnt
	}

	for i < j {
		mid := i + (j-i)>>1
		curCnt := getCnt(mid)
		if curCnt < k {
			//如果<=mid的距离对个数小于k,则应该让mid变大,继续在右区间查找
			i = mid + 1
		} else {
			/*
				1. 如果<=mid的距离对个数>k,则有可能是因为有多个距离对都等于mid.所以应该在左区间查找
				2. 如果<=mid的距离对个数等于k,则此时这个mid不一定存在,所以此时不能直接返回mid,只能继续减小二分范围
			*/
			j = mid
		}
	}

	return i
}

// @lc code=end
