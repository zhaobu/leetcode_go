package main

/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的 k-diff 数对
 */

// @lc code=start

/*
解法1 hash表
*/
func findPairs1(nums []int, k int) int {
	//过滤出不同的数字
	last := 0
	record := map[int]int{}
	repeatCnt := 0
	for _, v := range nums {
		cnt := record[v]
		if cnt == 0 {
			nums[last] = v
			last++
		} else if cnt == 1 {
			repeatCnt++
		}
		record[v] += 1
	}
	if k == 0 {
		return repeatCnt
	}

	nums = nums[:last]
	ret := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			cur := nums[j] - nums[i]
			if cur == k || -cur == k {
				ret++
			}
		}
	}
	return ret
}

/*
解法2 hash表
*/
func findPairs2(nums []int, k int) int {
	record := map[int]bool{}
	visited := map[int]bool{}

	for _, v := range nums {
		if visited[v-k] {
			record[v] = true
		}
		if visited[v+k] {
			record[v+k] = true
		}
		visited[v] = true
	}
	return len(record)
}

// @lc code=end
