package main

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberN(t *testing.T) {
	type args struct {
		nums []int
		k    int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{k: 1, m: 3,
				nums: []int{
					-2, -2, -2, 1, 1, 1, 4, 4, 4,
					-4,
				},
			},
			want: -4,
		},
		{
			name: "test2",
			args: args{k: 2, m: 4,
				nums: []int{
					-2, -2, -2, -2, 2, 2, 2, 2, 4, 4, 4, 4,
					-7, -7,
				},
			},
			want: -7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberN(tt.args.nums, tt.args.k, tt.args.m); got != tt.want {
				t.Errorf("singleNumberN() = %v, want %v", got, tt.want)
			}
		})
	}
}
