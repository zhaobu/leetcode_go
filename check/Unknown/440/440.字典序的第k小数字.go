package main

/*
 * @lc app=leetcode.cn id=440 lang=golang
 *
 * [440] 字典序的第K小数字
 */

// @lc code=start

/*
解法1 暴力法
386题是排序,可以直接求出第k个数字
*/
func findKthNumber1(n int, k int) int {
	cur := 1
	for i := 1; i < k; i++ {
		//尝试找到下一层的根节点
		if cur*10 <= n {
			cur *= 10
		} else {
			/*
				1. 下一层的根节点不满足条件,就尝试当前层往右
				2. cur+1必须在范围[1,n]之内
				3. cur的最后一个数不能超过9
			*/
			for cur+1 > n || cur%10 == 9 {
				cur /= 10 //如果cur+1后不满足条件,就回退到上一层的第一个数
			}
			cur++
		}
	}
	return cur
}

/*
解法2  字典树
*/
func findKthNumber(n int, k int) int {
	cur := 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	/*
		求以cur为根节点的这颗子树满足条件的节点个数
	*/
	getNum := func(cur int) int {
		count := 0
		first, last := cur, cur
		for first <= n {
			count += min(last, n) - first + 1
			first *= 10
			last = last*10 + 9
		}
		return count
	}

	i := 1 //第i小的数
	for i < k {
		num := getNum(cur)
		if i+num <= k {
			/*
				1. 如果以cur为根的树满足条件的节点+i<k,
				说明第k小的数不在这颗树当中,应该往右继续查找以兄弟节点为根的子树
			*/
			i += num
			cur++
		} else {
			i += 1
			cur *= 10
		}
	}
	return cur
}

// @lc code=end
