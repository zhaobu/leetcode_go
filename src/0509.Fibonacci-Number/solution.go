package Solution

/*
[abuladong:暴力递归](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func fib1(N int) int {
	if N < 2 {
		return N
	}
	return fib1(N-1) + fib1(N-2)
}

/*
[abuladong:带备忘录的递归解法](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func fib2(N int) int {
	record := make([]int, N+1)
	return _fib2(record, N)
}

func _fib2(record []int, N int) int {
	if N < 2 {
		return N
	}
	if record[N] == 0 {
		record[N] = _fib2(record, N-1) + _fib2(record, N-2)
	}
	return record[N]
}

/*
[abuladong: dp 数组的迭代解法](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func fib3(N int) int {
	if N < 2 {
		return N
	}
	dp := make([]int, N+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}

/*
[abuladong: dp 数组的迭代解法状态压缩](https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie-qing-an-shun-xu-yue-du/dong-tai-gui-hua-xiang-jie-jin-jie)
*/
func fib4(N int) int {
	if N < 2 {
		return N
	}
	pre, cur := 0, 1
	for i := 2; i <= N; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
