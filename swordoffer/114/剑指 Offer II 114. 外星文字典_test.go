package main

import (
	"testing"
)

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test3",
			args: args{
				words: []string{"wrt", "wrf", "er", "ett", "rftt"},
			},
			want: "wertf",
		},
		// {
		// 	name: "test4",
		// 	args: args{
		// 		words: []string{"z", "z"},
		// 	},
		// 	want: "z",
		// },
		// {
		// 	name: "test5",
		// 	args: args{
		// 		words: []string{"ab", "adc"},
		// 	},
		// 	want: "abcd",
		// },
		{
			name: "test6",
			args: args{
				words: []string{"wrt", "wrtkj"},
			},
			want: "jkrtw",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
