package main

/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start

/*
解法1
*/
func candy1(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}
	count := make([]int, n) //记录第i号孩子最少需要分得的糖果数量
	//从左往右遍历
	for i := 0; i < n; i++ {
		count[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			count[i] = count[i-1] + 1
		}
	}
	// fmt.Printf("count=%+v\n", count)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	//从右往左
	ret := count[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			count[i] = max(count[i], count[i+1]+1)
		}
		ret += count[i]
	}
	// fmt.Printf("count=%+v\n", count)
	return ret
}

/*
解法2
*/
func candy(ratings []int) int {
	n := len(ratings)
	if n == 1 {
		return 1
	}

	last := 1
	ret := last
	asc, desc := 1, 0
	for i := 1; i < n; i++ {
		/*
			1
			  2       6
			    3   5   7           12
				  4       8      11
				            9 10
		*/
		if ratings[i] >= ratings[i-1] {
			desc = 0
			if ratings[i] == ratings[i-1] {
				last = 1
			} else {
				last++
			}
			ret += last
			asc = last
		} else {
			desc++
			if desc == asc {
				desc++
			}
			ret += desc
			last = 1
		}
	}

	return ret
}

// @lc code=end
