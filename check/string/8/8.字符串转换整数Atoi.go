package main

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start

/*
解法1
*/
func myAtoi(s string) int {
	ret := 0
	const (
		Positive = 1
		Nagivate = -1
	)
	symbol := Positive
	//处理符号部分
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == '-' {
			symbol = Nagivate
			i++
			break
		} else if s[i] == '+' {
			symbol = Positive
			i++
			break
		} else {
			break
		}
	}

	const (
		Max = 1<<31 - 1
		Min = -1 << 31
	)

	//处理数字部分
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num := int(s[i] - '0')
			if ret > Max/10 || (ret == Max/10 && num > Max%10) {
				ret = Max
				break
			} else if ret < Min/10 || (ret == Min/10 && num > -Min%10) {
				ret = Min
				break
			}
			ret = ret*10 + symbol*num
		} else {
			break
		}
	}

	return ret
}

// @lc code=end
