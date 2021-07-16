package _9_queue

import "fmt"

type ArrayQueue struct {
	q        []interface{}
	capacity int
	head     int
	tail     int
}

func NewArrayQueue(n int) *ArrayQueue {
	return &ArrayQueue{
		q:        make([]interface{}, n),
		capacity: n,
		head:     0,
		tail:     0,
	}
}

func (this *ArrayQueue) EnQueue(v interface{}) bool {
	if this.tail == this.capacity { //队满
		return false
	}
	this.q[this.tail] = v
	this.tail++
	return true
}

func (this *ArrayQueue) DeQueue() interface{} {
	if this.head == this.tail { //队空
		return nil
	}
	v := this.q[this.head]
	this.head++
	return v
}

func (this *ArrayQueue) String() string {
	if this.head == this.tail {
		return "empty queue"
	}
	result := "head"
	for i := this.head; i <= this.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", this.q[i])
	}
	result += "<-tail"
	return result
}
