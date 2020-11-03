package Solution

func permute1(nums []int) [][]int {
	track := make([]int, 0, len(nums))
	res := [][]int{}
	backtrack(nums, track, &res)
	return res
}

func backtrack(nums, track []int, res *[][]int) {
	// 触发结束条件
	if len(nums) == len(track) {
		*res = append(*res, append([]int{}, track...))
		return
	}
	for i := range nums {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		backtrack(nums, track, res)
		track = track[0 : len(track)-1]
	}
	return
}

func contains(nums []int, value int) bool {
	for _, v := range nums {
		if v == value {
			return true
		}
	}
	return false
}

/*
[leetcode 解答,优化contains](https://leetcode-cn.com/problems/permutations/solution/gojie-fa-xiang-xi-zhu-shi-hui-su-mo-ban-zou-tian-x/)
*/

func permute2(nums []int) [][]int {
	track := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	res := [][]int{}
	backtrack2(nums, track, used, &res)
	return res
}

// 回溯核心
// nums: 原始列表
// track: 路径上的数字
// used: 是否访问过
func backtrack2(nums, track []int, used []bool, res *[][]int) {
	// 结束条件：走完了，也就是路径上的数字总数等于原始列表总数
	if len(nums) == len(track) {
		// 切片底层公用数据，所以要copy,把本次结果追加到最终结果上
		*res = append(*res, append([]int{}, track...))
		return
	}

	// 开始遍历原始数组的每个数字
	for i := 0; i < len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			// 没有访问过就选择它，然后标记成已访问过的
			used[i] = true
			// 做选择：将这个数字加入到路径的尾部，这里用数组模拟链表
			track = append(track, nums[i])
			backtrack2(nums, track, used, res)
			// 撤销刚才的选择，也就是恢复操作
			track = track[:len(track)-1]
			// 标记成未使用
			used[i] = false
		}
	}
}
