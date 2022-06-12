package main

/*
 * @lc app=leetcode.cn id=890 lang=golang
 *
 * [890] 查找和替换模式
 */

// @lc code=start

/*
解法1 构造映射
*/
func findAndReplacePattern(words []string, pattern string) []string {

	getCount := func(str string) (count int) {
		dict := [26]byte{}
		for _, ch := range str {
			if dict[ch-'a'] == 0 {
				count++
			}
			dict[ch-'a']++
		}
		return count
	}

	//计算pattern字符串中字符的种类
	charCount := getCount(pattern)

	check := func(word string) bool {
		if len(word) != len(pattern) {
			return false
		}
		/*
			dict[i]表示word中字对字符i对应pattern的的字符
		*/
		dict := [26]byte{}
		for i := range dict {
			dict[i] = 26
		}
		curCount := 0
		for i, v := range word {
			if dict[v-'a'] == 26 {
				//如果字符第一次出现就进行映射并计数
				curCount++
				dict[v-'a'] = pattern[i]
			} else if dict[v-'a'] != pattern[i] {
				return false
			}
		}

		return curCount == charCount
	}

	ret := []string{}
	for _, str := range words {
		if check(str) {
			ret = append(ret, str)
		}
	}

	return ret
}

// @lc code=end
