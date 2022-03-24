package main

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start

/*
解法1 backtrack
参考 https://labuladong.gitee.io/algo/4/30/108/
*/
func combine(n int, k int) [][]int {
	if n < 1 || k < 1 {
		return nil
	}

	ret := [][]int{}

	var dfs func(record []int, start int)
	dfs = func(record []int, start int) {
		if len(record) == k {
			ret = append(ret, append([]int{}, record...))
			return
		}
		//如果已经选到的元素+剩余可选的元素数量不足k个,不用再查找
		if len(record)+(n-start+1) < k {
			return
		}
		for i := start; i <= n; i++ {
			record = append(record, i)
			dfs(record, i+1)
			record = record[:len(record)-1]
		}
	}

	dfs([]int{}, 1)

	return ret
}

// @lc code=end
