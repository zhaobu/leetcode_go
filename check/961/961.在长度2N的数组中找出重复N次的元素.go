package main

/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 */

// @lc code=start

/*
解法1 数学
1. 假设重复元素为x,假设所有x之间都至少间隔2个元素
则x有n个, x之间有n-1个间隙,不同于x的数至少有2*(n-1)
2. 2*(n-1)+n = 2n+n-2. 而题目给定的元素个数是2n
当n>2时,2n+n-2 > 2n,不满足条件,所以有如下结论:
n<=2: 此时x之间可能间隔2个元素比如 [1,2,3,1]
n>2: 此时至少有2个x之间间隔小于2
*/
func repeatedNTimes1(nums []int) int {
	n := len(nums)
	//先考虑元素之间至少有2个x之间间隔<2的判断
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		} else if i > 1 && nums[i] == nums[i-2] {
			return nums[i]
		}
	}
	//再考虑当 n == 2 时类似[1,2,3,1]的特殊例子
	if n == 4 && nums[0] == nums[3] {
		return nums[0]
	}

	//不肯能到达这里
	return -1
}

/*
解法2 统计次数
题目给定范围 0 <= nums[i] <= 10^4
只有一个数出现2次,就是重复的数
*/
func repeatedNTimes(nums []int) int {
	// 直接统计出现次数
	count := [10000]byte{}
	for i := range nums {
		count[nums[i]]++
		if count[nums[i]] > 1 {
			return nums[i]
		}
	}
	return -1
}

// @lc code=end
