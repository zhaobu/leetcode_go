package Solution

/*
[labuladong 特殊数据结构：单调栈)](https://labuladong.gitbook.io/algo/shu-ju-jie-gou-xi-lie/2.4-shou-ba-shou-she-ji-shu-ju-jie-gou/dan-tiao-zhan#dan-tiao-zhan-mo-ban)
*/
func nextGreaterElements1(nums []int) []int {
	res := make([]int, len(nums))
	stack := []int{}
	n := len(nums)
	//因为是环形的,所以数组长度翻倍求解
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}
		temp := -1
		if len(stack) != 0 {
			temp = stack[len(stack)-1]
		}
		res[i%n] = temp
		stack = append(stack, nums[i%n])
	}
	return res
}
