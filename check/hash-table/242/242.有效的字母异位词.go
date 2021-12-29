package main

import "sort"

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start

// map 通用做法
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int, len(s))
	tMap := make(map[rune]int, len(t))
	for _, ch := range s {
		sMap[ch] += 1
	}
	for _, ch := range t {
		if _, ok := sMap[ch]; !ok {
			return false
		}
		tMap[ch] += 1
	}
	for ch, n := range sMap {
		if tMap[ch] != n {
			return false
		}
	}
	return true
}

//map 使用大小为26的数组代替map
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make([]int, 26)
	//统计s中每个字母的个数
	for _, ch := range s {
		sMap[ch-'a']++
	}

	//用s中每个字符出现的个数减去t中每个字符出现的个数
	for _, ch := range t {
		if sMap[ch-'a'] < 0 { //<0就已经表示s中个数较少
			return false
		}
		sMap[ch-'a']--
	}

	//如果是异味字符串,应该最后每个字符个数都为0
	for _, n := range sMap {
		if n != 0 {
			return false
		}
	}
	return true
}

//排序后
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sByte := []byte(s)
	tByte := []byte(t)
	sort.Slice(sByte, func(i, j int) bool { return sByte[i] < sByte[j] })
	sort.Slice(tByte, func(i, j int) bool { return tByte[i] < tByte[j] })
	return string(sByte) == string(tByte)
}

// @lc code=end
