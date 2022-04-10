package main

import (
	"fmt"
	"leetcode/check"
	. "leetcode/check"
	"reflect"
	"testing"
)

func Test_reverseOnce(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name        string
		args        args
		wantNewHead *ListNode
		wantNewTail *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewHead, gotNewTail := reverseOnce(tt.args.head, tt.args.n)
			if !reflect.DeepEqual(gotNewHead, tt.wantNewHead) {
				t.Errorf("reverseOnce() gotNewHead = %v, want %v", gotNewHead, tt.wantNewHead)
			}
			if !reflect.DeepEqual(gotNewTail, tt.wantNewTail) {
				t.Errorf("reverseOnce() gotNewTail = %v, want %v", gotNewTail, tt.wantNewTail)
			}
		})
	}
}

type ReverseKGroup struct {
	input *Input
	name  string
}

func (m *ReverseKGroup) GetInput() interface{} {
	return m.input
}

func (m *ReverseKGroup) GetName() string {
	return m.name
}

// func (m *ReverseKGroup) Copy() CaseIface {
// 	new := &ReverseKGroup{
// 		name:  m.name,
// 		input: m.input.copy(),
// 	}
// 	return new
// }

func (m *ReverseKGroup) Print() string {
	return fmt.Sprintf("name:%s,input:%s", m.name, m.input.Print())
}

func Test1(t *testing.T) {

	//	测试用例
	cases := []check.CaseIface{
		&ReverseKGroup{name: "test 1", input: &Input{head: UnmarshalListBySlice([]int{1, 2, 3, 4, 5})}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.GetName(), func(t *testing.T) {
			t.Logf("success  input:%+v", c.GetInput().(*Input).head.Print())
			out := reverseKGroup(c.GetInput().(*Input).head, 2)
			t.Logf("success  output:%+v\n", out.Print())
		})
	}
}
