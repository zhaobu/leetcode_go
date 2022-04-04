package main

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * 剑指 Offer 51. 数组中的逆序对
 */

// @lc code=start

/*
解法1 归并排序
*/
func reversePairs(nums []int) int {
	n := len(nums)
	ret := 0
	if n < 2 {
		return ret
	}

	//用来记录nums数组对应元素的下标
	numsIndex := make([]int, n)
	for i := range nums {
		numsIndex[i] = i
	}

	tmpNums := make([]int, n)  //用来帮助排序
	tmpIndex := make([]int, n) //用来帮助排序过程中元素下标的变化

	merge := func(nums []int, start, mid, end int) {
		copy(tmpNums[start:end+1], nums[start:end+1])
		copy(tmpIndex[start:end+1], numsIndex[start:end+1])

		pos := start
		for i, j := start, mid+1; pos <= end; pos++ {
			if i == mid+1 { //[start,mid]已经遍历完
				nums[pos] = tmpNums[j]
				numsIndex[pos] = tmpIndex[j]
				j++
			} else if j == end+1 { //[mid+1,end]已遍历完
				ret += j - (mid + 1)
				nums[pos] = tmpNums[i]
				numsIndex[pos] = tmpIndex[i]
				i++
			} else if tmpNums[i] <= tmpNums[j] { //左边元素较小
				ret += j - (mid + 1)
				nums[pos] = tmpNums[i]
				numsIndex[pos] = tmpIndex[i]
				i++
			} else { //右边元素较小
				nums[pos] = tmpNums[j]
				numsIndex[pos] = tmpIndex[j]
				j++
			}
		}
		//fmt.Printf("start=%d,end=%d,nums=%+v,numsIndex=%+v\n", start, end, nums, numsIndex)
	}

	var mergeSort func(nums []int, start, end int)
	mergeSort = func(nums []int, start, end int) {
		if start >= end {
			return
		}
		mid := start + (end-start)>>1
		mergeSort(nums, start, mid)
		mergeSort(nums, mid+1, end)
		merge(nums, start, mid, end)
	}

	mergeSort(nums, 0, n-1)

	//fmt.Printf("nums=%+v,numsIndex=%+v\n", nums, numsIndex)
	return ret
}

// @lc code=end
