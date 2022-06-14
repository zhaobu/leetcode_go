package main

import (
	"sort"
)

/*
 *
 *
 * 剑指 Offer 38. 字符串的排列
 */

// @lc code=start

/*
 解法1 回溯
*/
func permutation1(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	strArr := []byte(s)
	sort.Slice(strArr, func(i, j int) bool {
		return strArr[i] < strArr[j]
	})

	ret := []string{}
	visited := make([]bool, len(strArr))

	/*
		1. 每一个位置都能从所有字符中选一个
		2. i表示当前选到第i个字符,record记录已经选过的字符
	*/
	var dfs func(i int, record []byte)
	dfs = func(i int, record []byte) {
		if i == len(strArr) {
			ret = append(ret, string(record))
			return
		}

		for j := range strArr {
			/*
				1. 如果当前字符已经被用,则不能选
				2. 如果当前字符没有被用,则需要看他是不是和前一个字符相同,并且前一个字符已经被用.
				则第i位上就不能重复选同样的字符
			*/
			if visited[j] || (j > 0 && strArr[j] == strArr[j-1] && visited[j-1]) {
				continue
			}
			visited[j] = true
			dfs(i+1, append(record, strArr[j]))
			visited[j] = false
		}
	}

	dfs(0, make([]byte, 0, len(strArr)))

	return ret
}

/*
解法2 利用下一个排列思想
每次都找到比当前排列字典序更大的排列
*/
func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
}

// @lc code=end
