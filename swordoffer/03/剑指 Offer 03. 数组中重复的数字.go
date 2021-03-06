package main

import "fmt"

/*
 *
 *
 * 剑指 Offer 03. 数组中重复的数字
 */

// @lc code=start

/*
 解法1
*/
func findRepeatNumber1(nums []int) int {
	n := len(nums)
	cnts := make([]int, n)
	for _, v := range nums {
		cnts[v]++
		if cnts[v] == 2 {
			return v
		}
	}

	//不会走到这里
	return -1
}

/*
解法2 原地交换
[0,1,2,3,3,5]
*/
func findRepeatNumber2(nums []int) int {
	count := 0
	defer func() {
		fmt.Printf("解法2 nums=%+v,count=%d\n", nums, count)
	}()
	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			count++
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			i-- //注意交换后还要判断一次当前位置
		}
	}

	return -1
}

/*
解法3 原地交换另一种写法
[0,1,2,3,3,5]
*/
func findRepeatNumber(nums []int) int {
	count := 0
	defer func() {
		fmt.Printf("解法3 nums=%+v,count=%d\n", nums, count)
	}()
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			count++
			if nums[nums[i]] == nums[i] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return -1
}

// @lc code=end
