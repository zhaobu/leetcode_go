package main

import "sort"

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
//计数,数组作为key使用map进行分组
func groupAnagrams1(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	groupMap := map[[26]int][]string{}
	for _, str := range strs {
		arr := [26]int{}
		for _, ch := range str {
			arr[ch-'a']++
		}
		groupMap[arr] = append(groupMap[arr], str)
	}
	res := make([][]string, 0, len(groupMap))
	for _, v := range groupMap {
		res = append(res, v)
	}
	return res
}

//字符串转换为byte数组排序,string作为key使用map进行分组
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 1 {
		return [][]string{strs}
	}
	groupMap := map[string][]string{}
	for _, str := range strs {
		strByte := []byte(str)
		sort.Slice(strByte, func(i, j int) bool { return strByte[i] < strByte[j] })
		sortedStr := string(strByte)
		groupMap[sortedStr] = append(groupMap[sortedStr], str)
	}
	res := make([][]string, 0, len(groupMap))
	for _, v := range groupMap {
		res = append(res, v)
	}
	return res
}

// @lc code=end
