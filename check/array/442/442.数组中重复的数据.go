package main

import "fmt"

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start

/*
解法1 原地交换
*/
func findDuplicates1(nums []int) []int {
	ret := []int{}
	count := 0
	defer func() {
		fmt.Printf("解法1 nums=%+v,count=%d\n", nums, count)
	}()

	//把元素交换到正确的位置
	for i := range nums {
		//把nums[i]放到正确的位置上
		for nums[i] != nums[nums[i]-1] {
			count++
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

/*
解法2 原地交换另一种写法(比方法1效率更高)
和剑指 Offer 03. 数组中重复的数字思路类似
*/
func findDuplicates2(nums []int) []int {
	ret := []int{}
	// count := 0
	// defer func() {
	// 	fmt.Printf("解法2 nums=%+v,count=%d\n", nums, count)
	// }()
	//把元素交换到正确的位置
	for i := 0; i < len(nums); i++ {
		//把nums[i]放到争取的位置上
		if i+1 != nums[i] {
			// count++
			if nums[nums[i]-1] != nums[i] {
				nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
				i--
			}
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

/*
解法3 使用正负号
观察到数据范围是[1,n],所以可以考虑用负数标记数字是否已经出现过

*/
func findDuplicates(nums []int) []int {
	ret := []int{}
	//把元素交换到正确的位置

	for _, raw := range nums {
		//获取保存在下标i处的原始元素值
		if raw < 0 {
			raw = -raw
		}

		if nums[raw-1] < 0 {
			//如果这个元素以前出现过
			ret = append(ret, raw)
		} else {
			//如果这个元素之前没出现过,就标记出现过
			nums[raw-1] = -nums[raw-1]
		}
	}
	// fmt.Printf("nums=%+v\n", nums)
	return ret
}

// @lc code=end
