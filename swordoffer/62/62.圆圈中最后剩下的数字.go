package main

/*
 *
 *
 * 剑指 Offer 62. 圆圈中最后剩下的数字
 */

/*
解法1 暴力模拟(超时)
*/
func lastRemaining1(n int, m int) int {
	nums := make([]int, n)
	for i, _ := range nums {
		nums[i] = i
	}

	m--
	i := m % len(nums)
	for len(nums) > 1 {
		nums = append(nums[0:i], nums[i+1:]...)
		// fmt.Printf("nums=%+v\n", nums)
		i = (i + m) % len(nums)
	}

	return nums[0]
}

/*
解法2 递归
主要用到公式:
f(n,m) = (f(n-1, m) +m) % n
f(n,m)表示最终留下的元素的序号
公式推导参考图片
*/
func lastRemaining2(n int, m int) int {
	if n == 1 {
		//如果是只有1个数字0,则不能删除,留下的就是0
		return 0
	}
	x := lastRemaining2(n-1, m)

	return (m + x) % n
}

/*
解法3 迭代
*/

func lastRemaining(n int, m int) int {
	ret := 0

	for i := 2; i < n+1; i++ {
		ret = (m + ret) % i
	}
	return ret
}
