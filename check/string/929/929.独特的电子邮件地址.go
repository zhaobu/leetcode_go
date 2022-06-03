package main

import "strings"

/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] 独特的电子邮件地址
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	record := map[string]bool{}

	decode := func(str string) (raw string) {
		data := strings.Split(str, "@")
		build := strings.Builder{}
		//处理本地名
		for _, ch := range data[0] {
			if ch == '+' {
				break
			}
			if ch != '.' {
				build.WriteRune(ch)
			}
		}
		build.WriteString("@")
		build.WriteString(data[1])
		return build.String()
	}

	for _, v := range emails {
		record[decode(v)] = true
	}

	return len(record)
}

// @lc code=end
