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

func delRepeatNodes(head *ListNode) (resHead *ListNode) {
	if head == nil {
		return nil
	}

	var (
		sentry *ListNode = &ListNode{} //用来记录不重复的节点
		pre    *ListNode = sentry      //上一个不重复的节点
		cur    *ListNode = head        //当前正在判断是否重复的第一个节点
		next   *ListNode               //
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
		}
	}
	return sentry
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

func Test_delRepeatNodes(t *testing.T) {
	linkList := &LinkedList{head: &ListNode{}}
	cur := linkList.head
	arr := []int{1, 1, 2, 3, 3, 4}

	for _, v := range arr {
		newNode := &ListNode{val: v}
		cur.next = newNode
		cur = cur.next
	}
	linkList.Print()

	linkList.head = delRepeatNodes(linkList.head.next)

	linkList.Print()
}
