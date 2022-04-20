package main

/*
 * @lc app=leetcode.cn id=402 lang=golang
 *
 * [402] 移掉 K 位数字
 */

// @lc code=start

/*
解法1 暴力法
直接删除k次
1. 从左往右找到第一个i, num[i]>num[i+1]
2. 如果不存在1的情况,就说明num表示的数字是单调递增的,就删除最后一个数字
3. 执行上面的步骤k次
*/
func removeKdigits1(num string, k int) string {
	numByte := []byte(num)

	asc := false //整个字符已经是单调递增
	//执行删除1个字符k次
	for n := 0; n < k; n++ {
		m := len(numByte)
		if asc {
			//如果已经是单调递增,只需要再删除后面的k-n个字符即可
			numByte = numByte[:m-(k-n)]
			break
		} else {
			i := 0
			for i+1 < m && numByte[i] <= numByte[i+1] {
				i++
			}
			if i+1 == m {
				asc = true
			}
			//删除第i个字符
			numByte = append(numByte[:i], numByte[i+1:]...)
		}
	}

	//删除前导0
	i := 0
	for i < len(numByte) && numByte[i] == '0' {
		i++
	}
	numByte = numByte[i:]
	//如果最后字符串为空,要返回0
	if len(numByte) == 0 {
		return "0"
	}
	return string(numByte)
}

/*
解法2 单调栈
在暴力法的基础上,因为每次都是要删除从左到右第一个num[i]>num[i+1]的num[i]元素
所以可以考虑用一个单调递增栈保存从左到右已经遍历过的元素,这样在每次遍历到元素
num[i]时,既可以看栈顶元素,弹出那些比当前元素大的,每弹出一个,相当于解法1的删除了一个字符
*/
func removeKdigits(num string, k int) string {
	stack := []byte{}
	n := 0 //已经删除的字符个数
	for i := 0; i < len(num); i++ {
		for n < k && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			n++
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num[i])
	}
	//如果最后删除的元素不足k个,需要再删除掉结尾的k-n个字符
	stack = stack[:len(stack)-(k-n)]

	//删除前导0
	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}
	stack = stack[i:]
	//如果最后字符串为空,要返回0
	if len(stack) == 0 {
		return "0"
	}
	return string(stack)
}

// @lc code=end
