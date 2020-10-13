package Solution

/*
[暴力模式](https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/303.html#_02%E3%80%81sunday-%E5%8C%B9%E9%85%8D)
*/
func StrStr1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	l1 := len(haystack)
	l2 := len(needle)

	for i := 0; i <= l1-l2; i++ {
		for j := 0; j < l2; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == l2-1 {
				return i
			}
		}
	}

	return -1
}

// [sunday算法](https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/303.html#_02%E3%80%81sunday-%E5%8C%B9%E9%85%8D)
func StrStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	i := 0 //当前查询索引
	l1, l2 := len(haystack), len(needle)
	for i <= l1-l2 {
		//遍历模式串,看是否和匹配串相同
		for j := 0; j < l2; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			// 若匹配，则返回当前 idx
			if j == l2-1 {
				return i
			}
		}
		//匹配串的后一个字符不存在
		if i+l2 >= l1 {
			return -1
		}
		k := 0
		j := l2 - 1
		//遍历模式串,找到匹配串后一位字符距离模式串最后一个字符的距离
		for ; j >= 0; j-- {
			k++
			if haystack[i+l2] == needle[j] {
				break
			}
		}
		//如果未找到,直接跳过
		if j == -1 {
			i += l2 + 1
		} else {
			i += k
		}
	}
	return -1
}
