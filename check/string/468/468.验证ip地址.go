package main

/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start

/*
解法1 dfs(不推荐,容易漏掉各种情况)
*/
func validIPAddress(queryIP string) string {

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
解法2 先切割,再判断每一段
*/

// @lc code=end
