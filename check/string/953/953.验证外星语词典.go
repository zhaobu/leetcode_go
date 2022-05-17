package main

/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	dict := [26]int{}
	for i, ch := range order {
		dict[int(ch-'a')] = i
	}

	/*
	   1. word1在单词数组words中的位置比word2更小
	   2. 如果word1和word2满足字典序就返回true
	*/
	check := func(word1, word2 string) bool {
		n1, n2 := len(word1), len(word2)
		for i := 0; i < n1 && i < n2; i++ {
			index1, index2 := word1[i]-'a', word2[i]-'a'
			// fmt.Printf("ch1=%c,ch2=%c, index1=%d,index2=%d \n", word1[i], word2[i], index1, index2)
			if dict[index1] < dict[index2] {
				return true
			} else if dict[index1] > dict[index2] {
				return false
			}
		}
		/*
		   1. 走到这里说明word1,word2在下标n前面的字符都满足字典序
		   2. 在n之后就只有一个单词有字符了,如果更长的单词word2在后面,就满足字典序,如果相反则不满足
		*/
		if n1 > n2 {
			return false
		}
		return true
	}

	for i := 1; i < len(words); i++ {
		if !check(words[i-1], words[i]) {
			return false
		}
	}

	return true
}

// @lc code=end
