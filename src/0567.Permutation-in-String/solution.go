package Solution

/*
[滑动窗口模板解法 ](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484504&idx=1&sn=5ecbab87e42033cc0a62b635cc436977&chksm=9bd7fa50aca07346a3ffa6be6fccc445968c162af9532fa9c6304eaab2e3a1b79a4bbe758c0a&scene=21#wechat_redirect)
*/
func checkInclusion1(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	need, windows := make(map[byte]int, len(s1)), make(map[byte]int, len(s1))
	for i := range s1 {
		need[s1[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if windows[d] == need[d] {
					if valid == len(need) && right-left+1 == len(s1) {
						return true
					}
					valid--
				}
				windows[d]--
			}
		}
	}
	return false
}

/*
[leeetcode 滑动窗口解答]https://leetcode-cn.com/problems/permutation-in-string/solution/leetcode-567-zi-fu-chuan-de-pai-lie-by-eingate/
*/
func checkInclusion2(s1 string, s2 string) bool {
	//判断输入合理性以及s1若比s2还长，返回false
	if len(s1) > len(s2) {
		return false
	}
	needs := [26]int{}
	for i := range s1 {
		needs[s1[i]-'a']++
	}
	//left指针从s2最左端直到窗口右端碰到s2最右边
	for left := 0; left <= len(s2)-len(s1); left++ {
		windows := [26]int{}
		//将窗口覆盖元素加入
		for i := left; i < left+len(s1); i++ {
			windows[s2[i]-'a']++
		}
		//判断两个数组内容相等
		if needs == windows {
			return true
		}
	}
	return false
}

/*
[leeetcode 滑动窗口解答优化]https://leetcode-cn.com/problems/permutation-in-string/solution/leetcode-567-zi-fu-chuan-de-pai-lie-by-eingate/
*/
func checkInclusion3(s1 string, s2 string) bool {
	//判断输入合理性以及s1若比s2还长，返回false
	if len(s1) > len(s2) {
		return false
	}
	needs, windows := [26]int{}, [26]int{}
	for i := range s1 {
		needs[s1[i]-'a']++
		windows[s2[i]-'a']++
	}

	//left指针从s2最左端直到窗口右端碰到s2最右边
	for left := 0; left < len(s2)-len(s1); left++ {
		//判断两个数组内容相等
		if needs == windows {
			return true
		}
		windows[s2[left]-'a']--
		windows[s2[left+len(s1)]-'a']++
	}
	return needs == windows
}
