package main

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {
		// 	name: "test1",
		// 	args: args{board: [][]byte{
		// 		{'A', 'B', 'C', 'E'},
		// 		{'S', 'F', 'C', 'S'},
		// 		{'A', 'D', 'E', 'E'},
		// 	},
		// 		word: "ABCCED",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test2",
		// 	args: args{board: [][]byte{
		// 		{'A', 'B', 'C', 'E'},
		// 		{'S', 'F', 'C', 'S'},
		// 		{'A', 'D', 'E', 'E'},
		// 	},
		// 		word: "SEE",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test3",
		// 	args: args{board: [][]byte{
		// 		{'A', 'B', 'C', 'E'},
		// 		{'S', 'F', 'C', 'S'},
		// 		{'A', 'D', 'E', 'E'},
		// 	},
		// 		word: "ABCB",
		// 	},
		// 	want: false,
		// },
		// {
		// 	name: "test4",
		// 	args: args{board: [][]byte{
		// 		{'a'},
		// 	},
		// 		word: "a",
		// 	},
		// 	want: true,
		// },
		// {
		// 	name: "test5",
		// 	args: args{board: [][]byte{
		// 		{'C', 'A', 'A'},
		// 		{'A', 'A', 'A'},
		// 		{'B', 'C', 'D'},
		// 	},
		// 		word: "AAB",
		// 	},
		// 	want: true,
		// },
		{
			name: "test6",
			args: args{board: [][]byte{
				{'a', 'z'},
			},
				word: "ba",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
