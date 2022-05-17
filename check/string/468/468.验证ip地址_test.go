package main

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		queryIP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				queryIP: "172.16.254.1",
			},
			want: "IPv4",
		},
		{
			name: "test2",
			args: args{
				queryIP: "2001:0db8:85a3:0:0:8A2E:0370:7334",
			},
			want: "IPv6",
		},
		{
			name: "test3",
			args: args{
				queryIP: "2001:0db8:85a3:0:0:8A2E:0370:7334:",
			},
			want: "Neither",
		},
		{
			name: "test4",
			args: args{
				queryIP: "12.12.12",
			},
			want: "Neither",
		},
		{
			name: "test5",
			args: args{
				queryIP: "192.0.0.1",
			},
			want: "IPv4",
		},
		{
			name: "test6",
			args: args{
				queryIP: "01.01.01.01",
			},
			want: "Neither",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.queryIP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
