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

// [sunday算法优化写法](https://www.geekxh.com/1.3.%E5%AD%97%E7%AC%A6%E4%B8%B2%E7%B3%BB%E5%88%97/303.html#_02%E3%80%81sunday-%E5%8C%B9%E9%85%8D)
func StrStr3(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	i := 0 //目标串匹配索
	l1, l2 := len(haystack), len(needle)
	for i <= l1-l2 {
		k := 0
		j := l2 - 1
		exist := false
		same := true
		/*
			从后往前遍历模式串:遍历过程中
			1.查看当前匹配串是否和模式串相同,如果不同就把same=false
			2.计算匹配串后一个字符是否在模式串中,如果存在把exist=true,就用k保存从后往前的索引位置
			3:如果前2个条件都发生了,也就是exist && !same就可以退出匹配串的遍历过程
		*/
		for ; j >= 0; j-- {
			if i+l2 < l1 && haystack[i+l2] == needle[j] {
				exist = true
			}
			//如果没找到
			if !exist {
				k++
			}
			if haystack[i+j] != needle[j] {
				same = false
			}
			if exist && !same {
				break
			}
		}
		//如果匹配串和模式串相同
		if same {
			return i
		}
		if !exist { //如果匹配串后一个字符不在模式串中
			i += l2 + 1
		} else {
			i += k + 1
		}
	}
	return -1
}

func getNext(s string) []int {
	next := make([]int, len(s))
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}

// [leetcode 官方 kmp算法](https://leetcode-cn.com/problems/implement-strstr/solution/bang-ni-ba-kmpsuan-fa-xue-ge-tong-tou-ming-ming-ba/)
func StrStr4(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	l1 := len(haystack)
	l2 := len(needle)

	//计算next数组
	next := getNext(needle)

	j := 0
	for i := 0; i < l1; i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == l2 {
			return i - l2 + 1
		}
	}

	return -1
}
