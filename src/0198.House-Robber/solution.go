package Solution

/*
[小浩算法 ：动态规划)](https://www.geekxh.com/1.2.%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92%E7%B3%BB%E5%88%97/206.html#_03%E3%80%81go%E8%AF%AD%E8%A8%80%E7%A4%BA%E4%BE%8B)
*/
func Rob1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if i == 1 {
			nums[i] = max(nums[i], nums[i-1])
		} else {
			nums[i] = max(nums[i-1], nums[i-2]+nums[i])
		}
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a < b {

		return b
	}
	return a
}

/*
[labuladong 经典动态规划：打家劫舍系列问题 (带备忘录的自顶向下动态规划)](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484800&idx=1&sn=1016975b9e8df0b8f6df996a5fded0af&chksm=9bd7fb88aca0729eb2d450cca8111abd8f861236b04125ce556171cb520e298ddec4d90823b3&scene=21#wechat_redirect)
*/
func Rob2(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp := make(map[int]int, len(nums))
	return _rob2(dp, nums, 0)
}

func _rob2(dp map[int]int, nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	if v, ok := dp[start]; ok {
		return v
	}
	res := max(_rob2(dp, nums, start+1), nums[start]+_rob2(dp, nums, start+2))
	dp[start] = res
	return res
}

/*
[labuladong 经典动态规划：打家劫舍系列问题 (自底向上动态规划)](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484800&idx=1&sn=1016975b9e8df0b8f6df996a5fded0af&chksm=9bd7fb88aca0729eb2d450cca8111abd8f861236b04125ce556171cb520e298ddec4d90823b3&scene=21#wechat_redirect)
*/
func Rob3(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dp1, dp2, dpi := 0, 0, 0 //分别记录dp[i-1],dp[i-2],dp[i]
	for i := 0; i < len(nums); i++ {
		dpi = max(dp1, dp2+nums[i]) //分不偷nums[i]和偷nums[i]两种情况
		dp2 = dp1
		dp1 = dpi
	}
	return dpi
}
