package Solution

import "fmt"

/*
[小浩算法](https://www.geekxh.com/1.2.动态规划系列/201.html#_04、go语言示例)
*/
func ClimbStairs1(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

/*
[leetcode官方 方法一：动态规划)](https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode-solution/)
*/
func ClimbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	a := 1
	b := 2
	c := 0
	for i := 3; i < n+1; i++ {
		c = a + b
		a = b
		b = c
	}
	return b
}

/*
[windliang 解法一 暴力解法)](https://leetcode.wang/leetCode-70-Climbing-Stairs.html)
*/
func ClimbStairs3(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	return ClimbStairs3(n-1) + ClimbStairs3(n-2)
}

/*
[windliang 解法二 暴力解法优化)](https://leetcode.wang/leetCode-70-Climbing-Stairs.html)
*/
func ClimbStairs4(n int) int {
	record := map[int]int{}
	res := cal(n, record)
	fmt.Printf("record=%+v", record)
	return res
}
func cal(n int, record map[int]int) int {
	if n == 1 || n == 2 {
		return n
	}
	if v, ok := record[n]; ok {
		return v
	}
	res := cal(n-1, record) + cal(n-2, record)
	record[n] = res
	return res
}

// func cal(n int, record map[int]int) int {
// 	if n == 1 || n == 2 {
// 		return n
// 	}
// 	a, ok := record[n-1]
// 	if !ok {
// 		a = cal(n-1, record)
// 		record[n-1] = a
// 	}
// 	b, ok := record[n-2]
// 	if !ok {
// 		b = cal(n-2, record)
// 		record[n-2] = b
// 	}
// 	record[n] = a + b
// 	return a + b
// }
