package main

/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
type MedianFinder struct {
	small []int //小顶堆,存储较大部分元素
	big   []int //大顶堆,存储较小部分元素
	count int
}

func Constructor() MedianFinder {
	return MedianFinder{
		small: []int{},
		big:   []int{},
		count: 0,
	}
}

const (
	SMALL = 0 //小顶堆
	BIG   = 1 //大顶堆
)

//往堆后面插入元素
func insert(heap []int, num, heapType int) []int {
	heap = append(heap, num)
	/*
	       0
	    1      2
	  3   4  5   6

	*/
	cur := len(heap) - 1 //当前节点
	p := (cur - 1) >> 1  //父节点
	if heapType == SMALL {
		for p >= 0 && heap[cur] < heap[p] {
			heap[cur], heap[p] = heap[p], heap[cur]
			cur = p
			p = (p - 1) >> 1
		}
	} else {
		for p >= 0 && heap[cur] > heap[p] {
			heap[cur], heap[p] = heap[p], heap[cur]
			cur = p
			p = (p - 1) >> 1
		}
	}
	return heap
}

//删除堆顶元素,并调整堆
func heapify(heap []int, heapType int) []int {
	//堆顶元素和最后一个元素交换
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[:len(heap)-1]
	i, j := 0, len(heap)-1

	if heapType == SMALL {
		for i <= j {
			tmpIdx := i
			left, right := 2*i+1, 2*i+2
			if left <= j && heap[left] < heap[tmpIdx] {
				tmpIdx = left
			}
			if right <= j && heap[right] < heap[tmpIdx] {
				tmpIdx = right
			}
			if tmpIdx == i {
				break
			}
			heap[tmpIdx], heap[i] = heap[i], heap[tmpIdx]
			i = tmpIdx
		}
	} else {
		for i <= j {
			tmpIdx := i
			left, right := 2*i+1, 2*i+2
			if left <= j && heap[left] > heap[tmpIdx] {
				tmpIdx = left
			}
			if right <= j && heap[right] > heap[tmpIdx] {
				tmpIdx = right
			}
			if tmpIdx == i {
				break
			}
			heap[tmpIdx], heap[i] = heap[i], heap[tmpIdx]
			i = tmpIdx
		}
	}

	return heap
}

/*
smallTop, bigTop 分别为小顶堆堆顶元素,大顶堆堆顶元素
1. 如果 num<=bigTop, 表明加入的元素 <= 当前中位数, 则将num加入大顶堆
2. 如果 num>bigTop 表明加入的元素 >= 当前中位数, 则将num加入小顶堆
3. 如果count=0. 则直接将元素加入小顶堆
*/
func (this *MedianFinder) AddNum(num int) {
	if this.count == 0 || num <= this.big[0] {
		this.big = insert(this.big, num, BIG)
	} else {
		this.small = insert(this.small, num, SMALL)
	}
	this.count++
	/*
	   1. 偶数个元素时,两个堆中元素个数应该相等
	   2. 奇数个元素时,大顶堆元素个数应该比小顶堆多1
	*/
	targetCount := this.count / 2 //大顶堆应该存储的元素个数
	if this.count&1 == 1 {
		targetCount += 1
	}
	// fmt.Printf("111 count=%d, targetCount=%d,len1=%d,len2=%d \n", this.count, targetCount, len(this.small), len(this.big))
	for len(this.big) > targetCount { //大顶堆元素太多
		//删除大顶堆堆顶元素
		top := this.big[0]
		this.big = heapify(this.big, BIG)
		//插入到小顶堆中
		this.small = insert(this.small, top, SMALL)
	}

	for len(this.big) < targetCount { //大顶堆元素太少
		//删除小顶堆堆顶元素
		top := this.small[0]
		this.small = heapify(this.small, SMALL)
		//插入到大顶堆中
		this.big = insert(this.big, top, BIG)
	}

	// fmt.Printf("222 count=%d, targetCount=%d,len1=%d,len2=%d \n", this.count, targetCount, len(this.small), len(this.big))
}

func (this *MedianFinder) FindMedian() float64 {
	if this.count == 0 {
		return float64(0)
	}
	if this.count&1 == 1 { //奇数个时,大顶堆堆顶就是中位数
		return float64(this.big[0])
	} else { //偶数个时,为2个堆堆顶元素的平均数
		return float64(this.small[0]+this.big[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end
