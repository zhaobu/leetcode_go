/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	stack1 []int //入队时都入栈1
	stack2 []int //出队时先从栈2弹出,栈2为空后,再把栈1全部入栈2
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0, 30),
		stack2: make([]int, 0, 30),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1 = append(this.stack1, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	//栈2为空时,先把栈1元素全部出栈并入栈到栈2
	if len(this.stack2) == 0 {
		//栈1全部弹出并入栈2
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		//清空栈1
		this.stack1 = this.stack1[0:0]
	}

	// 从栈2出队列
	res := this.stack2[len(this.stack2)-1]
	this.stack2 = this.stack2[0 : len(this.stack2)-1]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	//栈2为空时,先把栈1元素全部出栈并入栈到栈2
	if len(this.stack2) == 0 {
		//栈1全部弹出并入栈2
		for i := len(this.stack1) - 1; i >= 0; i-- {
			this.stack2 = append(this.stack2, this.stack1[i])
		}
		//清空栈1
		this.stack1 = this.stack1[0:0]
	}

	// 栈2栈顶元素就是队头
	return this.stack2[len(this.stack2)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

