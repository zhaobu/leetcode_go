package main

import "sort"

/*
 * @lc app=leetcode.cn id=668 lang=golang
 *
 * [668] 乘法表中第k小的数
 */

// @lc code=start

/*
解法1 二分查找
1. 求第几小等价于求有多少数字不超过 x
2. 乘法表上的第i行上的某个数x.当前所在的列为x/i
当前行上<=x的数有x/i个
3. 所以求不超过数x的个数就是把每一行上不超过x的个数相加
4. 在i <= x/n的行上,也就是说最后一个数字i*n <= x.则这些行满足条件的数为n个
*/
func findKthNumber1(m int, n int, k int) int {
	//求<=num的数字的个数
	getNum := func(num int) int {
		pivot := num / n
		count := pivot * n
		for i := pivot + 1; i <= m; i++ {
			count += num / i
		}
		return count
	}
	/*
		二分查找第一个>=k时的位置
		1. 因为getNum求得的是<=mid的数字的个数,可能不会恰好等于k
		2. 当getNum求得的个数<k时,说明<=mid的个数少于k个,所以mid不是第k小的数字
		3. 当getNum求得的个数>=k时,说明<=mid的个数恰好等于k个或者是比k个多,但是第k小的数也一定和此时的mid相等
	*/
	i, j := 1, m*n
	for i <= j {
		mid := i + (j-i)>>1
		count := getNum(mid)
		if count < k {
			i = mid + 1
		} else {
			if mid == 1 || getNum(mid-1) < k {
				return mid
			}
			j = mid - 1
		}
	}

	return -1
}

/*
解法2 二分查找(调用库函数写法)
*/
func findKthNumber(m int, n int, k int) int {
	return sort.Search(m*n, func(num int) bool {
		pivot := num / n
		count := pivot * n
		for i := pivot + 1; i <= m; i++ {
			count += num / i
		}
		return count >= k
	})
}

// @lc code=end
