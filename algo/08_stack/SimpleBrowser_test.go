package _8_stack

import (
	"fmt"
	"testing"
)

func TestBrowser(t *testing.T) {
	b := NewBrowser()
	b.PushBack("www.qq.com")
	b.PushBack("www.baidu.com")
	b.PushBack("www.sina.com")
	if b.CanBack() {
		b.Back()
	}
	if b.CanForward() {
		b.Forward()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	if b.CanBack() {
		b.Back()
	}
	b.Open("www.taobao.com")
	if b.CanForward() {
		b.Forward()
	}
}

func TestArr(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("arr=%+v", arr)
	arr = arr[0 : len(arr)-1]
	fmt.Printf("arr=%+v", arr)
}
