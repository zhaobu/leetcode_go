package main

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// {
		// 	name: "test1",
		// 	args: args{
		// 		nums: []int{0, 0},
		// 	},
		// 	want: 1,
		// },
		{
			name: "test2",
			args: args{
				nums: []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
