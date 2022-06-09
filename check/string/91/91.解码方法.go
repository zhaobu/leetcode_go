package main

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
/*
解法1 回溯(超时)
*/
func numDecodings1(s string) int {
	ret := 0
	if s[0] == '0' {
		return ret
	}

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			ret++
			return
		}
		if s[i] != '1' && s[i] != '2' {
			return
		}

		for j := i; j < len(s); j++ {
			//1位数字
			dfs(i + 1)
			//2位数字
			if j+1 < len(s) {
				dfs(i + 2)
			}
		}
	}

	//找到第一个符合条件的数字
	for i := range s {
		if s[i] == '1' || s[i] == '2' {
			dfs(i)
			break
		}
	}

	return ret
}

/*
解法2 动态规划
*/
func numDecodings2(s string) int {
	ret := 0
	if s[0] == '0' {
		return ret
	}
	if len(s) == 1 {
		return 1
	}
	/*
	   dp[i]表示以s[i]结尾的字符串的解码方法总数
	   2:       2
	   22:     2,2; 22
	   226     2,2,6; 22,6; 2,26
	   2262    2,2,6,2; 22,6,2; 2,26,2
	   2212    2,2,1,2; 22,1,2; 2,21,2; 2,2,12; 22,12
	   2210    2,2,10; 22,10;
	*/
	check := func(s string) bool {
		if s[0] == '1' {
			return true
		} else if s[0] == '2' && s[1]-'0' <= 6 {
			return true
		}
		return false
	}
	dp := make([]int, len(s))
	dp[0] = 1

	for i := 1; i < len(s); i++ {
		if s[i] != '0' { //s[i]单独为1个数字,有dp[i-1]种拆分方式
			dp[i] += dp[i-1]
		}
		if check(s[i-1 : i+1]) {
			if i > 1 {
				dp[i] += dp[i-2] //s[i]和s[i-1]组合在一起,有dp[i-2]种方式
			} else {
				dp[i] += 1
			}
		}

		if dp[i] == 0 {
			//如果中途又某个字符不能单独解码也不能和前一个字符一起解码,整个字符串都不能解析
			return 0
		}
	}

	return dp[len(s)-1]
}

/*
解法3 动态规划(空间优化)
*/
func numDecodings(s string) int {
	ret := 0
	if s[0] == '0' {
		return ret
	}
	if len(s) == 1 {
		return 1
	}
	/*
	   dp[i]表示以s[i]结尾的字符串的解码方法总数
	   2:       2
	   22:     2,2; 22
	   226     2,2,6; 22,6; 2,26
	   2262    2,2,6,2; 22,6,2; 2,26,2
	   2212    2,2,1,2; 22,1,2; 2,21,2; 2,2,12; 22,12
	   2210    2,2,10; 22,10;
	*/
	check := func(s string) bool {
		if s[0] == '1' {
			return true
		} else if s[0] == '2' && s[1]-'0' <= 6 {
			return true
		}
		return false
	}

	dp1, dp2, dp := 1, 1, 0
	for i := 1; i < len(s); i++ {
		dp = 0
		if s[i] != '0' { //s[i]单独为1个数字,有dp[i-1]种拆分方式
			dp += dp1
		}
		if check(s[i-1 : i+1]) {
			dp += dp2 //s[i]和s[i-1]组合在一起,有dp[i-2]种方式
		}

		if dp == 0 {
			//如果中途又某个字符不能单独解码也不能和前一个字符一起解码,整个字符串都不能解析
			return 0
		}
		dp2, dp1 = dp1, dp
	}
	return dp
}

// @lc code=end
