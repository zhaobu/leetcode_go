package main

/*
 * @lc app=leetcode.cn id=2273 lang=golang
 *
 * [2273] 移除字母异位词后的结果数组
 */

// @lc code=start
/*
解法1 更直观写法
注意golang中数组可以直接比较,前提是2个数组类型相同且的长度相等
*/
func removeAnagrams1(words []string) []string {
	//用来计算每个字符个数
	cal := func(word string) (cnt [26]int) {
		for _, v := range word {
			cnt[v-'a']++
		}
		return cnt
	}

	ret := []string{}
	i, j := 0, 1
	for j < len(words) {
		/*
		   1. 是字母异位词,表示下标j处的单词应该删除
		   2. 删除后i保持不变,j++ 继续下一轮判断
		*/
		if cal(words[j]) == cal(words[i]) {
			j++
			continue
		}

		ret = append(ret, words[i])
		i = j
		j++
	}

	//结束时,i位置处还有一个单词是最后能保留的字符串
	ret = append(ret, words[i])

	return ret
}

/*
解法2 简洁写法
*/
func removeAnagrams(words []string) []string {
	ans := []string{words[0]}
	cmp := [26]int{}
	for _, word := range words[1:] {
		cnt := [26]int{}
		for _, b := range word {
			cnt[b-'a']++
		}
		for _, b := range ans[len(ans)-1] {
			cnt[b-'a']--
		}
		if cnt != cmp { // 不是字母异位词
			ans = append(ans, word)
		}
	}
	return ans
}

// @lc code=end
