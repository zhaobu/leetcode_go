package _8_stack

import "fmt"

/*
基于数组实现的栈
*/

type ArrayStack struct {
	//数据
	data []interface{}
	//栈顶指针
	top int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (t *ArrayStack) IsEmpty() bool {
	return t.top < 0
}

func (t *ArrayStack) Push(v interface{}) {
	t.top += 1

	if t.top >= len(t.data) {
		t.data = append(t.data, v)
	} else {
		t.data[t.top] = v
	}
}

func (t *ArrayStack) Pop() interface{} {
	if t.IsEmpty() {
		return nil
	}
	v := t.data[t.top]
	t.top -= 1
	return v
}

func (t *ArrayStack) Top() interface{} {
	if t.IsEmpty() {
		return nil
	}
	return t.data[t.top]
}

func (t *ArrayStack) Flush() {
	t.top = -1
}

func (t *ArrayStack) Print() {
	if t.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := t.top; i >= 0; i-- {
			fmt.Println(t.data[i])
		}
	}
}
