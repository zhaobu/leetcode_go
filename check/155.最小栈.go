/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	minVal []int
	stack  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:  make([]int, 0, 30),
		minVal: make([]int, 0, 30),
	}
}

func (this *MinStack) Push(val int) {
	if len(this.minVal) > 0 && this.minVal[len(this.minVal)-1] < val {
		this.minVal = append(this.minVal, this.minVal[len(this.minVal)-1])
	} else {
		this.minVal = append(this.minVal, val)
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[0 : len(this.stack)-1]
	this.minVal = this.minVal[0 : len(this.minVal)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minVal) == 0 {
		return -1
	}
	return this.minVal[len(this.minVal)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

