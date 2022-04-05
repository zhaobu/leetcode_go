package main

/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */

// @lc code=start

/*
解法1 归并排序 (暴力求解,超时)
*/
func countRangeSum1(nums []int, lower int, upper int) int {
	if len(nums) == 1 {
		return 1
	}

	ret := 0
	preSum := make([]int, len(nums)+1) //前缀和数组
	for i, v := range nums {
		preSum[i+1] += preSum[i] + v
	}
	// fmt.Printf("preSum=%+v\n", preSum)

	tmp := make([]int, len(nums)+1) //归并排序临时数组

	merge := func(nums []int, left, mid, right int) {
		copy(tmp[left:right+1], nums[left:right+1])

		/*
			利用求前缀和的技巧统计满足条件的区间和
		*/
		for i := left; i <= mid; i++ {
			for j := mid + 1; j <= right; j++ {
				delta := preSum[j] - preSum[i]
				// fmt.Printf("i=%d,j=%d,delta=%d\n", i, j, delta)
				if delta >= lower && delta <= upper {
					ret++
				}
			}
		}

		pos := left
		for i, j := left, mid+1; pos <= right; pos++ {
			if i == mid+1 { //左边先合并完
				nums[pos] = tmp[j]
				j++
			} else if j == right+1 { //右边先合并完
				nums[pos] = tmp[i]
				i++
			} else if tmp[i] <= tmp[j] { //左边较小
				nums[pos] = tmp[i]
				i++
			} else { //右边较小
				nums[pos] = tmp[j]
				j++
			}
		}
		// fmt.Printf("left=%d,mid=%d,right=%d,nums=%+v\n", left, mid, right, nums)
	}

	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)>>1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}

	mergeSort(preSum, 0, len(preSum)-1)
	return ret
}

/*
解法1 归并排序 优化
*/
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums) == 1 {
		return 1
	}

	ret := 0
	preSum := make([]int, len(nums)+1) //前缀和数组
	for i, v := range nums {
		preSum[i+1] += preSum[i] + v
	}
	// fmt.Printf("preSum=%+v\n", preSum)

	tmp := make([]int, len(nums)+1) //归并排序临时数组

	merge := func(nums []int, left, mid, right int) {
		copy(tmp[left:right+1], nums[left:right+1])

		start, end := mid+1, mid+1 //[start,end)为满足条件的以nums[i]为左边界的区间和的右边界的范围
		for i := left; i <= mid; i++ {
			/*
				1. i作为要求的区间和的左下标,nums[i]对应的满足条件的右下标的区间是 [start, end)
				2. 当i变为i+1时,j不需要重头开始,因为nums[i]变大,nums[i+1] 对应的区间一定会整体右移，类似滑动窗口
			*/
			for ; start <= right && nums[start]-nums[i] < lower; start++ {
			}
			for ; end <= right && nums[end]-nums[i] <= upper; end++ {
			}
			ret += end - start
		}

		pos := left
		for i, j := left, mid+1; pos <= right; pos++ {
			if i == mid+1 { //左边先合并完
				nums[pos] = tmp[j]
				j++
			} else if j == right+1 { //右边先合并完
				nums[pos] = tmp[i]
				i++
			} else if tmp[i] <= tmp[j] { //左边较小
				nums[pos] = tmp[i]
				i++
			} else { //右边较小
				nums[pos] = tmp[j]
				j++
			}
		}
		// fmt.Printf("left=%d,mid=%d,right=%d,nums=%+v\n", left, mid, right, nums)
	}

	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)>>1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(nums, left, mid, right)
	}

	mergeSort(preSum, 0, len(preSum)-1)
	return ret
}

// @lc code=end
