package main

import "testing"

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "test1",
		// 	args: args{matchsticks: []int{1, 1, 2, 2, 2}},
		// 	want: true,
		// },
		{
			name: "test2",
			args: args{matchsticks: []int{3, 3, 2, 2, 2}},
			want: false,
		},
		// {
		// 	name: "test3",
		// 	args: args{matchsticks: []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}},
		// 	want: true,
		// },
		// {
		// 	name: "test4",
		// 	args: args{matchsticks: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 4, 3, 2, 1}},
		// 	want: false,
		// },
		// {
		// 	name: "test5",
		// 	args: args{matchsticks: []int{14, 2, 7, 13, 6, 14, 10, 13, 9, 8, 18, 10, 12, 3, 14}},
		// 	want: false,
		// },
		// {
		// 	name: "test6",
		// 	args: args{matchsticks: []int{5, 5, 5, 5, 16, 4, 4, 4, 4, 4, 3, 3, 3, 3, 4}},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.args.matchsticks); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
