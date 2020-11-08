package Solution

func max(a, b int) int {
	if a < b {

		return b
	}
	return a
}

/*
[labuladong 经典动态规划：打家劫舍系列问题](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484800&idx=1&sn=1016975b9e8df0b8f6df996a5fded0af&chksm=9bd7fb88aca0729eb2d450cca8111abd8f861236b04125ce556171cb520e298ddec4d90823b3&scene=21#wechat_redirect)
*/
func Rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	return max(_rob1(nums, 0, len(nums)-2), _rob1(nums, 1, len(nums)-1))
}

func _rob1(nums []int, start, end int) int {
	dp1, dp2, dpi := 0, 0, 0 //分别表示dp[i-1],dp[i-2],dp[i]
	for i := start; i <= end; i++ {
		dpi = max(dp1, nums[i]+dp2)
		dp2 = dp1
		dp1 = dpi
	}
	return dpi
}
