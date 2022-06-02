package main

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start

/*
解法1 原地交换
*/
func findDuplicates(nums []int) []int {
	ret := []int{}

	//把元素交换到正确的位置
	for i := range nums {
		//把nums[i]放到争取的位置上
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			ret = append(ret, nums[i])
		}
	}
	// fmt.Printf("nums=%+v\n", nums)
	return ret
}

// @lc code=end
