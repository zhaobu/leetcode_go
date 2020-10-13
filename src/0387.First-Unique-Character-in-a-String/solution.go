package Solution

/*
[小浩算法  使用hash)](https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/302.html#_02%E3%80%81%E9%A2%98%E7%9B%AE%E5%9B%BE%E8%A7%A3)
*/
func FirstUniqChar1(s string) int {
	tb := map[rune]int{}
	for _, v := range s {
		tb[v]++
	}
	for i, v := range s {
		if tb[v] == 1 {
			return i
		}
	}
	return -1
}

/*
[小浩算法  使用数组)](https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/302.html#_02%E3%80%81%E9%A2%98%E7%9B%AE%E5%9B%BE%E8%A7%A3)
*/
func FirstUniqChar2(s string) int {
	tb := [26]int{}
	for i, v := range s {
		tb[v-'a'] = i
	}
	for i, v := range s {
		if i == tb[v-'a'] {
			return i
		} else {
			tb[v-'a'] = -1
		}
	}
	return -1
}
