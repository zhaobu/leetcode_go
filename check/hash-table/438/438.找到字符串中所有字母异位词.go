package main

import "fmt"

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start

/*
解法1 暴力法(超时)
*/
func findAnagrams1(s string, p string) []int {
	ret := []int{}
	if len(s) < len(p) || len(p) < 1 {
		return ret
	}
	byteS := []byte(s)
	byteP := []byte(p)

	var isAnagrams = func(a, b []byte) bool {
		fmt.Printf("a=%s,b=%s\n", string(a), string(b))
		aMap := [26]int{}

		for _, v := range a {
			aMap[v-'a']++
		}

		for _, v := range b {
			idx := v - 'a'
			if aMap[idx] < 1 {
				return false
			}
			aMap[idx]--
		}
		for _, v := range aMap {
			if v != 0 {
				return false
			}
		}
		return true
	}

	n := len(byteS) - len(byteP) + 1
	for i := 0; i < n; i++ {
		if isAnagrams(byteS[i:i+len(byteP)], byteP) {
			ret = append(ret, i)
		}
	}

	return ret
}

/*
解法2 滑动窗口
*/
func findAnagrams(s string, p string) []int {
	//统计p字符串中字符个数
	cnts := [26]int{}
	for _, ch := range p {
		cnts[ch-'a']++
	}

	ret := []int{}
	left, right := 0, 0
	curCnts := [26]int{} //记录当前窗口[left,right]内字符个数
	for right < len(s) {
		chIndex := s[right] - 'a'
		curCnts[chIndex]++ //窗口向右滑动一格

		//判断左边界是否需要收缩
		for left <= right && curCnts[chIndex] > cnts[chIndex] {
			curCnts[s[left]-'a']--
			left++
		}

		if curCnts == cnts {
			ret = append(ret, left)
		}
		right++
	}

	return ret
}

// @lc code=end
