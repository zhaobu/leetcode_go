package main

import (
	"container/heap"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start

/*
解法1 桶排序
1. 统计每个元素出现的次数
2. 找出最少次数 minCnt 和最多次数 maxCnt
3. 构建大小 maxCnt-minCnt+1 的二维桶,一维下标是出现的次数,二维是出现该次数的所有元素
4. 取出桶中出现次数前k的所有元素
*/

func topKFrequent1(nums []int, k int) []int {
	/*
		1. 统计每个元素出现的次数
		2. 初始化为k尽量减少map扩容次数
	*/
	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}
	//找出最少次数和最多次数
	minCnt, maxCnt := len(nums), 0
	for _, v := range cnts {
		if v < minCnt {
			minCnt = v
		}
		if v > maxCnt {
			maxCnt = v
		}
	}
	/*
	   1. 构建大小 maxCnt-minCnt+1的二维切片
	   2. 切片的一维下标是出现的次数,二维是出现该次数的所有元素
	*/
	bucket := make([][]int, maxCnt-minCnt+1)
	// fmt.Printf("maxCnt=%d, minCnt=%d, cnts=%+v\n", maxCnt, minCnt, cnts)
	//遍历每个元素出现的次数,依次入桶
	for num, cnt := range cnts {
		bucket[cnt-minCnt] = append(bucket[cnt-minCnt], num)
	}

	//取出桶中出现次数前k的所有元素
	ret := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(ret)+len(bucket[i]) > k { //最后一次只需要加入剩下的元素即可
			ret = append(ret, bucket[i][:k-len(ret)]...)
			break
		}
		//如果加入此桶元素后还不足就全部加入
		ret = append(ret, bucket[i]...)
	}
	return ret
}

/*
解法2 大顶堆(自己实现)
*/
func topKFrequent2(nums []int, k int) []int {

	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

	type Cnt struct {
		num   int
		count int
	}

	heapCnt := make([]*Cnt, 0, len(cnts))

	for num, cnt := range cnts {
		heapCnt = append(heapCnt, &Cnt{num: num, count: cnt})
	}
	// fmt.Printf("cnts=%+v,heapCnt=%+v\n", cnts, heapCnt)

	/*
			自顶向下调整堆大顶堆
		            0
		        1       2
		     3     4  5   6
	*/
	heapify := func(heapCnt []*Cnt, start, count int) {
		for start < count {
			pos := start
			left, right := 2*pos+1, 2*pos+2
			if left < count && heapCnt[left].count > heapCnt[pos].count {
				pos = left
			}
			if right < count && heapCnt[right].count > heapCnt[pos].count {
				pos = right
			}
			if pos == start {
				break
			}
			heapCnt[pos], heapCnt[start] = heapCnt[start], heapCnt[pos]
			start = pos
		}
	}

	//从最后一个非叶子节点开始堆化数组
	for p := (len(heapCnt) >> 1) - 1; p >= 0; p-- {
		heapify(heapCnt, p, len(heapCnt))
	}

	//删除堆顶元素k次
	count := len(heapCnt)
	for i := 0; i < k; i++ {
		//堆顶元素交换到最后
		count--
		heapCnt[count], heapCnt[0] = heapCnt[0], heapCnt[count]
		//调整剩下的元素
		heapify(heapCnt, 0, count)
	}
	ret := make([]int, 0, k)
	for i := len(heapCnt) - 1; i >= 0 && len(ret) < k; i-- {
		ret = append(ret, heapCnt[i].num)
	}

	return ret
}

/*
解法3 小顶堆(自己实现)
*/
func topKFrequent3(nums []int, k int) []int {

	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

	type Cnt struct {
		num   int
		count int
	}

	heapCnt := make([]*Cnt, 0, len(cnts))

	for num, cnt := range cnts {
		heapCnt = append(heapCnt, &Cnt{num: num, count: cnt})
	}
	// fmt.Printf("cnts=%+v,heapCnt=%+v\n", cnts, heapCnt)

	/*
			自顶向下调整堆大顶堆
		            0
		        1       2
		     3     4  5   6
	*/
	heapify := func(heapCnt []*Cnt, start, count int) {
		for start < count {
			pos := start
			left, right := 2*pos+1, 2*pos+2
			if left < count && heapCnt[left].count > heapCnt[pos].count {
				pos = left
			}
			if right < count && heapCnt[right].count > heapCnt[pos].count {
				pos = right
			}
			if pos == start {
				break
			}
			heapCnt[pos], heapCnt[start] = heapCnt[start], heapCnt[pos]
			start = pos
		}
	}

	//从最后一个非叶子节点开始堆化数组
	for p := (len(heapCnt) >> 1) - 1; p >= 0; p-- {
		heapify(heapCnt, p, len(heapCnt))
	}

	//删除堆顶元素k次
	count := len(heapCnt)
	for i := 0; i < k; i++ {
		//堆顶元素交换到最后
		count--
		heapCnt[count], heapCnt[0] = heapCnt[0], heapCnt[count]
		//调整剩下的元素
		heapify(heapCnt, 0, count)
	}
	ret := make([]int, 0, k)
	for i := len(heapCnt) - 1; i >= 0 && len(ret) < k; i-- {
		ret = append(ret, heapCnt[i].num)
	}

	return ret
}

/*
解法4 堆(库函数)
*/
func topKFrequent4(nums []int, k int) []int {

	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

	//构造大小为k的小顶堆
	heapCnt := &CntHeap{}
	heap.Init(heapCnt)

	for num, cnt := range cnts {
		heap.Push(heapCnt, [2]int{num, cnt})
		//如果元素超过k个就弹出堆顶元素
		if heapCnt.Len() > k {
			heap.Pop(heapCnt)
		}
	}
	ret := make([]int, 0, k)
	// 弹出全部堆中元素
	for heapCnt.Len() > 0 {
		ret = append(ret, heap.Pop(heapCnt).([2]int)[0])
	}
	return ret
}

//堆的定义下标0表示元素值,下标1表示出现次数
type CntHeap [][2]int

func (h CntHeap) Len() int           { return len(h) }
func (h CntHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h CntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

/*
1.往堆尾插入新元素
*/
func (h *CntHeap) Push(data interface{}) {
	*h = append(*h, data.([2]int))
}

//弹出堆尾元素
func (h *CntHeap) Pop() interface{} {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

/*
解法5 二分查找
*/
func topKFrequent(nums []int, k int) []int {

	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

}

// @lc code=end
