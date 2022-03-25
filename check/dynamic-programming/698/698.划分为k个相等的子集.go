package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start

/*
解法1 backtrack 超时
参考: https://labuladong.gitee.io/algo/4/30/107/
每个数字都要选择进入到 k 个子集中的某一个
*/
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	ret := false
	set := make([]int, k)
	sum /= k

	var isSame func() bool
	isSame = func() bool {
		for _, v := range set {
			if v != sum {
				return false
			}
		}
		return true
	}

	//降序排序更能达到剪枝条件
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	if nums[0] > sum {
		return false
	}
	fmt.Printf("nums=%+v\n", nums)
	var dfs func(start int)
	dfs = func(start int) {
		if start == len(nums) {
			if isSame() {
				ret = true
				// fmt.Printf("set=%+v,start=%d\n", set, start)
			}
			return
		}
		// 穷举 nums[start] 可能加入的子集
		for i := 0; i < k; i++ {
			// 剪枝，子集装装满了
			if set[i]+nums[start] > sum {
				continue
			}
			if ret {
				return
			}
			set[i] += nums[start]
			dfs(start + 1)
			set[i] -= nums[start]
			/*
				1. 很关键的剪枝,可以大幅提高效率
				2. 因为
			*/
			// fmt.Printf("set[i]=%+v\n", set[i])
			if set[i] == 0 {
				break
			}
		}
	}

	dfs(0)
	return ret
}

/*
解法2 动态规划
官方解法1 https://leetcode-cn.com/problems/partition-to-k-equal-sum-subsets/solution/hua-fen-wei-kge-xiang-deng-de-zi-ji-by-leetcode/
*/
func canPartitionKSubsets2(nums []int, k int) bool {
	n := len(nums)
	if n < k {
		return false
	}
	if k == 1 {
		return true
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	sort.Ints(nums)
	target := sum / k
	if nums[n-1] > target {
		return false
	}

	m := 1 << n
	dp := make([]bool, m)
	dp[0] = true
	curSum := make([]int, m)

	for i := 0; i < m; i++ {
		if !dp[i] {
			continue
		}
		for j := 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				continue
			}
			next := i | (1 << j)
			if dp[next] {
				continue
			}
			if curSum[i]%target+nums[j] <= target {
				curSum[next] = curSum[i] + nums[j]
				dp[next] = true
			} else {
				break
			}
		}
	}
	return dp[m-1]
}

/*
解法3 backtrack
对于每个子集，都要遍历 nums 中的 n 个数字，然后选择是否将当前遍历到的数字加到当前子集
*/
func canPartitionKSubsets3(nums []int, k int) bool {
	n := len(nums)
	if n < k {
		return false
	}
	if k == 1 {
		return true
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	//降序排序更能达到剪枝条件
	sort.Ints(nums)
	target := sum / k
	if nums[n-1] > target {
		return false
	}

	ret := false
	visit := 0
	dp := make(map[int]bool, k)
	var dfs func(index, curSum, start int)
	dfs = func(index, curSum, start int) {
		// 所有桶都被装满了，而且 nums 一定全部用完了
		if index == 1 {
			ret = true
			return
		}
		if ret {
			return
		}
		if start == n {
			return
		}
		// 装满了当前桶，递归穷举下一个桶的选择,下一个桶从 nums[0] 开始选数字
		if curSum == target {
			dfs(index-1, 0, 0)
			dp[visit] = true
			return
		}

		if dp[visit] {
			return
		}

		for i := start; i < n; i++ {
			if ret {
				return
			}
			//如果第 i 位是 1,nums[i] 已经被装入别的桶中
			if (visit>>i)&1 == 1 || curSum+nums[i] > target {
				continue
			}
			visit |= 1 << i
			dfs(index, curSum+nums[i], i+1)
			visit ^= 1 << i
		}
	}

	// k 号桶初始什么都没装，从 nums[0] 开始做选择
	dfs(k, 0, 0)
	return ret
}

// @lc code=end
