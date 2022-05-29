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
解法2 大顶堆(手写堆)
1. 用大顶堆堆化全部元素,堆顶元素保存的是出现次数最多的元素
2. 依次删除堆顶元素k次,就是出现次数前k多的元素
*/
func topKFrequent2(nums []int, k int) []int {

	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

	heapCnt := make([][2]int, 0, len(cnts))
	for num, cnt := range cnts {
		heapCnt = append(heapCnt, [2]int{num, cnt})
	}
	// fmt.Printf("cnts=%+v,heapCnt=%+v\n", cnts, heapCnt)

	/*
			自顶向下调整堆大顶堆
		            0
		        1       2
		     3     4  5   6
	*/
	heapify := func(heapCnt [][2]int, start, count int) {
		for start < count {
			pos := start
			left, right := 2*pos+1, 2*pos+2
			if left < count && heapCnt[left][1] > heapCnt[pos][1] {
				pos = left
			}
			if right < count && heapCnt[right][1] > heapCnt[pos][1] {
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
	n := len(heapCnt)
	for i := 1; i <= k; i++ {
		//堆顶元素交换到最后
		heapCnt[n-i], heapCnt[0] = heapCnt[0], heapCnt[n-i]
		//调整剩下的元素
		heapify(heapCnt, 0, k-i)
	}
	ret := make([]int, 0, k)
	for i := n - k; i < n; i++ {
		ret = append(ret, heapCnt[i][0])
	}

	return ret
}

/*
解法3 小顶堆(手写堆)
1. 设置一个容量为k的小顶堆,初始时只堆化前k个元素,堆顶元素是前k个元素中出现次数最少的
2. 从第k+1个元素开始,如果该元素出现次数比堆顶元素多,就和堆顶元素交换,并且从堆顶开始下沉,
注意此时堆中元素个数依旧为k
3. 这样到最后堆顶就是出现次数第k多的元素,堆中存储的就是出现次数前k多的所有元素
*/
func topKFrequent3(nums []int, k int) []int {

	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

	//构建元素出现次数的数组,方便进行堆操作
	heapCnt := make([][2]int, 0, len(cnts))
	for num, cnt := range cnts {
		heapCnt = append(heapCnt, [2]int{num, cnt})
	}
	// fmt.Printf("cnts=%+v,heapCnt=%+v\n", cnts, heapCnt)

	/*
			1. 自顶向下调整堆小顶堆
			2. 小顶堆的堆顶为出现次数第k多的元素,堆里面的元素出现次数都>=堆顶元素出现的次数
		            0
		        1       2
		     3     4  5   6
	*/
	heapify := func(heapCnt [][2]int, start, count int) {
		for start < count {
			pos := start
			left, right := 2*pos+1, 2*pos+2
			if left < count && heapCnt[left][1] < heapCnt[pos][1] {
				pos = left
			}
			if right < count && heapCnt[right][1] < heapCnt[pos][1] {
				pos = right
			}
			if pos == start {
				break
			}
			heapCnt[pos], heapCnt[start] = heapCnt[start], heapCnt[pos]
			start = pos
		}
	}

	//堆的大小为k,从最后一个非叶子节点开始堆化数组
	for p := (k >> 1) - 1; p >= 0; p-- {
		heapify(heapCnt, p, k)
	}

	//从第k+1个元素开始每个元素也入堆一次
	for i := k; i < len(heapCnt); i++ {
		/*
			和堆顶元素比较,如果出现次数大于堆顶元素,就和堆顶元素交换,
			相当于删除堆顶元素,然后堆顶元素下沉
		*/
		if heapCnt[i][1] > heapCnt[0][1] {
			heapCnt[i], heapCnt[0] = heapCnt[0], heapCnt[i]
			heapify(heapCnt, 0, k)
		}
	}

	ret := make([]int, 0, k)
	//堆中元素就是要求的结果
	for i := 0; i < k; i++ {
		ret = append(ret, heapCnt[i][0])
	}

	return ret
}

/*
解法4 堆(库函数)
go的堆实现不支持获取堆顶元素,只支持pop,所以必须先push,超过k个元素后再pop
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
1. 往堆尾插入新元素
2. 库函数Push实现是先调用此Push,插入到最后,然后上沉
*/
func (h *CntHeap) Push(data interface{}) {
	*h = append(*h, data.([2]int))
}

/*
1. 弹出堆尾元素
2. 库函数Pop实现是堆顶元素和最后一个元素交换,然后从堆顶开始下沉,最后调用
此函数弹出最后一个元素
*/
func (h *CntHeap) Pop() interface{} {
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

/*
解法5 二分查找
1. n表示所有元素出现的次数, n的取值范围为[minCnt, maxCnt]
2. 计算出出现次数>=n次的元素个数curCnt, n越小,curCnt越小,n越大,curCnt越大,满足单调递增特性
3. 二分查找第一个curCnt=k的n的值
4. 二分查找的初始范围是[minCnt, maxCnt]
*/
func topKFrequent(nums []int, k int) []int {

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
	//求出出现次数>=maxCnt的元素个数
	getCnt := func(maxCnt int) (count int) {
		for _, cnt := range cnts {
			if cnt >= maxCnt {
				count++
			}
		}
		return count
	}

	left, right := minCnt, maxCnt
	mid := 0 //表示元素出现次数的最小值
	for left <= right {
		mid = left + (right-left)>>1
		curCnt := getCnt(mid)
		if curCnt < k { //如果元素过少,应该让mid变小,
			right = mid - 1
		} else if curCnt > k { //如果元素太多,应该让mid变大
			left = mid + 1
		} else {
			break
		}
	}

	//过滤出出现次数前k的元素
	ret := make([]int, 0, k)
	for num, cnt := range cnts {
		if cnt >= mid {
			ret = append(ret, num)
		}
	}
	return ret
}

/*
解法6 快速排序
1. 统计每个元素出现的次数
2. 构建元素出现次数的数组 numsCnt, numsCnt[i][0]表示元素的值,numsCnt[i][1]表示元素出现的次数
3. 选一个分区点 pivot,pivot左边的元素出现次数都<numsCnt[pivot][1],pivot右边的元素出现次数都>numsCnt[pivot][1]
4. 求出 >=pivot 的元素个数rightCnt,这些元素出现的次数都 >=numsCnt[i][1].
如果rightCnt<k就继续在左分区找,如果rightCnt>k就继续在右分区找
*/
func topKFrequent6(nums []int, k int) []int {

	/*
		1. 统计每个元素出现的次数
		2. 初始化为k尽量减少map扩容次数
	*/
	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}

	//构建元素出现次数的数组
	numsCnt := make([][2]int, 0, len(cnts))
	for num, cnt := range cnts {
		numsCnt = append(numsCnt, [2]int{num, cnt})
	}
	/*
		1. 分区函数,把出现次数小于pivotCnt的放左边,出现次数大于pivotCnt的放右边
		2. 返回值pivot表示分区元素的下标
	*/
	parition := func(numsCnt [][2]int, left, right int) (pivot int) {
		pivotCnt := numsCnt[right][1]
		pivot = left //第一个出现次数>=pivotCnt的下标
		for i := left; i < right; i++ {
			if numsCnt[i][1] < pivotCnt {
				numsCnt[i], numsCnt[pivot] = numsCnt[pivot], numsCnt[i]
				pivot++
			}
		}
		numsCnt[pivot], numsCnt[right] = numsCnt[right], numsCnt[pivot]
		return pivot
	}

	left, right := 0, len(numsCnt)-1

	/*
		1. pivot 表示分区元素的下标,pivot左边的元素出现次数都<numsCnt[pivot][1],
		pivot右边的元素出现次数都>numsCnt[pivot][1]
		2. rightCnt 表示出现次数 >= numsCnt[pivot] 元素个数
	*/
	pivot, rightCnt := 0, 0
	for left <= right && rightCnt != k {
		pivot = parition(numsCnt, left, right)
		/*
			1. 表示出现次数>= numsCnt[pivot] 的元素个数
			2. [pivot,len(numsCnt)) 所有元素都满足条件.而不仅仅是[pivot,right]区间
		*/
		rightCnt = len(numsCnt) - pivot
		if rightCnt < k {
			// 如果元素太少,应该减小pivot,让更多的元素满足条件
			right = pivot - 1
		} else if rightCnt > k {
			// 如果元素太多,应该增大pivot,减少满足条件的元素
			left = pivot + 1
		}
	}

	//快排结束后[pivot,len(numsCnt))区间所有元素就是出现次数最多的k个元素
	ret := make([]int, 0, k)
	for i := pivot; i < len(numsCnt); i++ {
		ret = append(ret, numsCnt[i][0])
	}
	// fmt.Printf("numsCnt=%+v\n", numsCnt[pivot:])
	return ret
}


/* 
总结:
1. 堆实现时,推荐用小顶堆实现
2. 二分查找和快速排序思想类似,只不过二分是时间换空间,每次都重复计算元素个数
快排是时间还空间,直接保存每个元素出现的次数
*/
// @lc code=end
