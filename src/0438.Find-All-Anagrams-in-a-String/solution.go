package Solution

/*
[labuladong 滑动窗口算法 专题](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485141&idx=1&sn=0e4583ad935e76e9a3f6793792e60734&chksm=9bd7f8ddaca071cbb7570b2433290e5e2628d20473022a5517271de6d6e50783961bebc3dd3b&scene=21#wechat_redirect)
*/
func findAnagrams1(s string, p string) []int {
	need, window := make(map[byte]int, len(p)), make(map[byte]int, len(p))
	for i := range p {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	res := []int{} // 记录结果
	for right < len(s) {
		c := s[right]
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(p) {
			// 当窗口符合条件时，把起始索引加入 res
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

// [leetcode 题解](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/solution/golang-jie-jin-shuang-100-by-haa4qccppa/)
func findAnagrams2(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	//用26个小写字母替代map
	cn1 := [26]int{}
	cn2 := [26]int{}

	//初始化,分别统计s和p从0到len(p)-1的每个字符的个数
	for i := 0; i < len(p); i++ {
		cn1[p[i]-'a']++
		cn2[s[i]-'a']++
	}
	res := []int{}
	// 遍历s字符串,去掉下标为i的字符,新增下标为i+len(p)的字符
	for i := 0; i <= len(s)-len(p); i++ {
		if cn1 == cn2 {
			res = append(res, i)
		}
		if i+len(p) >= len(s) {
			break
		}
		cn2[s[i]-'a']--
		cn2[s[i+len(p)]-'a']++

	}
	return res
}
