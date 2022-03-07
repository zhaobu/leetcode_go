package main

/*
 * @lc app=leetcode.cn id=673 lang=golang
 *
 * [673] 最长递增子序列的个数
 */

// @lc code=start
/*
解法1 动态规划
*/
func findNumberOfLIS1(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	// dpLen[i]表示nums[i]结尾的最长递增子序列长度
	dpLen := make([]int, len(nums))
	dpLen[0] = 1

	// dpNum[i]表示nums[i]结尾的最长递增子序列的个数
	dpNum := make([]int, len(nums))
	dpNum[0] = 1

	maxLen := 1
	maxNum := 1
	for i := 1; i < len(nums); i++ {
		dpLen[i] = 1 //默认最长递增子序列为nums[i]
		dpNum[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dpLenNew := dpLen[j] + 1
				if dpLenNew > dpLen[i] { //遇到更长的能让nums[i]拼接在后面的子序列后就更新
					dpLen[i] = dpLenNew
					dpNum[i] = dpNum[j] //每次遇到更大的都要重置
				} else if dpLenNew == dpLen[i] { //遇到和已经求得的dpLen[i]相同长度的子序列就合计一下
					dpNum[i] += dpNum[j]
				}
			}
		}

		if dpLen[i] > maxLen {
			maxLen = dpLen[i]
			maxNum = dpNum[i]
		} else if dpLen[i] == maxLen {
			maxNum += dpNum[i]
		}

	}

	return maxNum
}

/*
解法1 二分查找

*/

// @lc code=end
