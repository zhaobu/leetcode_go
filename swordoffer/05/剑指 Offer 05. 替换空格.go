package main

import (
	"strings"
)

/*
 *
 *
 * 剑指 Offer 04. 二维数组中的查找
 */

// @lc code=start

/*
解法1
*/
func replaceSpace(s string) string {
	build := strings.Builder{}
	for i := range s {
		if s[i] != ' ' {
			build.WriteByte(s[i])
		} else {
			build.WriteString("%20")
		}
	}

	return build.String()
}

// @lc code=end
