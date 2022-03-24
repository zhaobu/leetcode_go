package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start

/*
解法1 backtrack
参考 https://labuladong.gitee.io/algo/4/30/108/
*/

/*
解法1 backtrack
*/
func subsets1(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}

	ret := [][]int{}

	var dfs func(record []int, start int)
	dfs = func(record []int, start int) {

		// 前序位置，每个节点的值都是一个子集
		ret = append(ret, append([]int{}, record...))

		for j := start; j < len(nums); j++ {
			// 做选择
			record = append(record, nums[j])
			// 通过 start 参数控制树枝的遍历，避免产生重复的子集, 次都只选后面的元素
			dfs(record, j+1)
			// 撤销选择
			record = record[:len(record)-1]
		}
	}

	dfs([]int{}, 0)

	return ret
}

/*
解法2 枚举2^n -1
官方解法1 https://leetcode-cn.com/problems/subsets/solution/zi-ji-by-leetcode-solution/
*/
func subsets(nums []int) [][]int {
	if len(nums) < 1 {
		return nil
	}

	n := len(nums)
	ret := [][]int{}
	record := make([]int, 0, n)
	m := 1 << n
	for i := 0; i < m; i++ {
		record = make([]int, 0, n)
		for j, v := range nums {
			if i&(1<<j) > 0 {
				record = append(record, v)
			}
		}
		ret = append(ret, record)
	}

	return ret
}

// @lc code=end
