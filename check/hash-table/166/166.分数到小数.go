package main

import "strconv"

/*
 * @lc app=leetcode.cn id=166 lang=golang
 *
 * [166] 分数到小数
 */

// @lc code=start
/*
解法1 模拟长除法
*/
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	//符号位
	ret := ""
	if numerator*denominator < 0 {
		ret = "-"
	}

	//整数部分
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}

	integer := numerator / denominator
	ret += strconv.Itoa(integer)
	if numerator%denominator == 0 {
		return ret
	}

	//不能整除就会存在小数
	ret += "."
	//处理掉整数部分
	if integer > 0 {
		numerator -= integer * denominator
	}

	decimal := "" //小数部分

	/*
		1. start表示重复小数部分开始的下标,end表示重复小数部分最后一个元素的后一个位置
		2. 比如小数3.01231 start=1, end= 4
		3. start初始化为-1,如果最后计算结果不是无限循环小数,则start=-1
	*/
	start, end := -1, 0

	/*
		1. record记录每种余数第一次出现时的下标
		2. 如果有一种余数第二次出现,就表示存在重复小数部分,否则就不会无限循环
	*/
	record := map[int]int{}
	for numerator != 0 {
		if i, ok := record[numerator]; ok {
			start = i
			break
		}
		record[numerator] = end
		end++
		numerator *= 10
		integer = numerator / denominator
		decimal += strconv.Itoa(integer)
		numerator -= integer * denominator
	}

	if start == -1 {
		//不存在循环小数
		return ret + decimal
	}

	//3.01231
	ret += decimal[:start] + "(" + decimal[start:end] + ")"

	return ret
}

// @lc code=end
