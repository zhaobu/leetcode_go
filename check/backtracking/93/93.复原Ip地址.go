package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 {
		return nil
	}
	if n > 12 {
		return nil
	}
	ret := []string{}

	var dfs func(record []byte, start, i int)
	dfs = func(record []byte, start, i int) {
		if i == 3 {
			if start < n {
				if s[start] != '0' || start == n-1 {
					num, _ := strconv.Atoi(s[start:])
					if num <= 255 {
						record = append(record, '.')
						record = append(record, s[start:]...)
						ret = append(ret, string(record))
					}
				}
			}
			return
		}
		// fmt.Printf("record=%+v\n", string(record))
		for j := 1; j <= 3 && j+start <= n; j++ {
			num, _ := strconv.Atoi(s[start : start+j])
			if num > 255 {
				continue
			}
			if s[start] == '0' && j > 1 {
				continue
			}
			// fmt.Printf("num=%+v\n", num)
			if len(record) > 0 && record[len(record)-1] != '.' {
				record = append(record, '.')
			}
			record = append(record, s[start:start+j]...)
			dfs(record, start+j, i+1)
			record = record[:len(record)-j]
		}
	}
	dfs(make([]byte, 0, n), 0, 0)

	return ret
}

// @lc code=end
