package main

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * 剑指 Offer 50. 第一个只出现一次的字符
 */

// @lc code=start

/*
 解法1 双指针
*/
func firstUniqChar1(s string) byte {
	n := len(s)
	if n < 1 {
		return ' '
	}
	if n < 2 {
		return s[0]
	}

	cnts := [26]int{}
	for i := range s {
		cnts[s[i]-'a']++
	}

	for i := range s {
		if cnts[s[i]-'a'] == 1 {
			return s[i]
		}
	}

	return ' '
}

/*
解法2 map
*/
func firstUniqChar(s string) byte {
	n := len(s)
	if n < 1 {
		return ' '
	}
	if n < 2 {
		return s[0]
	}

	cnts := [26]int{}
	for i := range cnts {
		cnts[i] = n //初始化为n,不初始化为0是因为cnts[i]要用来表示下标,下标可能为0
	}

	for i := range s {
		if cnts[s[i]-'a'] == n { //第一次出现时记录下标
			cnts[s[i]-'a'] = i
		} else { //重复出现时就变为n+1
			cnts[s[i]-'a'] = n + 1
		}
	}
	// fmt.Printf("cnts=%+v\n", cnts)
	minIndex := n
	for _, v := range cnts {
		if v < minIndex { //过滤出只出现一次且下标最小的元素
			minIndex = v
		}
	}
	if minIndex == n {
		return ' '
	}
	return s[minIndex]
}

// @lc code=end
