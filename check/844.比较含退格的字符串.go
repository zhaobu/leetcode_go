/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	stackS := make([]rune, 0, len(s))
	stackT := make([]rune, 0, len(t))

	for _, v := range s {
		if v != '#' {
			stackS = append(stackS, v)
		} else if len(stackS) > 0 {
			stackS = stackS[0 : len(stackS)-1]
		}
	}

	for _, v := range t {
		if v != '#' {
			stackT = append(stackT, v)
		} else if len(stackT) > 0 {
			stackT = stackT[0 : len(stackT)-1]
		}
	}

	return string(stackS) == string(stackT)
}

//TODO双指针解法
// @lc code=end

