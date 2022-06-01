package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start

/*
类似题目: https://www.lintcode.com/problem/386/
和leetcode 340题一样
*/
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	// write your code here
	n := len(s)
	if n < 1 || k < 1 {
		return 0
	}
	if n < k {
		return n
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	cnts := map[byte]int{}
	ret := 0
	left, right := 0, 0

	for right < n {
		cnts[s[right]]++
		for left <= right && len(cnts) > k {
			cnts[s[left]]--
			if cnts[s[left]] == 0 {
				delete(cnts, s[left])
			}
			left++
		}
		right++
		/*
			1. 运行到这里时,left,right之间的数据必定满足条件
			2. 因为每个right,都会判断一次,所以必定能记录到最大长度
		*/
		ret = max(ret, right-left)
	}

	return ret
}

/*
解法1 滑动窗口
*/
func longestSubstring1(s string, k int) int {
	n := len(s)
	if n < k {
		return 0
	}
	if k == 1 {
		return n
	}
	//先统计字符出现的个数
	cnts := [26]int{}
	totalKind := 0
	for i := range s {
		cnts[s[i]-'a']++
		if cnts[s[i]-'a'] == 1 {
			totalKind++
		}
	}

	ret := 0
	for maxKind := 1; maxKind <= totalKind; maxKind++ {
		cnts = [26]int{} //记录每种字符数量
		charKind, kind := 0, 0
		left, right := 0, 0

		for right < n {
			index := s[right] - 'a'
			cnts[index]++
			if cnts[index] == 1 { //字符种类+1
				charKind++
			}
			if cnts[index] == k { //满足出现次数>=k的字符种类+1
				kind++
			}

			//如果字符种类超过最大字符种类限制,需要减小窗口
			for charKind > maxKind {
				index = s[left] - 'a'
				if cnts[index] == 1 {
					charKind--
				}
				if cnts[index] == k {
					kind--
				}
				cnts[index]--
				left++
			}

			right++
			/*
				1. 运行到这里字符种类都是满足<=maxKind的
				2. 如果所有字符出现次数>=k,就更新记录
			*/
			if charKind == kind {
				if right-left > ret {
					ret = right - left
				}
			}
		}

	}

	return ret
}

/*
解法2 分治
*/
func longestSubstring(s string, k int) (ret int) {
	n := len(s)
	if n < k {
		return 0
	}
	if k == 1 {
		return n
	}

	//统计字符串s中每个字符的频率
	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}

	//找出一个分割字符,这个字符的频率<k
	split := byte(0)
	for i, cnt := range freq {
		if cnt > 0 && cnt < k {
			split = byte(i) + 'a'
			break
		}
	}
	//如果不存在频率<k的字符,说明整个字符串都符合要求
	if split == 0 {
		return len(s)
	}
	//继续去每一段分割后的字符串中查找
	for _, str := range strings.Split(s, string(split)) {
		tmp := longestSubstring(str, k)
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

// @lc code=end
