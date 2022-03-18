package main

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
/*
解法1 dfs
*/
func permute1(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
	}
	var dfs func(leftNums, pickedNums []int)

	dfs = func(leftNums, pickedNums []int) {
		// fmt.Printf("1  leftNums=%+v\n", leftNums)
		if len(pickedNums) == len(nums) {
			ret = append(ret, append([]int{}, pickedNums...))
			return
		}
		// fmt.Printf("2  leftNums=%+v\n", leftNums)
		for i, v := range leftNums {
			newLeftNums := make([]int, 0, len(leftNums)-1)
			for j, _ := range leftNums {
				if j != i {
					newLeftNums = append(newLeftNums, leftNums[j])
				}
			}
			// fmt.Printf("3  leftNums=%+v, newLeftNums=%+v \n", leftNums, newLeftNums)
			dfs(newLeftNums, append(pickedNums, v))
		}
	}

	dfs(nums, make([]int, 0, len(nums)))

	return ret
}

/*
解法2 dfs
*/
func permute(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
	}

	used := make([]int, len(nums)) //用来标记那些元素已使用
	var dfs func(pickedNums []int)
	dfs = func(pickedNums []int) {
		if len(pickedNums) == len(nums) {
			ret = append(ret, append([]int{}, pickedNums...))
			return
		}
		for i, v := range nums {
			if used[i] == 1 {
				continue
			}
			used[i] = 1
			dfs(append(pickedNums, v))
			used[i] = 0
		}
	}

	dfs(make([]int, 0, len(nums)))

	return ret
}

// @lc code=end
