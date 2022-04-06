package main

import "container/heap"

/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的 K 对数字
 */

// @lc code=start

/*
解法1 优先级对列 使用库自带的优先级对列
*/
func kSmallestPairs1(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
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
*/
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	ret := make([][]int, 0, k)

	i, j := 0, 0
	m, n := len(nums1), len(nums2)
	for (i < m && j < n) && k > 0 {
		ret = append(ret, []int{nums1[i], nums2[j]})
		k--
		if i+1 >= m {
			j++
		} else if j+1 >= n {
			i++
		} else {
			if nums1[i+1]+nums2[j] <= nums1[i]+nums2[j+1] {
				i++
			} else {
				j++
			}
		}
	}

	return ret
}

// @lc code=end
