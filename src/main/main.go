package main

import "fmt"

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

func main() {
	arr := []int{1, 2, 3}
	for _, v := range arr {
		arr = append(arr, v)
	}
	fmt.Println(arr)
}
