package main

import (
	"container/heap"
	"fmt"
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
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
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
func kSmallestPairs1(nums1, nums2 []int, k int) (ans [][]int) {

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
		// heap = append(heap, data)
		// //堆化
		// i := len(heap) - 1
		// for (i>>1) > 0 && less(heap[]) {

		// }
	}

	heapify := func(start int) {

	}

	//往堆中插入k个元素
	for i := 0; i < k && i < m; i++ {
		heap = append(heap, &Data{i: i, j: 0})
	}

	//自顶向下堆化
	heapify(0)

	return ret
}

// @lc code=end
