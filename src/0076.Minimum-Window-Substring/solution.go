package Solution

/*
[labuladong 滑动窗口算法 专题](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485141&idx=1&sn=0e4583ad935e76e9a3f6793792e60734&chksm=9bd7f8ddaca071cbb7570b2433290e5e2628d20473022a5517271de6d6e50783961bebc3dd3b&scene=21#wechat_redirect)
*/
func minWindow1(s string, t string) string {
	//need记录t每个字符需要的个数,window记录窗口中出现t中的每个字符的个数
	need, window := make(map[byte]int, len(t)), make(map[byte]int, len(t))
	for i := range t {
		need[t[i]]++
	}
	//left,right分别为左闭右开的滑动窗口的左右边界,valid为已经满足要求的字符个数
	left, right, valid := 0, 0, 0
	//记录最小覆盖子串的起始索引及长度,minLen初始化为一个不可能达到的长度
	start, minLen := 0, len(s)+1
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] { //过滤重复出现的情况
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < minLen {
				start = left
				minLen = right - left
				if minLen == len(t) { //不可能存在比len(t)还短的情况
					return s[start : start+minLen]
				}
			}
			// d 是将移出窗口的字符
			d := s[left]
			// 左移窗口
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
	// 返回最小覆盖子串
	if minLen == len(s)+1 {
		return ""
	}
	return s[start : start+minLen]
}
