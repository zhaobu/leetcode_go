package main

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
/*
解法1 循环检查二进制位
*/
func hammingWeight1(num uint32) int {
	ret := 0
	for i := 0; i < 32; i++ {
		if num&(1<<i) > 0 {
			ret++
		}
	}
	return ret
}

/*
解法2 不断消除最后一位上的1
*/
func hammingWeight(num uint32) int {
	ret := 0
	for num > 0 {
		num &= num - 1
		ret++
	}
	return ret
}

/*
解法3 利用求最右边为1的方式(不推荐,用方法2更好)
*/
func hammingWeight3(num uint32) int {
	ret := 0
	for num > 0 {
		//求出最右边的1
		firstOne := num & (-num)
		//消掉最右边的1
		num &= ^firstOne

		ret++
	}
	return ret
}

// @lc code=end
