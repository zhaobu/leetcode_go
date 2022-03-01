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
func longestPalindrome2(s string) string {
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

/*d
解法3
扩展中心算法
1. 下标i从0到len-1位置,每个位置向2边扩展,计算最大的回文串
2. 每个位置i都可以作为回文串的中心,能求出所有奇数长度的回文串
3. 每个位置i和右边位置i+1中间的间隙作为中心计算一次,能求出所有偶数长度的回文串
*/
func longestPalindrome3(s string) string {
	if len(s) < 2 {
		return s
	}
	start := 0  //最大回文串的开始下标
	maxLen := 0 //最大回文串的长度-1(最后再+1就可以表示长度,不要在循环里面每次都+1)

	// 从i开始向左,j开始向右找到最大回文串
	var findPalindrome = func(i, j int) {
		// fmt.Printf("开始时s=%s,i=%d,j=%d\t", s, i, j)
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		// fmt.Printf("结束时s=%s,i=%d,j=%d\n", s, i, j)
		length := j - i - 1
		if length > maxLen {
			start = i + 1
			maxLen = length
		}
		return
	}
	for i := 0; i < len(s); i++ {
		findPalindrome(i-1, i+1)
		findPalindrome(i, i+1)
	}
	return s[start : start+maxLen]
}

/*
解法4
扩展中心算法优化
核心思想: 以连续的相同字符组成的子串作为扩展中心
1. 下标 i 从 0 开始遍历到 len-1
2. 找到 i 右边第一个和 s[i] 不相等的字符,下标记为 r,
i 的左边位置 i-1 记为 l
3. 由 l 开始向左, r 开始向右扩展到不是回文串
4. 找到最长回文串后, i=r ,继续1,2,3
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start := 0  //最大回文串的开始下标
	maxLen := 0 //最大回文串的长度

	// 从i开始向左,j开始向右找到最大回文串
	var findPalindrome = func(i, j int) {
		// fmt.Printf("开始时s=%s,i=%d,j=%d\t", s, i, j)
		for i >= 0 && j < len(s) && s[i] == s[j] {
			i--
			j++
		}
		/*
			1. 拿baabad字符串举例,进循环时 i=0,j=3, 循环结束时,
			最长回文串为s[0,3], i=-1 ,j=4所以length=4-(-1)-1
			2. 拿aaaad字符串举例,进循环时 i=-1,j=4, 循环结束时,
			最长回文串为s[0,3], i=-1,j=4所以length=4-(-1)-1
		*/

		// fmt.Printf("结束时s=%s,i=%d,j=%d\n", s, i, j)
		length := j - i - 1
		if length > maxLen {
			start = i + 1
			maxLen = length
		}
		return
	}

	for i := 0; i < len(s); {
		r := i + 1
		for r < len(s) && s[r] == s[i] {
			r++
		}
		findPalindrome(i-1, r)
		i = r
	}
	return s[start : start+maxLen]
}

// @lc code=end
