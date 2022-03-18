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
使用used来标记已经哪些元素已使用
*/
func permute2(nums []int) [][]int {
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

/*
解法3 dfs
1. 直接用nums内的元素交换实现不同排列
2. dfs表示当前正在交换下标为swapIndex的元素
3. 每一层都是从swapIndex开始一直交换到len(nums)-1
*/
func permute(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 1 {
		return ret
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
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
			dfs(swapIndex + 1)
			nums[swapIndex], nums[i] = nums[i], nums[swapIndex]
		}
	}

	dfs(0)

	return ret
}

// @lc code=end
