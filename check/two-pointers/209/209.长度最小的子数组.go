package main

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start

/*
解法1 滑动窗口
*/
func minSubArrayLen1(target int, nums []int) int {
	m := len(nums)

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	ret := m + 1
	left, right := 0, 0
	sum := 0
	for right < m {
		sum += nums[right]
		for sum >= target {
			ret = min(ret, right-left+1)
			if ret == 1 {
				return ret
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	if ret == m+1 {
		return 0
	}
	return ret
}

/*
解法2 暴力法
*/
func minSubArrayLen2(target int, nums []int) int {
	m := len(nums)

	ret := m + 1
	for i := 0; i < m; i++ {
		sum := 0
		for j := i; j < m; j++ {
			sum += nums[j]
			if sum >= target {
				cur := j - i + 1
				if cur < ret {
					if cur == 1 {
						return 1
					}
					ret = cur
				}
				break
			}
		}
	}
	if ret == m+1 {
		return 0
	}
	return ret
}

/*
解法3 前缀和+二分查找
解法2的优化
*/
func minSubArrayLen(target int, nums []int) int {
	m := len(nums)

	/*
		1. preSum[i]表示nums数组中[0,i-1]的前缀和
	*/
	preSum := make([]int, m+1)
	for i := 0; i < m; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	/*
		在[start,end]范围内二分查找第一个>=k的下标
	*/
	search := func(start, end, k int) int {
		// fmt.Printf("start=%d,end=%d,k=%d\n", start, end, k)
		for start <= end {
			mid := start + (end-start)>>1
			if preSum[mid] < k {
				start = mid + 1
			} else {
				if mid == 0 || preSum[mid-1] < k {
					return mid
				}
				end = mid - 1
			}
		}
		return -1
	}

	ret := m + 1
	for i := 0; i < m; i++ {
		/*
			1. 对于preSum的下标[i+1,m],实际上对于nums来说是求[i,m-1]范围内的元素和
		*/
		j := search(i+1, m, preSum[i]+target)
		if j == -1 { //未找到满足条件的下标就直接跳过
			continue
		}
		/*
			1. 求得的下标j表明nums数组[i,j-1]范围内的元素之和>=target
		*/
		cur := j - i
		if cur < ret {
			if cur == 1 {
				return 1
			}
			ret = cur
		}
	}
	if ret == m+1 {
		return 0
	}
	return ret
}

// @lc code=end
