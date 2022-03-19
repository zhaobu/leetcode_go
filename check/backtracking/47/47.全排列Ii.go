package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start

/*
解法1 dfs
*/
func permuteUnique1(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
	}

	used := make([]bool, len(nums)) //用来标记那些元素已使用
	var repeat func(i int) bool
	repeat = func(i int) bool {
		for j := 0; j <= i; j++ {
			if nums[j] == nums[i] && used[j] {
				return true
			}
		}
		return false
	}

	var dfs func(pickedNums []int, depth int)
	dfs = func(pickedNums []int, depth int) {
		if len(pickedNums) == len(nums) {
			ret = append(ret, append([]int{}, pickedNums...))
			return
		}
		for i, v := range nums {
			/*
				1. 只需要关心当前层pickedNums应该选什么
				2. 当前层pickedNums应该选择不重复的元素
				3. 不重复也就意味着在nums数组[0,i]范围内都没有用过
			*/
			if repeat(i) {
				// fmt.Printf("0 depth=%d,i=%+v,used=%+v,pickerNums=%+v\n\n", depth, i, used, pickedNums)
				continue
			}
			used[i] = true
			// pickedNums = append(pickedNums, v)
			// fmt.Printf("1 depth=%d,i=%+v,used=%+v,pickerNums=%+v\n\n", depth, i, used, pickedNums)
			dfs(append(pickedNums, v), depth+1)
			used[i] = false
			// pickedNums = pickedNums[:len(pickedNums)-1]

		}
	}
	// fmt.Printf("nums=%+v\n", nums)
	dfs(make([]int, 0, len(nums)), 0)

	return ret
}

/*
解法2 dfs
1. 先排序,这样相同元素就是连续的
2. 结合解法1,重复元素相邻,所以判断是否重复只需要和前面一个元素比较即可
nums[i] == nums[i-1]&&used[i-1]跳过: 表示重复元素用后面一个
nums[i] == nums[i-1]&&!used[i-1]跳过: 表示重复元素用用前面一个
*/
func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
	}
	sort.Ints(nums)
	used := make([]bool, len(nums)) //用来标记那些元素已使用
	var dfs func(pickedNums []int)
	dfs = func(pickedNums []int) {
		if len(pickedNums) == len(nums) {
			ret = append(ret, append([]int{}, pickedNums...))
			return
		}
		for i, v := range nums {
			if used[i] || (i > 0 && nums[i] == nums[i-1] && used[i-1]) {
				continue
			}
			used[i] = true
			dfs(append(pickedNums, v))
			used[i] = false
		}
	}

	dfs(make([]int, 0, len(nums)))

	return ret
}

/*
解法3
*/
func permuteUnique3(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
	}

	var repeat func(swapIndex, i int) bool
	repeat = func(swapIndex, i int) bool {
		for j := swapIndex; j < i; j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
		return false
	}
	/*
		1. 表示当前正在交换下标为swapIndex的元素
		2. 从swapIndex开始一直交换到len(nums)-1
	*/
	var dfs func(swapIndex int)
	dfs = func(swapIndex int) {
		if swapIndex == len(nums) {
			ret = append(ret, append([]int{}, nums...))
			return
		}
		for i := swapIndex; i < len(nums); i++ {
			/*
				1. 当前要确定的是下标为swapIndex的元素
				2. 只要在[swapIndex,i)区间上没出现过nums[i],都可以作为swapIndex位置的元素
			*/
			if repeat(swapIndex, i) {
				continue
			}
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
			dfs(swapIndex + 1)
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
		}
	}

	dfs(0)

	return ret
}

// @lc code=end
