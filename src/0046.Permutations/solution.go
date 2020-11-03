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
