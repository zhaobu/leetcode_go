package main

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	m := len(nums)
	if m < 1 {
		return 1
	}

	//把不在[1,m]范围内的数变为0
	for i := 0; i < m; i++ {
		if nums[i] < 1 || nums[i] > m {
			nums[i] = m + 1
		}
	}

	/*
	   1.
	   2.
	*/
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := 0; i < m; i++ {
		index := abs(nums[i]) - 1         //当前元素真正的值用来当做下标
		if index < m && nums[index] > 0 { //说明当前下标对应的元素还未标记
			nums[index] = -nums[index] //标记当前下标对应的元素
		}
	}
	// fmt.Printf("nums=%+v\n", nums)
	//找到第一个>=0的下标
	i := 0
	for i < m && nums[i] < 0 {
		i++
	}
	if i == m {
		return m + 1
	} else {
		return i + 1
	}
}

// @lc code=end
