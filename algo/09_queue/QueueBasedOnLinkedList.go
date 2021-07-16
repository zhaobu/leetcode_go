package _9_queue

import "fmt"

type ListNode struct {
	val  interface{}
	next *ListNode
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (this *LinkedListQueue) EnQueue(v interface{}) {
	if this.tail == nil {
		this.tail = &ListNode{val: v}
		this.head = this.tail
	} else {
		this.tail.next = &ListNode{val: v}
		this.tail = this.tail.next
	}
	this.length++
}

func (this *LinkedListQueue) DeQueue() interface{} {
	if this.head == nil {
		return nil
	}
	v := this.head.val
	this.head = this.head.next
	this.length--
	return v
}

func (this *LinkedListQueue) String() string {
	if this.head == nil {
		return "empty queue"
	}
	result := "head<-"
	for cur := this.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.val)
	}
	result += "<-tail"
	return result
}
