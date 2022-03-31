package main

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {

	rank := -1 //本轮确定的已排序元素的下标
	k = k - 1  //最终要确定的下标
	left, right := 0, len(nums)-1

	var parition func(start, end int) int
	parition = func(start, end int) int {
		pivot := nums[end] //左边>=pivot, 右边<pivot
		i := start         //nums[i]表示已处理区间的最后一个元素
		j := start
		for ; j < end; j++ {
			if nums[j] > pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		// if i > start {
		nums[i], nums[end] = nums[end], nums[i]
		// }
		return i
	}

	for rank != k && left <= right {
		rank = parition(left, right)
		if rank > k {
			right = rank - 1
		} else if rank < k {
			left = rank + 1
		}
	}

	return nums[rank]
}

// @lc code=end
