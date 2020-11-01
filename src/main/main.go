package main

import "leetcode/utils/zaplog"

// func init() {
// 	zaplog.InitLog(&zaplog.Config{
// 		Logdir:   "./",
// 		LogName:  "main.log",
// 		Loglevel: "debug",
// 		Debug:    true},
// 	)
// }

type Student struct {
	Name string
	Age  int
}

func StudentAge1(s Student) *Student {
	if s.Age > 18 {
		s.Name = "aaa"
	}
	return &s
}

func StudentAge2(s *Student) *Student {
	if s.Age > 18 {
		s.Name = "aaa"
	}
	return s
}

var record []int

func fib1(N int) (res int) {
	if N == 1 || N == 2 {
		return 1
	}
	if record[N-1] == 0 {
		record[N-1] = fib1(N-1) + fib1(N-2)
	}
	return record[N-1]
}

func fib2(N int) (res int) {
	if N == 1 || N == 2 {
		return 1
	}
	pre, cur := 1, 1
	for i := 3; i <= N; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}

func main() {
	num := 300
	record = make([]int, num)
	res := fib2(num)
	zaplog.Infof("fib(%d)=%d", num, res)

	// for i := 1; i < 10; i++ {
	// 	res := fib1(i)
	// 	zaplog.Infof("fib1(%d)=%d", i, res)
	// }
}
