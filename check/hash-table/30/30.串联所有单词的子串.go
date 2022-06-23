package main

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start

/*
解法1 滑动窗口
是 438. 找到字符串中所有字母异位词（https://leetcode.cn/problems/find-all-anagrams-in-a-string/）题的
进阶版
*/
func findSubstring(s string, words []string) []int {
	sLen, n := len(s), len(words[0])
	if sLen < n {
		return nil
	}

	//words中每个单词的数量
	cnts := map[string]int{}
	for _, word := range words {
		cnts[word]++
	}

	//判断2个map是内容完全相同
	check := func(cnt1, cnt2 map[string]int) bool {
		if len(cnt1) != len(cnt2) {
			return false
		}
		for i := range cnt1 {
			if cnt2[i] != cnt1[i] {
				return false
			}
		}
		return true
	}

	ret := []int{}

	for start := 0; start < n; start++ {
		/*
			1. 可以依从从字符串的[0,n-1]开始进行窗口滑动
			也就是前面的字符都不要
			2. 每次窗口滑动都要重新统计窗口内的单词个数
		*/
		curCnt := map[string]int{}
		left, right := start, start
		end := sLen - n
		for right <= end {
			rightWord := s[right : right+n]
			curCnt[rightWord]++
			right += n
			for left <= right && curCnt[rightWord] > cnts[rightWord] {
				leftWord := s[left : left+n]
				curCnt[leftWord]--
				if curCnt[leftWord] == 0 {
					delete(curCnt, leftWord)
				}
				left += n
			}
			if check(curCnt, cnts) {
				// fmt.Printf("cnts=%+v, curCnts=%+v\n", cnts, curCnt)
				ret = append(ret, left)
			}
		}
	}

	return ret
}

// @lc code=end
