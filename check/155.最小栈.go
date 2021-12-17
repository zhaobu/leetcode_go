/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

type MinStack struct {
	ValStack []int
	MinStack []int
}

func Constructor() MinStack {
	return MinStack{
		ValStack: []int{},
		MinStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.ValStack = append(this.ValStack, val)
	if len(this.MinStack) < 1 { //空栈则当前元素最小
		this.MinStack = append(this.MinStack, val)
	} else {
		min := this.MinStack[len(this.MinStack)-1] //取出栈顶元素和当前元素比较,谁小谁就入栈
		if val < min {
			min = val
		}
		this.MinStack = append(this.MinStack, min)
	}
}

func (this *MinStack) Pop() {
	if len(this.ValStack) < 1 {
		return
	}
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
	this.ValStack = this.ValStack[:len(this.ValStack)-1]
}

func (this *MinStack) Top() int {
	return this.ValStack[len(this.ValStack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinStack[len(this.MinStack)-1]
}

/* type MyStack struct {
	Val int
	Min int
}
type MinStack struct {
	Stack []*MyStack
}

func Constructor() MinStack {
	return MinStack{Stack: []*MyStack{}}
}

func (this *MinStack) Push(val int) {
	if len(this.Stack) < 1 {
		this.Stack = append(this.Stack, &MyStack{Val: val, Min: val})
	} else {
		min := this.Stack[len(this.Stack)-1].Min
		if val < min {
			min = val
		}
		this.Stack = append(this.Stack, &MyStack{Val: val, Min: min})
	}

}

func (this *MinStack) Pop() {
	if len(this.Stack) < 1 {
		return
	}
	this.Stack = this.Stack[:len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.Stack[len(this.Stack)-1].Min
} */

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

