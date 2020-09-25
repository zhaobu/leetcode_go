package Solution

func LongestCommonPrefix1(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	res := ""
	//遍历strs[0]每个字符
	for i := range strs[0] {
		ch := strs[0][i]
		for _, v := range strs {
			if i >= len(v) {
				return res
			}
			if v[i] != ch {
				return res
			}
		}

		res += string(ch)
	}

	return res
}
