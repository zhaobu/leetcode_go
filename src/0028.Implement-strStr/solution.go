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

func StrStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	return -1
}
