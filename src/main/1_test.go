package main

import (
	"fmt"
	"leetcode/utils/zaplog"
	"testing"
)

func init() {
	zaplog.InitLog(&zaplog.Config{
		Logdir:   "./",
		LogName:  "1_test.log",
		Loglevel: "debug",
		Debug:    true},
	)
}

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

//删除重复节点,保留一个
func delRepeatNode(head *ListNode) (resHead *ListNode) {
	if head == nil {
		return nil
	}

	var (
		cur  *ListNode = head
		next *ListNode
	)
	for cur != nil {
		next = cur.next
		for next != nil && next.val == cur.val {
			next = next.next
		}

		cur.next = next
		cur = next
	}
	return head
}

//删除所有重复节点
/*
思路:
1.用一个带头链表记录结果,避免判断头节点就是重复时做特殊处理
2.整体就是从第一个节点开始依次判断后面的节点是否和本节点相同,直到找到和本节点不同的下一个节点.同时要记录
相同节点的个数,用来判断是否重复.这样依次每种不同值的节点都会判断一次
3. 判断到某个值的节点不重复时加入结果,重复时要把结果.next指向nil.
*/
func delRepeatNodes(head *ListNode) (resHead *ListNode) {
	if head == nil {
		return nil
	}

	resHead = &ListNode{} //带头链表,用来记录不重复的节点
	var (
		pre  *ListNode = resHead //上一个不重复的节点
		cur  *ListNode = head    //当前正在判断是否重复的第一个节点
		next *ListNode           //
	)
	for cur != nil {
		next = cur
		count := 1 //记录和cur值相同的节点个数
		//从cur开始往下找和cur.val相同的节点
		for next != nil && next.next != nil && next.val == next.next.val {
			count++
			next = next.next
		}
		if count == 1 { //如果和cur相同的节点只有一个
			// 把当前节点加入到不重复节点链表
			pre.next = cur
			pre = cur
			cur = cur.next
			continue
		} else {
			//如果和cur相同的节点多余1个,此时next指向本次所有相同节点的最后一个
			cur = next.next
			//如果本次判断的节点不能加入,要把pre的next指向nil
			pre.next = nil
		}
	}
	return resHead.next
}

//删除所有重复节点
func delRepeatNodes1(head *ListNode) (resHead *ListNode) {
	if head == nil {
		return nil
	}

	resHead = &ListNode{} //带头链表,用来记录不重复的节点
	var (
		pre *ListNode = resHead //上一个不重复的节点
		cur *ListNode = head    //当前正在判断是否重复的第一个节点
	)
	count := 1 //记录节点出现的次数
	for cur != nil {
		if cur.next == nil { //遍历到了链表尾节点
			if count == 1 {
				pre.next = cur
				pre = cur
			}
		} else {
			if cur.next.val == cur.val { //下一个节点和当前节点一样继续遍历
				count++
			} else { //下一个节点和当前节点不同就判断如何记录结果
				if count == 1 {
					pre.next = cur
					pre = cur
				} else {
					//如果本次判断的节点不能加入,要把pre的next指向nil
					pre.next = nil
				}
				//下一个节点和当前节点不一样时将继续下一轮判断
				count = 1
			}
		}
		cur = cur.next
	}
	return resHead.next
}

func (t *LinkedList) Print() {
	cur := t.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.val)
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	zaplog.Info(format)
}

var (
	nodes []int = []int{1, 1, 2, 3, 4, 4, 6, 6}
)

func Test_delRepeatNodes(t *testing.T) {
	linkList := &LinkedList{head: &ListNode{}}
	cur := linkList.head
	nodes = []int{1, 1, 2, 3, 4, 4, 6, 6}

	for _, v := range nodes {
		newNode := &ListNode{val: v}
		cur.next = newNode
		cur = cur.next
	}
	linkList.Print()

	linkList.head.next = delRepeatNodes(linkList.head.next)

	linkList.Print()
}

func Test_delRepeatNodes1(t *testing.T) {
	linkList := &LinkedList{head: &ListNode{}}
	cur := linkList.head
	nodes = []int{1, 1, 2, 3, 4, 4, 6, 6}

	for _, v := range nodes {
		newNode := &ListNode{val: v}
		cur.next = newNode
		cur = cur.next
	}
	linkList.Print()

	linkList.head.next = delRepeatNodes1(linkList.head.next)

	linkList.Print()
}

func Test_delRepeatNode(t *testing.T) {
	linkList := &LinkedList{head: &ListNode{}}
	cur := linkList.head
	nodes = []int{1, 1, 2, 3, 4, 4, 6, 6}

	for _, v := range nodes {
		newNode := &ListNode{val: v}
		cur.next = newNode
		cur = cur.next
	}
	linkList.Print()

	linkList.head.next = delRepeatNode(linkList.head.next)

	linkList.Print()
}
