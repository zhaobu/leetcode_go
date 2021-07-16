package _8_stack

import "fmt"

/*
基于链表实现的栈
*/
type node struct {
	next *node
	val  interface{}
}

type LinkedListStack struct {
	//栈顶节点
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (t *LinkedListStack) IsEmpty() bool {
	return t.topNode == nil
}

func (t *LinkedListStack) Push(v interface{}) {
	t.topNode = &node{next: t.topNode, val: v}
}

func (t *LinkedListStack) Pop() interface{} {
	if t.IsEmpty() {
		return nil
	}
	v := t.topNode.val
	t.topNode = t.topNode.next
	return v
}

func (t *LinkedListStack) Top() interface{} {
	if t.IsEmpty() {
		return nil
	}
	return t.topNode.val
}

func (t *LinkedListStack) Flush() {
	t.topNode = nil
}

func (t *LinkedListStack) Print() {
	if t.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := t.topNode
		for nil != cur {
			fmt.Println(cur.val)
			cur = cur.next
		}
	}
}
