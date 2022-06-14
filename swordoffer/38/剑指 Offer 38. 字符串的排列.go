package main

import (
	"sort"
)

/*
 *
 *
 * 剑指 Offer 38. 字符串的排列
 */

// @lc code=start

/*
 解法1 回溯
*/
func permutation1(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	strArr := []byte(s)
	sort.Slice(strArr, func(i, j int) bool {
		return strArr[i] < strArr[j]
	})

	ret := []string{}
	visited := make([]bool, len(strArr))

	/*
		1. 每一个位置都能从所有字符中选一个
		2. i表示当前选到第i个字符,record记录已经选过的字符
	*/
	var dfs func(i int, record []byte)
	dfs = func(i int, record []byte) {
		if i == len(strArr) {
			ret = append(ret, string(record))
			return
		}

		for j := range strArr {
			/*
				1. 如果当前字符已经被用,则不能选
				2. 如果当前字符没有被用,则需要看他是不是和前一个字符相同,并且前一个字符已经被用.
				则第i位上就不能重复选同样的字符
			*/
			if visited[j] || (j > 0 && strArr[j] == strArr[j-1] && visited[j-1]) {
				continue
			}
			visited[j] = true
			dfs(i+1, append(record, strArr[j]))
			visited[j] = false
		}
	}

	dfs(0, make([]byte, 0, len(strArr)))

	return ret
}

/*
解法2 交换元素
1. 每次都从剩下元素中选一个放在当前位置,不能重复选取同样的元素
2. 选取过程中利用交换进行选取
*/
func permutation2(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}

	strArr := []byte(s)
	ret := []string{}

	repeated := func(start, i int) bool {
		for start < i {
			if strArr[start] == strArr[i] {
				return true
			}
		}
		return false
	}

	/*
		1. 当前正在选取第start个元素
		2. 选取的范围是[start,len(strArr)-1]
	*/
	var selectOne func(start int)
	selectOne = func(start int) {
		if start == len(strArr) {
			ret = append(ret, string(append([]byte{}, strArr...)))
			return
		}
		for i := start; i < len(strArr); i++ {
			//如果strArr[i]在[start,i)范围上选取过,就不用重复选
			if repeated(start, i) {
				continue
			}
			strArr[start], strArr[i] = strArr[i], strArr[start]
			selectOne(start + 1)
			strArr[start], strArr[i] = strArr[i], strArr[start]
		}
	}

	selectOne(0)

	return ret
}

/*
解法3 利用下一个字典序排列思想
每次都找到比当前排列字典序更大的排列
*/
func permutation(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	//排序,方便从最小的元素开始进行下一个更大字典序的查找
	strArr := []byte(s)
	sort.Slice(strArr, func(i, j int) bool {
		return strArr[i] < strArr[j]
	})

	/*
		123456
		123465
		123546
		123564
		123654
	*/
	n := len(strArr)

	nextPermutation := func() string {
		/*
			1. 从[i,n)范围内都是非递增的
			2. i是第一个strArr[i] > strArr[i-1]的下标
		*/
		i := n - 1
		for ; i > 0; i-- {
			if strArr[i] > strArr[i-1] {
				break
			}
		}
		if i == 0 {
			return ""
		}
		/*
			1. 到[i,n)区间上第一个>strArr[i-1]的元素,并进行互换
			2. 因为[i,n)区间是非递增区间,所以从后往前找
		*/
		for j := n - 1; j >= i; j-- {
			if strArr[j] > strArr[i-1] {
				strArr[i-1], strArr[j] = strArr[j], strArr[i-1]
				break
			}
		}

		/*
			1. 在区间[i,n)范围内要进行非递减排列
			2. 因为当前已经是非递增排列的,所以只需要对半交换即可
		*/
		for i, j := i, n-1; i < j; i, j = i+1, j-1 {
			strArr[i], strArr[j] = strArr[j], strArr[i]
		}
		return string(strArr)
	}

	ret := []string{}
	cur := string(strArr)
	for len(cur) > 0 {
		ret = append(ret, cur)
		cur = nextPermutation()
	}

	return ret
}

// @lc code=end
