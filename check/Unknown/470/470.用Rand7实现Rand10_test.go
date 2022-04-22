package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_rand10(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{
			name: "test1",
			n:    10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := map[int]int{}
			for i := 0; i < tt.n; i++ {
				num := rand10()
				record[num] += 1
				// t.Logf("i=%d, rand10() = %v", i, num)
			}
			t.Logf("record=%+v", record)
		})
	}
}
func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(8-1) + 1 //[0,7)+1 = [1,8)
}
