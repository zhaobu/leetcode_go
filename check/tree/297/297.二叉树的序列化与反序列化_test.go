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

var testArgs [][]*TestArgs

func init() {
	//前序遍历测试数据
	pre := []*TestArgs{
		{
			root: BuildTree("123##45"),
			data: "1,2,#,#,3,4,#,#,5,#,#,",
		},
	}

	//中序遍历测试数据
	inor := []*TestArgs{
		{
			root: BuildTree("123##45"),
			data: "#,2,#,1,#,4,#,3,#,5,#,",
		},
	}

	//后序遍历测试数据
	post := []*TestArgs{
		{
			root: BuildTree("123##45"),
			data: "#,#,2,#,#,4,#,#,5,3,1,",
		},
	}

	testArgs = [][]*TestArgs{pre, inor, post}
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
		// {
		// 	name: "测试前序",
		// 	this: &Codec{},
		// 	args: args{
		// 		root: testArgs[PreOrder][0].root,
		// 	},
		// 	want: testArgs[PreOrder][0].data,
		// },
		// {
		// 	name: "测试中序",
		// 	this: &Codec{},
		// 	args: args{
		// 		root: testArgs[InOrder][0].root,
		// 	},
		// 	want: testArgs[InOrder][0].data,
		// },
		{
			name: "测试后序",
			this: &Codec{},
			args: args{
				root: testArgs[POstOrder][0].root,
			},
			want: testArgs[POstOrder][0].data,
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
		// {
		// 	name: "测试前序",
		// 	this: &Codec{},
		// 	args: args{
		// 		data: testArgs[PreOrder][0].data,
		// 	},
		// 	want: testArgs[PreOrder][0].root,
		// },
		// {
		// 	name: "测试中序",
		// 	this: &Codec{},
		// 	args: args{
		// 		data: testArgs[InOrder][0].data,
		// 	},
		// 	want: testArgs[InOrder][0].root,
		// },
		{
			name: "测试后序",
			this: &Codec{},
			args: args{
				data: testArgs[POstOrder][0].data,
			},
			want: testArgs[POstOrder][0].root,
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
