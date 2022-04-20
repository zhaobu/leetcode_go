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
解法1 并查集

*/
func longestConsecutive(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	return 0
}

// @lc code=end
