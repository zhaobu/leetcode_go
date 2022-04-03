package main

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 */

// @lc code=start

/*
解法1 快速排序
*/
func sortArray1(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	parition := func(nums []int, left, right int) int {
		index := rand.Intn(right-left+1) + left             //随机选择一个位置作为本次的分区点
		nums[index], nums[right] = nums[right], nums[index] //把分区点移动到end位置,方便遍历[start,end)区间
		pivot := nums[right]                                //pivot表示区分点
		i, j := left, left
		/*
			i的左边表示已处理区间,也就是[start,i)都是<=pivot的元素
			j 用来遍历所有元素,因为pivot已经被交换到right位置,所以j只需要遍历到right-1
		*/
		for ; j < right; j++ {
			if nums[j] <= pivot { //遍历所有元素,把<=pivot的元素都移动到[start,i)
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[right] = nums[right], nums[i]
		return i
	}

	var quickSort func(nums []int, left, right int)
	quickSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		pivot := parition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}

	//设置随机种子
	rand.Seed(time.Hour.Nanoseconds())
	quickSort(nums, 0, len(nums)-1)
	return nums
}

/*
解法2 归并排序
*/
func sortArray2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	ret := make([]int, len(nums))

	merge := func(left, mid, right int) {
		start := left //区间[left,right]上最后一个已排序元素的位置后一个位置
		i, j := left, mid+1
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				ret[start] = nums[i]
				i++
			} else {
				ret[start] = nums[j]
				j++
			}
			start++
		}

		//复制剩下的元素
		for ; i <= mid; i++ {
			ret[start] = nums[i]
			start++
		}
		for ; j <= right; j++ {
			ret[start] = nums[j]
			start++
		}
		//记得最后一定要复制到原数组,保存本次排序的结果
		copy(nums[left:right+1], ret[left:right+1])
		// fmt.Printf("ret=%+v,left=%d,mid=%d,right=%d\n", ret, left, mid, right)
	}

	var mergeSort func(nums []int, left, right int)
	mergeSort = func(nums []int, left, right int) {
		if left >= right {
			return
		}
		mid := left + (right-left)>>1
		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)
		merge(left, mid, right)
	}

	mergeSort(nums, 0, len(nums)-1)

	return ret
}

/*
解法3 堆排序 小顶堆
1. 建立大小为n的小顶堆
2. 依次删除堆顶元素n次,每次删除前记录堆顶元素
*/
func sortArray3(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	//自顶向下堆化
	heapify := func(nums []int, cur, count int) {
		minPos := cur //保存当前节点和其左右节点中最小值的下标
		for {
			left, right := 2*cur+1, 2*cur+2
			if left < count && nums[left] < nums[minPos] {
				minPos = left
			}
			if right < count && nums[right] < nums[minPos] {
				minPos = right
			}
			if minPos == cur {
				break
			}
			nums[cur], nums[minPos] = nums[minPos], nums[cur]
			cur = minPos
		}
	}

	n := len(nums) //堆元素个数
	//从最后一个非叶子节点开始堆化,建小顶堆
	for i := (n >> 1) - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}

	ret := make([]int, 0, n)

	//依次删除堆顶元素n次
	for i := 1; i <= n; i++ {
		ret = append(ret, nums[0])
		nums[0], nums[n-i] = nums[n-i], nums[0] //堆顶元素和最后一个交换,相当于删除堆顶元素
		heapify(nums, 0, n-i)
	}
	return ret
}

/*
解法3 堆排序 大顶堆
1. 建立大小为n的大顶堆
2. 依次删除堆顶元素n-1次,因为每次堆顶元素都移动到最后,所以最终nums就是升序的
*/
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	//自顶向下堆化
	heapify := func(nums []int, cur, count int) {
		maxPos := cur //保存当前节点和其左右节点中最小值的下标
		for {
			left, right := 2*cur+1, 2*cur+2
			if left < count && nums[left] > nums[maxPos] {
				maxPos = left
			}
			if right < count && nums[right] > nums[maxPos] {
				maxPos = right
			}
			if maxPos == cur {
				break
			}
			nums[cur], nums[maxPos] = nums[maxPos], nums[cur]
			cur = maxPos
		}
	}

	n := len(nums) //堆元素个数
	//从最后一个非叶子节点开始堆化,建大顶堆
	for i := (n >> 1) - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}

	//依次删除堆顶元素n次
	for i := 1; i < n; i++ {
		nums[0], nums[n-i] = nums[n-i], nums[0] //堆顶元素和最后一个交换,相当于删除堆顶元素
		heapify(nums, 0, n-i)
	}
	return nums
}

// @lc code=end
