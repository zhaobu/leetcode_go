package Solution

/*
[leetcode 官方 方法二：二分查找](https://leetcode-cn.com/problems/sqrtx/solution/x-de-ping-fang-gen-by-leetcode-solution/)
*/
func MySqrt1(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

/*
[二分查找(套用模板,最后返回时打补丁)](https://labuladong.gitbook.io/algo/suan-fa-si-wei-xi-lie/er-fen-cha-zhao-xiang-jie#san-xun-zhao-you-ce-bian-jie-de-er-fen-cha-zhao)
*/
func MySqrt2(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r-l)>>1
		a := mid * mid
		if a == x {
			return mid
		} else if a < x {
			l = mid + 1
		} else if a > x {
			r = mid - 1
		}
	}
	return r
}
