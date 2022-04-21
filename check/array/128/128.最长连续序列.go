package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start

/*
解法1 hash表
*/
func longestConsecutive1(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	record := map[int]bool{} //用来记录哪些元素存在
	for i := 0; i < len(nums); i++ {
		record[nums[i]] = true
	}

	ret := 0
	for cur := range record {
		/*
			1. 比如说假如存在1,2,3,4 这些数字,那只需要在cur=1时才需要进到内存循环里面计算一次
			2. 因为要计算的是最长的连续序列,从1开始到4肯定比从其他数字开始到4计算出长度要长
			3. 因为对于每一段连续的序列,都是在外层循环遍历到序列最小的元素时才需要进入到内层循环
			同时,内层循环又是直接往后遍历到该序列的最大值,所以对于每一个元素都会在内外层循环各遍历1次,
			所以算法的时间复杂度为O(2n)
		*/
		if record[cur-1] {
			continue
		}
		count := 1 //count表示从cur开始最长的连续序列的长度
		for j := cur + 1; record[j]; j++ {
			count++
		}
		if count > ret {
			ret = count
			/*
				1. 如果判断到某一段连续序列长度已经过半,那肯定不会存在比它更长的了,
				因为留给其他序列的位置最多也就是一半
			*/
			if ret >= len(nums)>>1 {
				return ret
			}
		}
	}

	return ret
}

/*
解法2 并查集
1. 利用并查集把相邻元素合并为一个集合
2. 在执行union时,把较小元素的根节点指向较大元素的根节点
3. 遍历并查集的每个集合,集合的根节点-集合中最小的元素就是当前集合的长度
*/
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	count := 0                       //集合的个数
	parents := map[int]int{}         //并查集父节点
	for i := 0; i < len(nums); i++ { //初始化并查集,每种元素一个集合
		parents[nums[i]] = nums[i]
		count++
	}

	//查找v所在的集合
	find := func(v int) int {
		if _, ok := parents[v]; !ok {
			return -1
		}
		//路径分裂:使路径上的每个节点都指向其祖父节点（parent的parent）
		for v != parents[v] {
			p := parents[v]
			parents[v] = parents[p]
			v = p
		}
		return v
	}

	/*
		两个集合联合在一起
		Quick Union 的 union(v1, v2)：让 v1 的根节点指向 v2 的根节点
	*/
	union := func(v1, v2 int) {
		p1 := find(v1)
		p2 := find(v2)
		if p1 == p2 {
			return
		}
		parents[p1] = p2
		count--
		return
	}

	//遍历到元素v时,如果v+1也存在,就把他们连通成一个集合
	for _, v := range parents {
		if _, ok := parents[v+1]; ok {
			union(v, v+1)
		}
	}
	ret := 0
	/*
		1. 因为是union(v,v+1),也就是把v的根节点指向v+1的根节点.所以根节点的值始终是最大的
		2. 遍历所有元素,当i为最长序列所在集合时,集合的根节点就是最长序列的最后一个元素
		所以元素个数也就是find(i) - i + 1
	*/
	for i := range parents {
		cur := find(i) - i + 1
		if cur > ret {
			ret = cur
		}
	}

	return ret
}

// @lc code=end
