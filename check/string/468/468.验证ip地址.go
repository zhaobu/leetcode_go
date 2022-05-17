package main

import "strings"

/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start

/*
解法1 dfs(不推荐,容易漏掉各种情况)
*/
func validIPAddress1(queryIP string) string {

	/*
		验证ipv4地址的合法字符
	*/
	ipv4Char := func(ch byte) bool {
		if ch == '.' {
			return true
		}
		if ch >= '0' && ch <= '9' {
			return true
		}
		return false
	}

	/*
		1. 验证是否合法ipv4地址
		2. index表示字符串下标,count表示当前在ip地址的第几段,preSum用来记录当前ip段的数值累积
	*/
	var checkIPv4 func(index, count, preSum int) bool
	checkIPv4 = func(index, count, preSum int) bool {
		if index >= len(queryIP) {
			if count != 3 { //遍历结束时,必须满足有4段
				return false
			}
			return true
		}
		ch := queryIP[index]
		if !ipv4Char(ch) {
			return false
		}
		if ch == '.' {
			if index == 0 || index == len(queryIP)-1 {
				return false
			}
			if queryIP[index-1] == '.' {
				return false
			}
			count++
			if count > 3 {
				return false
			}
			preSum = 0
		} else {
			if ch == '0' && preSum == 0 && index+1 < len(queryIP) && queryIP[index+1] != '.' {
				return false
			}
			preSum = preSum*10 + int(ch-'0')
			if preSum > 255 {
				return false
			}
		}
		return checkIPv4(index+1, count, preSum)
	}

	if checkIPv4(0, 0, 0) {
		return "IPv4"
	}

	ipv6Char := func(ch byte) bool {
		if ch == ':' {
			return true
		}
		if ch >= '0' && ch <= '9' {
			return true
		}
		if (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F') {
			return true
		}
		return false
	}

	var checkIPv6 func(index, count, preNumCount int) bool
	checkIPv6 = func(index, count, preNumCount int) bool {
		if index >= len(queryIP) {
			if count != 7 {
				return false
			}
			return true
		}
		ch := queryIP[index]
		if !ipv6Char(ch) {
			return false
		}
		if ch == ':' {
			if index == 0 || index == len(queryIP)-1 {
				return false
			}
			if queryIP[index-1] == ':' {
				return false
			}
			count++
			if count > 7 {
				return false
			}
			preNumCount = 0
		} else {
			preNumCount++
			if preNumCount > 4 {
				return false
			}
		}
		return checkIPv6(index+1, count, preNumCount)
	}

	if checkIPv6(0, 0, 0) {
		return "IPv6"
	}

	return "Neither"
}

/*
解法2 先分割,再判断每一段
*/
func validIPAddress(queryIP string) string {

	/*
		1. 验证是否合法ipv4地址
	*/
	checkIPv4 := func(str string) bool {
		if len(str) == 0 || len(str) > 3 { //最多3位数255
			return false
		}
		num := 0
		for i, ch := range str {
			if ch < '0' || ch > '9' {
				return false
			}
			if ch == '0' && i == 0 && len(str) > 1 {
				return false
			}
			num = num*10 + int(ch-'0')
		}
		if num > 255 {
			return false
		}
		return true
	}

	checkIPv6 := func(str string) bool {
		if len(str) == 0 || len(str) > 4 { //最多4位数
			return false
		}
		for _, ch := range str {
			if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')) {
				return false
			}
		}
		return true
	}

	//先尝试判断是个ipv4
	strs := strings.Split(queryIP, ".")
	if len(strs) == 4 {
		valid := true
		for i := range strs {
			if !checkIPv4(strs[i]) {
				valid = false
				break
			}
		}
		if valid {
			return "IPv4"
		}
	}

	//再尝试判断是个ipv6
	strs = strings.Split(queryIP, ":")
	if len(strs) == 8 {
		valid := true
		for i := range strs {
			if !checkIPv6(strs[i]) {
				valid = false
				break
			}
		}
		if valid {
			return "IPv6"
		}
	}

	return "Neither"
}

// @lc code=end
