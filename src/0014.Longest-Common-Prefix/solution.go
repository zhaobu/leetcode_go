package Solution

import "strings"

/* 纵向扫描
时间复杂度：O(mn)O(mn)，其中 mm 是字符串数组中的字符串的平均长度，nn 是字符串的数量。最坏情况下，字符串数组中的每个字符串的每个字符都会被比较一次。

空间复杂度：O(1)O(1)。使用的额外空间复杂度为常数。
*/
func LongestCommonPrefix1(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	res := ""
	//遍历strs[0]每个字符
	for i := range strs[0] {
		ch := strs[0][i]
		for _, v := range strs {
			if i >= len(v) || v[i] != ch {
				return res
			}
		}

		res += string(ch)
	}

	return res
}

/* 使用小浩算法的思路,使用strings.Index

 */
func LongestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
