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
