package main

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start

/*
解法1 快排思想
类似快慢指针形式的快排写法
*/
func findKthLargest1(nums []int, k int) int {

	rank := -1 //本轮确定的已排序元素的下标
	k = k - 1  //最终要确定的下标
	left, right := 0, len(nums)-1

	var parition func(start, end int) int
	parition = func(start, end int) int {
		index := rand.Intn(end-start+1) + start         //在[start,end]区间随机选一个作为分区点
		nums[end], nums[index] = nums[index], nums[end] //把选取的分区点和最后一个位置交换,方便下面遍历
		pivot := nums[end]                              //记录分区点的位置
		i := start                                      //nums[i]表示已处理区间的后一个元素
		for j := start; j < end; j++ {                  //注意循环区间应该是[start,end)
			if nums[j] > pivot { //把所有>= pivot的元素移动到区间[start,i)
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		//遍历完成后i就是pivot所在的下标
		nums[i], nums[end] = nums[end], nums[i]
		return i
	}

	rand.Seed(time.Now().UnixNano())
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

/*
解法2 快排思想
填坑形式的快排写法
*/
func findKthLargest2(nums []int, k int) int {

	rank := -1 //本轮确定的已排序元素的下标
	k = k - 1  //最终要确定的下标
	left, right := 0, len(nums)-1

	var parition func(start, end int) int
	parition = func(start, end int) int {
		index := rand.Intn(end-start+1) + start //在[start,end]区间随机选一个作为分区点
		pivot := nums[index]                    //记录分区点的位置
		i, j := start, end                      //要处理的区间范围
		/*
			1. pivot已经保存了nums[index]处元素的值,所以可以理解成在数组nums[index]上挖了个坑，
			可以将其它数据填充到这来
			2. 本来第一个坑应该是index位置,互换index和i,这样第一个坑就是i
		*/
		nums[i], nums[index] = nums[index], nums[i]
		for i < j {
			// 从右向左找第一个>=pivot的位置
			for i < j && nums[j] < pivot {
				j--
			}
			if i < j {
				nums[i] = nums[j] //将nums[j]填到nums[i]中，nums[j]就形成了一个新的坑
				i++

			}
			// 从左向右找第一个小于pivot的位置
			for i < j && nums[i] >= pivot {
				i++
			}
			if i < j {
				nums[j] = nums[i] //将nums[i]填到nums[j]中，nums[i]就形成了一个新的坑
				j--
			}
		}
		//退出时，i等于j。将pivot填到这个坑中。
		nums[i] = pivot
		return i
	}
	rand.Seed(time.Now().UnixNano())
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

/*
解法3 二叉堆 小顶堆
如何在一个包含 n 个数据的数组中，查找前 K 大数据呢?
维护一个大小为 K 的小顶堆，顺序遍历数组，从数组中取出取数据
1. 如果堆中元素<k,元素直接入堆
2. 如果堆中元素=k则与堆顶元素比较:
如果比堆顶元素大，我们就把堆顶元素删除，并且将这个元素插入到堆中
如果比堆顶元素小，则不做处理,继续遍历数组。
数组中的数据都遍历完之后，堆中的数据就是前 K 大数据
*/
func findKthLargest3(nums []int, k int) int {
	heap := make([]int, k) //大小为k的小顶堆
	count := 0             //堆中元素个数

	// 自上往下堆化
	heapify := func(i int) {
		for {
			minPos := i //当前节点和2个子节点中最小元素的下标
			if (2*i+1) < count && heap[minPos] > heap[2*i+1] {
				minPos = 2*i + 1
			}
			if (2*i+2) < count && heap[minPos] > heap[2*i+2] {
				minPos = 2*i + 2
			}
			if minPos == i {
				break
			}
			heap[i], heap[minPos] = heap[minPos], heap[i] //交换最大节点和当前节点
			i = minPos
		}
	}

	//自下往上堆化
	insert := func(data int) {
		heap[count] = data
		count++
		i := count - 1    //当前节点从堆的最后一个元素开始
		p := (i - 1) >> 1 //父节点下标
		for p >= 0 && heap[i] < heap[p] {
			heap[i], heap[p] = heap[p], heap[i] //交换父节点和当前节点
			i = p
			p = (i - 1) >> 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if count < k { //如果堆不满,插入新元素,从下到上堆化
			insert(nums[i])
		} else if nums[i] > heap[0] { //如果堆满,直接用新元素替换堆顶元素,相当于删除,然后再从上到下堆化
			heap[0] = nums[i]
			heapify(0)
		}

	}

	return heap[0]
}

/*
解法4 二叉堆 小顶堆
思想和解法2一样,只不过写法不一样,先建好大小为k的小顶堆,然后再遍历剩下的元素
*/
func findKthLargest(nums []int, k int) int {

	// 自上往下堆化
	heapify := func(heap []int, count, i int) {
		for {
			minPos := i //当前节点和2个子节点中最小元素的下标
			if (2*i+1) < count && heap[minPos] > heap[2*i+1] {
				minPos = 2*i + 1
			}
			if (2*i+2) < count && heap[minPos] > heap[2*i+2] {
				minPos = 2*i + 2
			}
			if minPos == i {
				break
			}
			heap[i], heap[minPos] = heap[minPos], heap[i] //交换最大节点和当前节点
			i = minPos
		}
	}

	//建大小为k的小顶堆
	for i := k>>1 - 1; i >= 0; i-- {
		heapify(nums, k, i)
	}

	//后面每个元素都和堆顶元素比较
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i] //直接用新元素替换堆顶元素,相当于删除,然后再从上到下堆化
			heapify(nums, k, 0)
		}

	}

	return nums[0]
}

/*
解法5 二叉堆 小顶堆
1. 遍历所有元素堆化构建大顶堆
2. 删除n-k次堆顶元素,剩下k个元素时,堆顶元素就是第k大的元素
*/
func findKthLargest5(nums []int, k int) int {

	// 自上往下堆化
	heapify := func(heap []int, count, i int) {
		for {
			minPos := i //当前节点和2个子节点中最小元素的下标
			if (2*i+1) < count && heap[minPos] > heap[2*i+1] {
				minPos = 2*i + 1
			}
			if (2*i+2) < count && heap[minPos] > heap[2*i+2] {
				minPos = 2*i + 2
			}
			if minPos == i {
				break
			}
			heap[i], heap[minPos] = heap[minPos], heap[i] //交换最大节点和当前节点
			i = minPos
		}
	}

	n := len(nums) //堆中元素个数
	//建堆
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	m := n - k
	//删除堆顶元素m次
	for i := 1; i <= m; i++ {
		//把最后一个元素和堆顶元素交换,相当于删除堆顶元素,然后自上往下堆化
		nums[0], nums[n-i] = nums[n-i], nums[0]
		heapify(nums, n-i, 0)
	}
	return nums[0]
}

/*
解法6 二叉堆 大顶堆
1. 遍历所有元素堆化构建大顶堆
2. 删除堆顶元素k-1次,堆顶元素就是第k大的元素
*/
func findKthLargest6(nums []int, k int) int {

	// 自上往下堆化
	heapify := func(heap []int, count, i int) {
		for {
			maxPos := i //当前节点和2个子节点中最大元素的下标
			if (2*i+1) < count && heap[maxPos] < heap[2*i+1] {
				maxPos = 2*i + 1
			}
			if (2*i+2) < count && heap[maxPos] < heap[2*i+2] {
				maxPos = 2*i + 2
			}
			if maxPos == i {
				break
			}
			heap[i], heap[maxPos] = heap[maxPos], heap[i] //交换最大节点和当前节点
			i = maxPos
		}
	}

	n := len(nums) //堆中元素个数
	//建堆
	for i := n>>1 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}

	m := k - 1
	//删除堆顶元素m次
	for i := 1; i <= m; i++ {
		//把最后一个元素和堆顶元素交换,相当于删除堆顶元素,然后自上往下堆化
		nums[0], nums[n-i] = nums[n-i], nums[0]
		heapify(nums, n-i, 0)
	}
	return nums[0]
}

// @lc code=end
