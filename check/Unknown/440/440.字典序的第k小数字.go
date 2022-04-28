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
		层序遍历,求以cur为根节点的这颗子树满足条件的节点个数
	*/
	getNum := func(cur int) (count int, nodes []int) {
		first, last := cur, cur //当前层第一个节点为first,最后一个节点为last
		for first <= n {
			count += min(last, n) - first + 1
			// for i := first; i <= n && i <= last; i++ {
			// 	nodes = append(nodes, i)
			// }
			first *= 10        //下一层第一个节点为当前层第一个节点*10
			last = last*10 + 9 //下一层最后一个节点为当前层最后一个节点*10+9
		}
		return
	}

	i := 1 //第i小的数
	// ret := []int{}
	for i < k {
		num, _ := getNum(cur)
		if i+num <= k {
			/*
				1. 如果以cur为根的树满足条件的节点num+i<k,说明第k小的数不在这颗树当中
				2. 当前子树节点应该全部跳过,所以i = i+num
			*/
			i += num
			// ret = append(ret, nodes...)
			cur++ //往右继续查找第一个兄弟节点为根的子树
		} else {
			/*
				1. 如果以cur为根的树满足条件的节点num+i>k,说明第k小的节点就在以cur为根节点的子树当中
				2. 把当前树的根节点cur加到结果当中,继续从当前树的第一个孩子节点开始查找
			*/
			i += 1
			// ret = append(ret, nodes[:2]...)
			cur *= 10 //cur跳到子树最左边第一个孩子节点开始继续查找
		}
	}
	// fmt.Printf("len=%d,ret=%+v\n", len(ret), ret)
	return cur
}

// @lc code=end
