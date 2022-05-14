package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
/*
解法1 字符串分割
*/
func compareVersion1(version1 string, version2 string) int {
	v1Arr := strings.Split(version1, ".")
	v2Arr := strings.Split(version2, ".")
	n1, n2 := len(v1Arr), len(v2Arr)
	ret := 0
	for i := 0; i < len(v1Arr) || i < len(v2Arr); i++ {
		v1, v2 := 0, 0
		if i < n1 {
			v1, _ = strconv.Atoi(v1Arr[i])
		}
		if i < n2 {
			v2, _ = strconv.Atoi(v2Arr[i])
		}
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}

	return ret
}

/*
解法2 双指针
*/
func compareVersion(version1 string, version2 string) int {
	n1, n2 := len(version1), len(version2)
	ret := 0
	i, j := 0, 0
	for i < n1 || j < n2 {
		v1, v2 := 0, 0
		for ; i < n1; i++ {
			if version1[i] == '.' {
				i++
				break
			}
			v1 = v1*10 + int(version1[i]-'0')
		}

		for ; j < n2; j++ {
			if version2[j] == '.' {
				j++
				break
			}
			v2 = v2*10 + int(version2[j]-'0')
		}
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}

	return ret
}

// @lc code=end
