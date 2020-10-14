package Solution

import "strings"

/*
[leetcode官方 方法一：穷举法](https://leetcode-cn.com/problems/rotate-string/solution/xuan-zhuan-zi-fu-chuan-by-leetcode/)
*/
func RotateString1(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == "" {
		return true
	}
	n := len(A)
loop:
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if A[(i+j)%n] != B[j] {
				continue loop
			}
		}
		return true
	}
	return false
}

func RotateString2(A string, B string) bool {
	return len(A) == len(B) && strings.Contains(A+A, B)
}
