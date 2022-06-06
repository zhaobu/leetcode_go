package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test2",
			args: args{
				nums:   []int{5, 1, 3},
				target: 2,
			},
			want: -1,
		},
		{
			name: "test3",
			args: args{
				nums:   []int{3, 1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				nums:   []int{1, 3},
				target: 2,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
