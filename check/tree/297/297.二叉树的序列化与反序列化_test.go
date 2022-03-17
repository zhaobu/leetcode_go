package main

import (
	. "leetcode/check/tree"
	"reflect"
	"testing"
)

type TestArgs struct {
	root *TreeNode
	data string
}

const (
	PreOrder = iota
	InOrder
	POstOrder
)

var testArgs []*TestArgs

func init() {
	testArgs = []*TestArgs{
		{
			root: BuildTree("123##45"),
			data: "1,2,#,#,3,4,#,#,5,#,#,",
		},
	}
}

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			this: &Codec{},
			args: args{
				root: testArgs[0].root,
			},
			want: testArgs[0].data,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		this *Codec
		args args
		want *TreeNode
	}{
		{
			name: "test1",
			this: &Codec{},
			args: args{
				data: testArgs[0].data,
			},
			want: testArgs[0].root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
