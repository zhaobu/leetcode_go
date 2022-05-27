package main

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start

/*
解法1 高斯求和
*/
func missingNumber1(nums []int) int {
	n := len(nums)
	ans := (0 + n) * (n + 1) / 2
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	return ans - sum
}

/*
解法2
想到异或的特点是a^a=0, a^0=a

假设n=5
0,1,2,3,4,5 全部的数字
0,1, ,3,4,5 缺失一个数字
所以只需要把每一个下标异或后的值和每个元素异或后的值再异或一次就是要求的结果
*/
func missingNumber(nums []int) int {
	/*
	   0,1,2,3,4,5
	   0   2,3,4,5
	*/
	n := len(nums)
	ret := 0
	for i := 0; i < n; i++ {
		ret ^= i ^ nums[i]
	}
	return ret ^ n
}

// @lc code=end
