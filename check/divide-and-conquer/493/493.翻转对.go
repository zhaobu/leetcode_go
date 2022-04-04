package main

/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start

/*
解法1 归并排序 升序
*/
func reversePairs(nums []int) int {
	ret := 0
	if len(nums) < 2 {
		return ret
	}

	tmp := make([]int, len(nums))

	merge := func(nums []int, left, mid, right int) {
		copy(tmp[left:right+1], nums[left:right+1])

		/*
		 1. 合并前判断重要翻转对
		 2. 因为2个待合并的左右部分都是升序的,对于 nums[i], left <= i <= mid，我们在找到的符合 nums[i] > 2*nums[j] 的 nums[j],
		  mid+1 <= j <= right，也必然也符合 nums[i+1] > 2*nums[j]
		*/
		j := mid + 1
		for i := left; i <= mid; i++ {
			/*
				1. 对于每一个nums[i],如果nums[i] > 2*nums[j],因为nums[j]左边的都小于nums[j],所以也必定满足,右边的必定不满足
				2. 当i变为i+1时,j不需要重头开始,因为nums[i]变大,上一次循环结束的j位置处必定满足nums[i] > 2*nums[j],j只需要继续往后遍历即可
			*/
			for ; j <= right && nums[i] > 2*nums[j]; j++ {
			}
			ret += j - (mid + 1)
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

	mergeSort(nums, 0, len(nums)-1)
	return ret
}

// @lc code=end
