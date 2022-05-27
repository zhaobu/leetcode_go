package main

/*
 * [面试题 17.11. 单词距离](https://leetcode.cn/problems/find-closest-lcci/)
 */

/*
解法1 用2个map存储
*/
func findClosest1(words []string, word1 string, word2 string) int {
	record := map[string][]int{}

	for i, v := range words {
		record[v] = append(record[v], i)
	}

	index1, index2 := record[word1], record[word2]

	abs := func(a, b int) int {
		if a > b {
			return a - b
		}
		return b - a
	}

	ret := len(words)
	for _, idx1 := range index1 {
		for _, idx2 := range index2 {
			cur := abs(idx1, idx2)
			if cur < ret {
				ret = cur
			}
		}
	}
	return ret
}

/*
解法2 一遍遍历
*/

func findClosest(words []string, word1 string, word2 string) int {
	idx1, idx2 := -1, -1
	ret := len(words)
	for i, v := range words {
		if v == word1 {
			idx1 = i
			if idx2 >= 0 && idx1-idx2 < ret {
				ret = idx1 - idx2
			}
		} else if v == word2 {
			idx2 = i
			if idx1 >= 0 && idx2-idx1 < ret {
				ret = idx2 - idx1
			}
		}
	}

	return ret
}
