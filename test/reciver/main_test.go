package main

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRet [][]int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := threeSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("threeSum() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name    string
		args    args
		wantRet []int
	}{
		{
			name: "test1",
			args: args{
				nums1: []int{1, 1},
				nums2: []int{},
			},
			wantRet: []int{},
		},
		{
			name: "test1",
			args: args{
				nums1: []int{},
				nums2: []int{1, 1},
			},
			wantRet: []int{},
		},
		{
			name: "test1",
			args: args{
				nums1: []int{1, 1, 1, 2},
				nums2: []int{},
			},
			wantRet: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := merge(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(gotRet, tt.wantRet) {
				t.Errorf("merge() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
