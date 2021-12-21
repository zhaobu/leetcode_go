package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	f2 := 1 // 1个台阶
	f1 := 2 //2个台阶
	fn := 0 //n级台阶
	for i := 3; i <= n; i++ {
		fn = f1 + f2 //f(n)=f(n-1)+f(n-2)
		f2 = f1      //更新f(n-2)的值为f(n-1)
		f1 = fn      //更新f(n-1)为最新的fn
	}
	return fn
}

// @lc code=end
