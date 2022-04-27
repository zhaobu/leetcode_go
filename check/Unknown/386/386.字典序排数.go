package main

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

// @lc code=start

/*
解法1 深度优先搜索
最小的字典序数字为1,每次找下一个满足条件的数字时:
1. 基本思想就是前序遍历10叉树
2. 优先考虑尝试在 number 后面附加一个零，即 number×10，如果 number×10≤n，
那么说明 number×10 是下一个字典序整数
3. 如果不能就尝试往当前层右边走,也就是number+1,但是number+1必须满足<=n,同时个位上的数字如果=9
也就是表明已经是最后一个数字,此时就要往上一层的第一个数回退,也就是number/10
*/
func lexicalOrder1(n int) []int {
	ret := make([]int, n)
	cur := 1
	for i := 0; i < n; i++ {
		ret[i] = cur
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

	return ret
}

/*
解法2 dfs递归
*/
func lexicalOrder(n int) []int {
	ret := make([]int, 0, n)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur > n {
			return
		}
		ret = append(ret, cur)
		next := cur * 10
		for i := 0; i < 10; i++ {
			dfs(next + i)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return ret
}

// @lc code=end
