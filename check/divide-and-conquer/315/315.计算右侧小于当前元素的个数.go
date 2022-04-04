package main

import "fmt"

/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start

/*
解法1 归并排序变形,升序
1. 利用归并排序merge时做统计,merge时两个数组都已经是升序的,并且左边部分的元素在原数组中也同样是在右边部分的左边
2. 在merge和合并左边部分A和右边部分B时,如果左边元素a<=右边元素b,则本次应该合并a.
可以确定此时因为数组是升序排列,从[mid+1,j)之间的元素已经被加入到已合并部分,
所以也就表明从[mid+1,j)也就是j-mid-1个元素全部都是小于元素a的,
*/
func countSmaller1(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	if n == 1 {
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
				ret[tmpIndex[i]] += j - (mid + 1)
				nums[pos] = tmpNums[i]
				numsIndex[pos] = tmpIndex[i]
				i++
			} else if tmpNums[i] <= tmpNums[j] { //左边元素较小
				ret[tmpIndex[i]] += j - (mid + 1)
				nums[pos] = tmpNums[i]
				numsIndex[pos] = tmpIndex[i]
				i++
			} else { //右边元素较小
				nums[pos] = tmpNums[j]
				numsIndex[pos] = tmpIndex[j]
				j++
			}
		}
		fmt.Printf("start=%d,end=%d,nums=%+v,numsIndex=%+v\n", start, end, nums, numsIndex)
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

	// fmt.Printf("nums=%+v,numsIndex=%+v\n", nums, numsIndex)
	return ret
}

/*
解法1 归并排序变形,降序
*/
func countSmaller(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	if n == 1 {
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
				ret[tmpIndex[i]] += end - j + 1
				nums[pos] = tmpNums[i]
				numsIndex[pos] = tmpIndex[i]
				i++
			} else if tmpNums[i] > tmpNums[j] { //左边元素较大
				ret[tmpIndex[i]] += end - j + 1
				nums[pos] = tmpNums[i]
				numsIndex[pos] = tmpIndex[i]
				i++
			} else { //右边元素较大
				nums[pos] = tmpNums[j]
				numsIndex[pos] = tmpIndex[j]
				j++
			}
		}
		// fmt.Printf("start=%d,end=%d,nums=%+v,numsIndex=%+v\n", start, end, nums, numsIndex)
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

	// fmt.Printf("nums=%+v,numsIndex=%+v\n", nums, numsIndex)
	return ret
}

// @lc code=end
