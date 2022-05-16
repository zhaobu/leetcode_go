package main

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	check := func(i int) bool {
		var last byte
		for _, v := range strs {
			if len(v) <= i {
				return false
			}
			if last == 0 {
				last = v[i]
			} else if v[i] != last {
				return false
			}
		}
		return true
	}

	i := 0
	for {
		if !check(i) {
			break
		}
		i++
	}

	return strs[0][:i]
}

// @lc code=end
