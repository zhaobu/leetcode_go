package main /*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start
func rand10() int {
	first, second := 0, 0
	/*
		第一个数取[1,6]
	*/
	for {
		first = rand7()
		if first <= 6 {
			break
		}
	}

	/*
		第二个数取[1,5]
	*/
	for {
		second = rand7()
		if second <= 5 {
			break
		}
	}
	//然后再判断first的奇偶
	if first&1 == 0 {
		return second
	} else {
		return 5 + second
	}
}

// @lc code=end
