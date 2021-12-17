/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	if len(digits) < 1 {
		return digits
	}
	var (
		lastOver = 1 //上一个进位
	)
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+lastOver < 10 { //如果和小于10就直接相加,并且结束
			digits[i] += lastOver
			lastOver = 0
			break
		} else {
			digits[i] = 0
			lastOver = 1
		}
	}
	if lastOver == 0 {
		return digits
	} else {
		return append([]int{lastOver}, digits...) //如果是9999这种,就在前面加1
	}
}

// @lc code=end

