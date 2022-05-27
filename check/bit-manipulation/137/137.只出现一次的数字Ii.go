package main

/*
 * @lc app=leetcode.cn id=137 lang=golang
 *
 * [137] 只出现一次的数字 II
 */

// @lc code=start

/*
解法1 位运算
可以推广到:一个数组中有一种数出现k次，其他数都出现了m次，m > 1, k < m, 找到出现了k次的数
https://www.cnblogs.com/greyzeng/p/15385402.html
*/
func singleNumber1(nums []int) int {
	/*
		1. 把数组中的所有数的二进制位上的信息累加起来了
		2. 比如help[1]表示数组中所有数二进制中1位置的值之和,也就是
		二进制中第1位置上为1的数的个数
	*/
	cnt := [32]int{}
	for _, v := range nums {
		for i := 0; i < 32; i++ {
			/*
				如果数v的第i位上为1,就让cnt[i]加1
			*/
			cnt[i] += ((v >> i) & 1)
		}
	}

	ret := int32(0) //一定要用32位数,因为要进行移位运算
	for i := 0; i < 32; i++ {
		if cnt[i]%3 == 1 {
			ret |= 1 << i
		}
	}
	return int(ret)
}

/*
解法2 更简洁写法
*/
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

// @lc code=end
