package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start

/*
解法1 模拟计算乘法过程

	   123
	 * 456
	_________
	   738
	  615
     492
  _______________
      6888
	 492
*/
func multiply1(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	multiplyOnce := func(s1, a string) string {
		b, _ := strconv.Atoi(a)
		pre := 0
		retStr := ""
		for i := len(s1) - 1; i >= 0; i-- {
			c := int(s1[i] - '0')
			cur := c*b + pre
			pre = cur / 10
			retStr = strconv.Itoa(cur%10) + retStr
		}
		if pre > 0 {
			retStr = strconv.Itoa(pre) + retStr
		}
		// fmt.Printf("multiplyOnce s1=%s,a=%s,retStr=%s\n", s1, a, retStr)
		return retStr
	}

	addOnce := func(s1, s2 string, count int) string {
		if s1 == "" {
			return s2
		}

		retStr := s1[len(s1)-count:]
		s1 = s1[:len(s1)-count]
		pre := 0
		for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
			a, b := 0, 0
			if i >= 0 {
				a = int(s1[i] - '0')
			}
			if j >= 0 {
				b = int(s2[j] - '0')
			}
			cur := a + b + pre
			retStr = strconv.Itoa(cur%10) + retStr
			pre = cur / 10
		}
		if pre > 0 {
			retStr = strconv.Itoa(pre) + retStr
		}
		// fmt.Printf("addOnce s1=%s,s2=%s,retStr=%s\n", s1, s2, retStr)

		return retStr
	}

	ret := ""
	count := 0 //控制对齐相加
	for j := len(num2) - 1; j >= 0; j-- {
		curRet := multiplyOnce(num1, num2[j:j+1])
		ret = addOnce(ret, curRet, count)
		count++
	}

	return ret
}

/*
解法2 模拟计算乘法过程
先用数组保存
 		  123
 		  456
       _____________
 		   12 15 18
 		 8 10 12
	   4 5  6
	___________________
	   5  6   0   8  8
*/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	// fmt.Printf("ansArr=%+v\n", ansArr)
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	ans := ""
	idx := 0
	if ansArr[0] == 0 {
		idx = 1
	}
	// fmt.Printf("ansArr=%+v\n", ansArr)
	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArr[idx])
	}
	return ans
}

// @lc code=end
