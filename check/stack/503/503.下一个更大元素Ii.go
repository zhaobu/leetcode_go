package main

/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
/*
解法1 暴力法
*/
func nextGreaterElements1(nums []int) []int {
	n := len(nums)

	ret := make([]int, n)
	for i := 0; i < len(nums); i++ {
		ret[i] = -1
		for j := (i + 1) % n; j != i; j = (j + 1) % n {
			if nums[j] > nums[i] {
				ret[i] = nums[j]
				break
			}
		}
	}
	return ret
}

/*
解法2 单调栈
*/

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ret := make([]int, n)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		ret[i] = -1
		for j := i % n; j != i; j = (j + 1) % n {
			if len(stack) < 1 || nums[j] < nums[stack[len(stack)-1]] {
				stack = append(stack, j)
			} else {
				for len(stack) > 0 && nums[j] >= nums[stack[len(stack)-1]] {
					top := stack[len(stack)-1]
					ret[top] = nums[j]
					stack = stack[:len(stack)-1]
				}
			}
			break
		}

	}
	return ret
}

// @lc code=end
