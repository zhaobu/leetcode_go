package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

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
