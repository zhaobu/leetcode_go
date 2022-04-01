package main

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test1",
		// 	args: args{
		// 		nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		// 		k:    4,
		// 	},
		// 	want: 4,
		// },
		// {
		// 	name: "test2",
		// 	args: args{
		// 		nums: []int{3, 1, 2, 4},
		// 		k:    2,
		// 	},
		// 	want: 3,
		// },
		{
			name: "test3",
			args: args{
				nums: []int{2, 1},
				k:    1,
			},
			want: 2,
		},
		{
			name: "test4",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
				k:    1,
			},
			want: 6,
		},
		{
			name: "test5",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
