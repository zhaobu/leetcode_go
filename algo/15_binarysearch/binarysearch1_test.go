package _5_binarysearch

import (
	"fmt"
	"testing"
)

type args struct {
	a []int
	v int
}
type Case struct {
	name string
	args args
	want int
}

func TestBinarySearch1(t *testing.T) {
	nums := []int{1, 3, 5, 6, 8}
	finds := []int{0, 1, 4, 5, 8, 10}
	results := []int{-1, 0, -1, 2, 4, -1}

	tests := []*Case{}
	for i := range finds {
		tests = append(tests, &Case{
			name: fmt.Sprintf("test%d", i),
			args: args{
				a: nums,
				v: finds[i],
			},
			want: results[i],
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch1(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("BinarySearch1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchFirst1(t *testing.T) {
	nums := []int{1, 2, 2, 2, 4, 8}
	finds := []int{0, 1, 2, 4, 5, 10}
	results := []int{-1, 0, 1, 4, -1, -1}

	tests := []*Case{}
	for i := range finds {
		tests = append(tests, &Case{
			name: fmt.Sprintf("test%d", i),
			args: args{
				a: nums,
				v: finds[i],
			},
			want: results[i],
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchFirst1(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("BinarySearchFirst1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchLast1(t *testing.T) {
	nums := []int{1, 2, 2, 2, 3, 4}
	finds := []int{0, 1, 2, 4, 5, 10}
	results := []int{-1, 0, 3, 5, -1, -1}

	tests := []*Case{}
	for i := range finds {
		tests = append(tests, &Case{
			name: fmt.Sprintf("test%d", i),
			args: args{
				a: nums,
				v: finds[i],
			},
			want: results[i],
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchLast1(tt.args.a, tt.args.v); got != tt.want {
				t.Errorf("BinarySearchLast1() = %v, want %v", got, tt.want)
			}
		})
	}
}
