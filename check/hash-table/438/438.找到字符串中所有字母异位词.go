package main

import "fmt"

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start

func findAnagrams(s string, p string) []int {
	ret := []int{}
	if len(s) < len(p) || len(p) < 1 {
		return ret
	}
	sMap := [26]int{}
	pMap := [26]int{}

	for i, v := range p {
		sMap[s[i]-'a']++
		pMap[v-'a']++
	}

	n := len(s) - len(p)
	for i := 0; i <= n; i++ {
		if sMap == pMap {
			ret = append(ret, i)
		}
		if i < n {
			sMap[s[i]-'a']--
			sMap[s[i+len(p)]-'a']++
		}
	}
	return ret
}

/*
解法1
暴力法
超时
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

// @lc code=end
