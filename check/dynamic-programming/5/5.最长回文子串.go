package main

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start

/*
暴力法
*/
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	start := 0
	maxLen := 0

	var isPalindrome = func(start, end int) bool {
		for ; start < end; start, end = start+1, end-1 {
			if s[start] != s[end] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if isPalindrome(i, j) && j-i > maxLen {
				start = i
				maxLen = j - i
			}
		}
	}
	return s[start : start+maxLen+1]
}

/*
解法2
动态规划
dp定义: dp[i][j]表示字符串s[i,j]是否为回文串
dp初始态:
dp状态转移方程:
当字符串长度<=3也就是 j-i<=2时 dp[i][j] = s[i]==s[j]
当字符串长度>3也就是j-i>2时 dp[i][j] = s[i]==s[j] && dp[i+1][j-1]

*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// dp[i][j]表示字符串s[i,j]是否为回文串
	dp := make([][]bool, len(s))

	//初始化列
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}

	/*
		  4  e           *
		  3  d        *  *
		  2  c     *  *  *
		  1  b  *  *  *  *
		  0  a  b  c  d  e
			 0	1  2  3  4
	*/
	start := 0  //最大回文串的开始下标
	maxLen := 0 //最大回文串的长度-1(最后再+1就可以表示长度,不要在循环里面每次都+1)
	// i要从大到小
	for i := len(s) - 1; i >= 0; i-- {
		// j要从小到大
		for j := i; j < len(s); j++ {
			length := j - i
			// 可以写成一行 dp[i][j] = s[i] == s[j] && (length <= 2 || dp[i+1][j-1])
			if s[i] == s[j] {
				if length <= 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && length > maxLen {
				maxLen = length
				start = i
			}
		}
	}
	// for i, v := range dp {
	// 	fmt.Printf("dp[%d]=%+v\n", i, v)
	// }
	return s[start : start+maxLen+1]
}

// @lc code=end
