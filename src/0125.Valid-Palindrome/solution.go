package Solution

import "strings"

/*
 */
func IsPalindrome1(s string) bool {
	if s == "" {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		for ; i < j && !isalnum(s[i]); i++ {
		}
		for ; i < j && !isalnum(s[j]); j-- {
		}
		if strings.ToLower(string(s[i])) != strings.ToLower(string(s[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
