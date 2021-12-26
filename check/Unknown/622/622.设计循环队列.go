package main

/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	q    []int //循环队列
	head int   //队头
	tail int   //队尾
	cap  int   //队列容量(数值上等于len(q)-1)
}

func Constructor(k int) MyCircularQueue {
	if k < 1 {
		return MyCircularQueue{}
	}
	return MyCircularQueue{
		cap:  k,
		q:    make([]int, k+1), //注意循环队列要牺牲一个元素大小的空间
		head: 0,
		tail: 0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	/*
		因为tail初始值为0,tail指向队尾的下一个位置,所以插入元素时tail就表示要插入的位置
		所以要先插入元素,再移动tail
	*/
	this.q[this.tail] = value
	this.tail = (this.tail + 1) % len(this.q)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.q)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.q[(len(this.q)+this.tail-1)%len(this.q)]
}

/*
队空条件：head==tail为true
*/
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

/*
队满条件：(tail+1)%capacity==head为true
*/
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%len(this.q) == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
