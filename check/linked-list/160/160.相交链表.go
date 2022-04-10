/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
解法1
1. p1,p2分别从headA和headB开始遍历
2. 如果p1遍历完headA,就开始遍历headB,如果p2遍历完headB,就开始遍历headA
3. 这样假设headA不相交部分为a,相交部分为c, headB不相交部分为b则
p1总遍历的长度为a+c+b, p2总遍历的长度为b+c+a.所以2个指针遍历的总长度就相等.
最终如果p1==p2就说明走到了相交节点,或者最终都为nil表明没有相交
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB
	t1, t2 := true, true //分别用来保证p1,p2只会切换一次链表
	for p1 != p2 && p1 != nil && p2 != nil {
		p1 = p1.Next
		p2 = p2.Next
		if t1 && p1 == nil { //遍历完headA后就开始遍历headB
			p1 = headB
			t1 = false
		}
		if t2 && p2 == nil { //遍历完headB后就开始遍历headA
			p2 = headA
			t2 = false
		}
	}

	//最终要么走到相交节点,要么遍历到节点尾部
	return p1
}

// @lc code=end

