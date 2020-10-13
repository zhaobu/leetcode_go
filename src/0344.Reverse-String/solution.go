package Solution

/*
[小浩算法 )](https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/301.html#_01%E3%80%81%E9%A2%98%E7%9B%AE%E5%88%86%E6%9E%90)
*/
func ReverseString1(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
