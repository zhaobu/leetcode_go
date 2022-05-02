package main

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

/*
解法1 暴力复制
直接遍历复制
*/
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	help := &Node{}
	cur, pre := head, help

	//第一步先复制新链表的主干,也就是只关心Next
	for cur != nil {
		next := &Node{
			Val: cur.Val,
		}
		pre.Next = next
		pre = pre.Next
		cur = cur.Next
	}

	//第二步再根据旧链表的Random复制新链表Random
	cur1, cur2 := help.Next, head
	for cur1 != nil {
		if cur2.Random != nil {
			/*
				1. 复制出来的链表和原始链表长度一样,Random指向的节点距离头节点的个数也一样
				2. node2为原始链表中的节点,也就是原始链表cur2的Random指向的节点,
				node1为复制出来的链表中的节点,也是复制链表节点cur1的Random应该指向的节点
			*/
			node1, node2 := help.Next, head
			for node2 != cur2.Random {
				node1 = node1.Next
				node2 = node2.Next
			}
			cur1.Random = node1
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	return help.Next
}

/*
解法2 dfs
*/
func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}

	var dfs func(record map[*Node]*Node, node *Node) *Node
	dfs = func(record map[*Node]*Node, node *Node) *Node {
		if node == nil {
			return nil
		}
		if v, ok := record[node]; ok {
			return v
		}
		clone := &Node{
			Val: node.Val,
		}
		record[node] = clone
		clone.Next = dfs(record, node.Next)
		clone.Random = dfs(record, node.Random)
		return clone
	}

	return dfs(map[*Node]*Node{}, head)
}

/*
解法3 间隔复制
*/
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	//在每个节点后面复制一个节点,并且和当前链表连接在一起
	cur := head
	for cur != nil {
		clone := &Node{
			Val: cur.Val,
		}
		clone.Next = cur.Next
		cur.Next = clone
		cur = cur.Next.Next
	}
	/*
		处理Random节点
		1. head节点从0开始,偶数位节点为原始节点,奇数位为原节点复制出来的节点
		2. 复制节点的Random = 原节点的Random的下一个节点
	*/
	cur = head //开始指向原始链表的头节点
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next //间隔一个才是原始链表的下一个节点
	}

	//把原始节点和复制节点分离开
	help := &Node{}        //辅助游标节点
	cur, pre := head, help //cur:指向当前遍历到的原始链表的节点; pre:保存复制节点和原始链表拆分开后形成的复制链表的最后一个节点
	for cur != nil {
		pre.Next = cur.Next      //已经复制部分和新的复制节点连接
		pre = pre.Next           //复制链表后移,准备下一次操作
		cur.Next = cur.Next.Next //还原原始链表的指向
		cur = cur.Next           //继续遍历原始链表的下一个节点
	}

	return help.Next
}

// @lc code=end
