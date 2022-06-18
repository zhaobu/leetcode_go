package check

import (
	"fmt"
	"math/rand"
)

type Node struct {
	Val  int
	Next *Node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// func (l *ListNode) String() string {
// 	res := ""
// 	num := 0
// 	for l != nil && num < 10000 {
// 		res += fmt.Sprintf("val:%d,address:%p --> ", l.Val, l)
// 		l = l.Next
// 		num++
// 	}
// 	return res
// }

func (l *ListNode) String() string {
	res := "["
	num := 0
	for l != nil && num < 10000 {
		res += fmt.Sprintf("%d,", l.Val)
		l = l.Next
		num++
	}
	res += "]"
	return res
}

func (l *ListNode) Copy() *ListNode {
	res := &ListNode{}
	cur := res

	//防止循环链表进入死循环
	vals := make(map[*ListNode]int)
	for i := l; i != nil; i = i.Next {
		_, ok := vals[i]
		if ok {
			break
		}
		vals[i] = 1
		cur.Next = &ListNode{Val: i.Val}
		cur = cur.Next
	}
	return res.Next
}

//定义结构
type Input struct {
	head *ListNode
}

func (i *Input) copy() *Input {
	return &Input{
		head: i.head.Copy(),
	}
}

func (i *Input) String() string {
	return fmt.Sprintf("head:%s", i.head.String())
}

// 给list添加一个环
func getRandList(hasCycle bool, len int) *ListNode {

	head := UnmarshalListByRand(100, len)
	//	生成没有环的链表
	if !hasCycle {
		return head
	}

	tmp := head

	// 生成有环的链表，让第index个及节点指向头结点
	index := rand.Intn(len)
	for i := 0; i < index; i++ {
		tmp = tmp.Next
	}

	tmp.Next = head
	return head
}

// 随机初始化链表
func UnmarshalListByRand(max_num int, len int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head

	for i := 0; i < len; i++ {
		tmp.Next = &ListNode{Val: rand.Intn(max_num), Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}

//	根据数组反序列化链表
func UnmarshalListBySlice(nums []int) *ListNode {
	head := &ListNode{Val: -1, Next: nil}
	tmp := head
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v, Next: nil}
		tmp = tmp.Next
	}
	return head.Next
}
