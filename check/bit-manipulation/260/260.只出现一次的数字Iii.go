package main

/*
 * @lc app=leetcode.cn id=260 lang=golang
 *
 * [260] 只出现一次的数字 III
 */

// @lc code=start

/*
解法1 位运算
1. 求最低位为1的数可以用(^a+1)&a也可以用(-a)&a
2. 正数的原码,反码,补码都相等
拿52来举例:
0110100 52的原码
1001011 取反
1001100 取反后+1

1110100 -52的原码
1001011 -52的反码(负数的反码是符号位不变,其他位取反)
1001100 -52的补码(负数的补码是反码+1)

也就是说一个正数取反后+1就等于它的相反数的补码
而在计算机中存储的都是补码,所以 (^a+1)&a 等价于 (-a)^a
*/
func singleNumber(nums []int) []int {
	//先求出a^b
	ans := 0
	for i := range nums {
		ans ^= nums[i]
	}
	/*

	 1. 0110100 原码
	 2. 1001011 取反后
	 3. 1001100 取反后+1
	 4. 1和3的结果相与 0000100
	*/
	//再求出最低位的1,假设是第i位
	firstOne := (^ans + 1) & ans

	/*
		1. 假设a的第i位为1, b的第i位为0
		2. 按第i位为0和1把所有数分为2类
		3. 出现2次的数要么属于第i位为0的一类,要么属于第i位为1的一类
		4. a和b不可能在同一类
		5. 把第i位为0的一类数全部异或,则最后结果必定定于b
		把第i位为1的一类数全部异或,最后结果必定等于a
	*/
	a := 0
	for i := range nums {
		if nums[i]&firstOne > 0 { //找出第i位为1的一类数
			a ^= nums[i]
		}
	}

	//b = a^(a^b)
	return []int{a, a ^ ans}
}

// @lc code=end
