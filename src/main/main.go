package main

import (
	"errors"
	"fmt"
)

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

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}
