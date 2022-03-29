package main

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start

/*
解法1 先转换为string,再转换回来

*/
func reverse1(x int) int {
	if x == 0 {
		return 0
	}
	m := x
	if x < 0 { //负数变正数
		m = -x
	}
	xByte := []byte(strconv.Itoa(m))
	for i, j := 0, len(xByte)-1; i < j; i, j = i+1, j-1 {
		xByte[i], xByte[j] = xByte[j], xByte[i]
	}
	newX, _ := strconv.Atoi(string(xByte))
	if x < 0 { //变回负数
		newX = -newX
	}
	if newX < -(1<<31) || newX > 1<<31-1 {
		return 0
	}

	return newX
}

/*
解法2 每次取个位数,然后拼接
*/
func reverse2(x int) int {
	if x == 0 {
		return 0
	}
	m := x
	if x < 0 { //负数变正数
		m = -x
	}
	newX := 0
	for m != 0 {
		newX = newX*10 + m%10
		m /= 10
	}
	if x < 0 { //变回负数
		newX = -newX
	}
	if newX < -(1<<31) || newX > 1<<31-1 {
		return 0
	}

	return newX
}

/*
解法3 官方题解
*/
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}

// @lc code=end
