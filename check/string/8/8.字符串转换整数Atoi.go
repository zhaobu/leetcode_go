package main

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start

/*
解法1 直接转换
*/
func myAtoi1(s string) int {
	ret := 0
	const (
		Positive = 1
		Nagivate = -1
	)
	sign := Positive
	//处理符号部分
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else if s[i] == '-' {
			sign = Nagivate
			i++
			break
		} else if s[i] == '+' {
			sign = Positive
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
			ret = ret*10 + sign*num
		} else {
			break
		}
	}

	return ret
}

/*
解法2 状态机
*/
func myAtoi(s string) int {
	ret := 0
	const (
		Positive = 1
		Nagivate = -1
	)
	const (
		Start    = 0
		Signed   = 1
		InNumber = 2
		End      = 3
	)
	stateTable := [][]int{
		{Start, Signed, InNumber, End},
		{End, End, InNumber, End},
		{End, End, InNumber, End},
		{End, End, End, End},
	}

	sign := Positive
	state := Start
	getCharIndex := func(char rune) int {
		if char == ' ' {
			return 0
		} else if char == '+' || char == '-' {
			return 1
		} else if char >= '0' && char <= '9' {
			return 2
		}
		return 3
	}
	const (
		Max = 1<<31 - 1
		Min = -1 << 31
	)

	automaton := func(char rune) {
		state = stateTable[state][getCharIndex(char)]
		if state == InNumber {
			num := int(char - '0')
			if ret > Max/10 || (ret == Max/10 && num > Max%10) {
				ret = Max
			} else if ret < Min/10 || (ret == Min/10 && num > -Min%10) {
				ret = Min
			} else {
				ret = ret*10 + sign*num
			}
		} else if state == Signed {
			if char == '-' {
				sign = Nagivate
			}
		}
	}
	for _, v := range s {
		automaton(v)
	}

	return ret
}

// @lc code=end
