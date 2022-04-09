package main

import (
	"container/heap"
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start

/*
解法1 优先级对列 使用库自带的优先级对列
可以参考:https://leetcode-cn.com/problems/find-k-pairs-with-smallest-sums/solution/bu-chong-guan-fang-ti-jie-you-xian-dui-l-htf8/

1. 构建大小为min(k,m)的小顶堆
2. 因为nums1,nums2都是升序数组,要求的是前k个和最小的数对,假设（i,j）表示的是nums1下标i与nums2下标j
则一定有(i,j)和(i,j+1),(i+1,j),(i+1,j+1)四个中最小的,并且有:
(i,j)<(i,j+1)<(i+1,j+1)
(i,j)<(i+1,j)<(i+1,j+1)
把i看成行,j看成列,初始化时,初始化每一行, 如果m<K. 只需要初始化前m行即可,然后每次弹出栈时,再把j相加
不断往右下角走
*/
func kSmallestPairs1(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	//初始化每一行
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}

	fmt.Printf("hp:%+v\n", h.data)
	//遍历每一行,列不断相加
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		fmt.Printf("Pop:%+v\n", p)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
			fmt.Printf("Push:%+v\n", pair{i, j + 1})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}
func (h *hp) Push(v interface{}) {
	h.data = append(h.data, v.(pair))
}
func (h *hp) Pop() interface{} {
	a := h.data
	v := a[len(a)-1]
	h.data = a[:len(a)-1]
	return v
}

/*
解法2 优先级对列 自己实现二叉堆
思路和解法1一样,自己实现堆
*/
func kSmallestPairs2(nums1, nums2 []int, k int) (ans [][]int) {

	type Data struct {
		i int
		j int
	}

	m, n := len(nums1), len(nums2)
	ret := make([][]int, 0, k)

	heap := make([]*Data, 0, k)

	less := func(d1, d2 *Data) bool {
		return nums1[d1.i]+nums2[d1.j] < nums1[d2.i]+nums2[d2.j]
	}

	insert := func(data *Data) {
		//往堆中插入元素
		heap = append(heap, data)
		//堆化
		i := len(heap) - 1
		p := (i - 1) >> 1
		for p >= 0 && less(heap[i], heap[p]) {
			heap[i], heap[p] = heap[p], heap[i]
			i = p
			p = (i - 1) >> 1
		}
	}

	heapify := func(start int) {
		count := len(heap)
		for {
			minPos := start
			left, right := 2*start+1, 2*start+2
			if left < count && less(heap[left], heap[minPos]) {
				minPos = left
			}
			if right < count && less(heap[right], heap[minPos]) {
				minPos = right
			}
			if minPos == start {
				break
			}
			heap[start], heap[minPos] = heap[minPos], heap[start]
			start = minPos
		}
	}

	//构建大小为k的小顶堆
	for i := 0; i < k && i < m; i++ {
		insert(&Data{i: i, j: 0})
	}

	//不断删除堆顶元素k次,并重新让新元素入堆
	for len(heap) > 0 && k > 0 {
		top := heap[0]
		ret = append(ret, []int{nums1[top.i], nums2[top.j]})
		k--
		if top.j+1 < n {
			//如果删除堆顶元素后还要插入新元素,直接替换堆顶相当于删除
			heap[0] = &Data{i: top.i, j: top.j + 1}
		} else {
			//如果删除堆顶元素后不需要插入新元素
			heap = heap[1:]
		}
		//重新调整堆
		heapify(0)
	}
	return ret
}

/*
解法3 二分查找
*/
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	ans = make([][]int, 0, k)
	// 二分查找第 k 小的数对和
	left, right := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]+1
	/*
		1. sort.Search里面用二分是mid的取值范围是[0,right-left),也就是把sum看成是
		二分时的下标.
		2. sum+=left后sum的范围就是[left,right)
		3. sort.Search求得的是最小的sum满足小于等于left+sum的数对数目刚好大于等于 k
	*/
	pairSum := left + sort.Search(right-left, func(sum int) bool {
		sum += left
		cnt := 0
		i, j := 0, n-1
		for i < m && j >= 0 {
			if nums1[i]+nums2[j] > sum {
				j--
			} else {
				/*
					1. 如果nums1[i]+nums2[j]<=sum,则j的取值范围[0,j]之内都满足.
					2. 所以对于每一个i,都有j+1个数对满足条件
				*/
				cnt += j + 1
				i++
			}
		}
		return cnt >= k
	})

	// 找数对和小于 pairSum 的数对
	i := n - 1
	for _, num1 := range nums1 {
		/*
			因为num1是从小到大遍历的,所以当num1变大时,num2的第一个满足条件的下标肯定会变小
			所以i不用每次都从n-1开始遍历,只需要才能继续从上一个num1求的的i判断
		*/
		for i >= 0 && num1+nums2[i] >= pairSum {
			i--
		}
		for _, num2 := range nums2[:i+1] {
			ans = append(ans, []int{num1, num2})
			if len(ans) == k {
				return
			}
		}
	}

	// 找数对和等于 pairSum 的数对
	i = n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] > pairSum {
			i--
		}
		for j := i; j >= 0 && num1+nums2[j] == pairSum; j-- {
			ans = append(ans, []int{num1, nums2[j]})
			if len(ans) == k {
				return
			}
		}
	}
	return
}

// @lc code=end
