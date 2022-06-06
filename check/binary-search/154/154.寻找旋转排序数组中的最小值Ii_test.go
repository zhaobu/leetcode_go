package main

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{[]int{
		// 		11, 13, 15, 17,
		// 	}},
		// 	want: 11,
		// },
		{
			name: "test2",
			args: args{[]int{
				2, 2, 2, 0, 1, 2,
			}},
			want: 0,
		},
		{
			name: "test3",
			args: args{[]int{
				1, 3, 3,
			}},
			want: 1,
		},
		{
			name: "test4",
			args: args{[]int{
				2, 3, 1,
			}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
