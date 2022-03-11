package main

/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start

/*
解法1 动态规划
定义参考官方解法: https://leetcode-cn.com/problems/wiggle-subsequence/solution/bai-dong-xu-lie-by-leetcode-solution-yh2m/

上升摆动序列: 最后一个元素呈上升趋势。如序列 [1,3,2,4][1,3,2,4] 即为「上升摆动序列」
下降摆动序列: 最后一个元素呈下降趋势。如序列 [4,2,3,1][4,2,3,1] 即为「下降摆动序列」
长度为 1 的序列，它既是「上升摆动序列」，也是「下降摆动序列」
峰: 当且仅当该元素两侧的相邻元素均小于它。如序列 [1,3,2,4][1,3,2,4] 中，3 就是一个「峰」
谷: 当且仅当该元素两侧的相邻元素均大于它。如序列 [1,3,2,4][1,3,2,4] 中，2 就是一个「谷」
对于位于序列两端的元素，只有一侧的相邻元素小于或大于它，我们也称其为「峰」或「谷」。如序列 [1,3,2,4][1,3,2,4] 中，
1 也是一个「谷」，4 也是一个「峰」


状态定义:
1. up[i] 表示以前 i 个元素中的某一个为结尾的最长的「上升摆动序列」的长度。
2. down[i] 表示以前 i 个元素中的某一个为结尾的最长的「下降摆动序列」的长度。
所以最终要求的结果是max(up[len-1], down[len-1])

basecase:
up[0] = 1
down[0] = 1

转移方程:
1. 如果nums[i]==nums[i-1],假如i-1时用到了nums[i-1],新来nums[i]持平，既没法上升也没法下降
假如之前没用到nums[i-1],新来的也用不到
up[i]=up[i-1], down[i]=down[i-1]

2. 如果nums[i] > nums[i-1]
up[i]: 转移来源只会是up[i-1]或者down[i-1]
考虑从up[i-1]转移: 因为up[i-1]已经是[0,i-1]之间最长上升摆动序列,所以再加一个字符nums[i]也不能再继续上升
则 up[i] = up[i-1]
考虑从down[i-1]转移: 假设down[i-1]最后一个元素是 num[j], 0<=j<=i-1 则有如下情况:
情况1: nums[j] <= nums[i] > nums[i-1] .因为nums[j]是波谷,所以nums[i]一定可以加在后面变成上升摆动序列
up[i] = down[i-1] + 1
情况2: nums[j] > nums[i] > nums[i-1]  则可以让down[i-1]的最后一个元素从nums[j]替换为
nums[i-1],然后加上nums[i]也可以变成上升摆动序列.
up[i] = down[i-1] + 1
综合来看 up[i] = max(up[i-1],down[i-1] + 1)

down[i]: 转移来源只会是 up[i-1] 或者 down[i-1]
考虑从 down[i-1] 转移: 因为 down[i-1] 已经是 [0,i-1] 之间最长下降摆动序列,所以再加一个字符 nums[i] 也不能再继续下降
则 down[i] = down[i-1]
考虑从 up[i-1] 转移: 假设 up[i-1] 最后一个元素是 num[j],0<=j<=i-1 则有如下情况:
情况1: nums[j] <= nums[i] > nums[i-1]. 从nums[j]到nums[i] 无法下降,此种情况不可能转移
情况2: nums[j] > nums[i] > nums[i-1]. nums[j]是up[i-1]的最后一个元素,但nums[j] > nums[i-1]
说明摆动序列可以更长,此种情况会被包含在down[i-1]里面
则 down[i] = down[i-1]
综合来看 down[i] = down[i-1]

*/
func wiggleMaxLength1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	up := make([]int, n)
	down := make([]int, n)

	up[0] = 1
	down[0] = 1
	maxLen := 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			up[i] = up[i-1]
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(down[i-1], up[i-1]+1)
		} else {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		}

		maxLen = max(maxLen, max(up[i], down[i]))
	}
	return maxLen
}

/*
解法2 解法1状态压缩
根据状态转移方程优化
*/
func wiggleMaxLength2(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	up := 1
	down := 1
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			down = max(down, up+1)
		} else if nums[i] > nums[i-1] {
			up = max(up, down+1)
		}

	}
	return max(up, down)
}

/*
解法3 贪心
求最长摆动子序列其实就是求波峰加上波谷的和
*/

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	// res := []int{nums[0]}
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans = 2
		// res = append(res, nums[1])
	}
	for i := 2; i < n; i++ {
		diff := nums[i] - nums[i-1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
			// res = append(res, nums[i])
		}
	}
	// fmt.Printf("nums=%+v\n,res=%+v\n", nums, res)
	return ans
}

// @lc code=end
