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
func subsets(nums []int) [][]int {
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

// @lc code=end

