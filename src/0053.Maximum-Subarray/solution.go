package Solution

/*
[leetcode 官方 方法一：动态规划)](https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/)
*/
func MaxSubArray1(nums []int) int {
	maxAns := nums[0]
	pre := 0
	for _, v := range nums {
		pre = max(pre+v, v)
		maxAns = max(maxAns, pre)
	}
	return maxAns
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxSubArray2(nums []int) int {
	maxAns := nums[0]
	sum := 0
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		maxAns = max(maxAns, sum)
	}
	return maxAns
}
