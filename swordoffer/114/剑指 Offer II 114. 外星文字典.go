package main

/*
 *
 *
 * 剑指 Offer II 114. 外星文字典
 */

// @lc code=start

/*
 解法1 拓扑排序(广度优先)
*/
func alienOrder(words []string) string {
	n := len(words)

	indeg := make([]byte, 26)   //每个点的入度, indeg[i]=k 表示比字母i优先级高的字母有k个
	edges := make([][]byte, 26) //邻接表存储图的边, edges[i]表示所有比i优先级低的字母

	//标记word中每个出现的字母
	charSign := func(word string) {
		for i := range word {
			if indeg[word[i]-'a'] == 26 {
				indeg[word[i]-'a'] = 0
			}
		}
	}

	//每个字母都初始化为26,表示没出现过
	for i := range indeg {
		indeg[i] = 26
	}

	charSign(words[0])
next:
	for i := 1; i < n; i++ {
		word1, word2 := words[i-1], words[i]
		charSign(words[i])
		for j := 0; j < len(word1) && j < len(word2); j++ {
			if word1[j] != word2[j] {
				high, low := word1[j]-'a', word2[j]-'a'
				/*
					1. high的优先级比low更高
					2. indeg[i]表示的是比i的优先级更高的字母个数
					2. edges[i]表示所有比i优先级更低的字母
				*/
				edges[high] = append(edges[high], low)
				indeg[low]++ //比low优先级高的字符个数加1
				continue next
			}
		}
		if len(word1) > len(word2) {
			//题目要求前面字母都相同时,排在前面的字符串长度一定要更短
			return ""
		}
	}

	queue := []byte{}
	count := 0 //统计所有参与比较过的字母
	/*
		1. 挑选出所有入度为0的顶点
		2. 入度=-1表示未参加过比较,
		入度为0表示参加过比较,并且优先级在所有参与过比较的字符中最高
		入度为k,表示有k个字母比他的优先级更高
	*/
	for i, v := range indeg {
		if v < 26 {
			count++
			if v == 0 {
				queue = append(queue, byte(i))
			}
		}
	}

	//尝试进行拓扑排序
	ret := make([]byte, 0, count)
	for len(queue) > 0 {
		//入度为0的都是字典序更靠前的字母
		head := queue[0]
		queue = queue[1:]
		count--
		ret = append(ret, head+'a')
		/*
			1. edges[head]表示所有比head优先级低的字母
			2. 当head参与拓扑排序后,比这些字母优先级高的字母都会减少1个
		*/
		for _, low := range edges[head] {
			if indeg[low]--; indeg[low] == 0 {
				queue = append(queue, low)
			}
		}
	}
	if count != 0 {
		return ""
	}
	return string(ret)
}

/*
解法2 官方题解
用map保存字母
*/
func alienOrder2(words []string) string {
	g := map[byte][]byte{}
	inDeg := map[byte]int{}
	for _, c := range words[0] {
		inDeg[byte(c)] = 0
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			inDeg[byte(c)] += 0
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				inDeg[t[j]]++
				continue next
			}
		}
		if len(s) > len(t) {
			return ""
		}
	}

	order := make([]byte, len(inDeg))
	q := order[:0]
	for u, d := range inDeg {
		if d == 0 {
			q = append(q, u)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			if inDeg[v]--; inDeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if cap(q) == 0 {
		return string(order)
	}
	return ""
}

// @lc code=end
