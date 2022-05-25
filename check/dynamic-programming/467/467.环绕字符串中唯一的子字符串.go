package main

/*
 * @lc app=leetcode.cn id=467 lang=golang
 *
 * [467] 环绕字符串中唯一的子字符串
 */

// @lc code=start
/*
解法1
1. 符合要求的子串都是属于s的连续递增子串,不过s是个无线循环字符串,所以需要考虑首位相连
2. 找出以每个字符结尾的最长连续子串,最后求和
*/
func findSubstringInWraproundString1(p string) int {

	//判断right是不是left的下一个字符
	isNext := func(left, right byte) bool {
		return (left-'a'+1)%26 == right-'a'
	}

	//dp[i]表示以字符'a'+i结尾且在 s 中的子串的最长长度
	dp := make([]int, 26)
	dp[p[0]-'a'] = 1
	k := 1 //记录连续递增子串的长度
	for i := 1; i < len(p); i++ {
		index := int(p[i] - 'a')
		if isNext(p[i-1], p[i]) {
			k++
		} else {
			k = 1
		}
		if k > dp[index] {
			dp[index] = k
		}
	}

	ret := 0
	for i := range dp {
		ret += dp[i]
	}
	return ret
}

// @lc code=end
