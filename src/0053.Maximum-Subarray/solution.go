package Solution

/*
[leetcode 官方 方法一：动态规划)](https://leetcode-cn.com/problems/maximum-subarray/solution/zui-da-zi-xu-he-by-leetcode-solution/)
*/
func MaxSubArray1(nums []int) int {
	maxAns := nums[0] //最大和
	pre := 0
	for _, v := range nums {
		//计算以下标 i 结尾的子数组的元素的最大的和
		pre = max(pre+v, v)
		//比较之前的maxAns和本次循环计算出的最大和
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

/*
[解法二 动态规划思路二](https://leetcode.wang/leetCode-53-Maximum-Subarray.html)
*/
func MaxSubArray2(nums []int) int {
	maxAns := nums[0] //最大和
	sum := 0          //相当于dp[i-1]
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

/*
[解法一 动态规划思路一](https://leetcode.wang/leetCode-53-Maximum-Subarray.html)
*/
func MaxSubArray3(nums []int) int {
	maxAns := nums[0]
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			//dp[j]表示从下标j开始,长度为i的子数组的元素和
			dp[j] = dp[j] + nums[j+i-1]
			if dp[j] > maxAns {
				maxAns = dp[j]
			}
		}
	}
	return maxAns
}
