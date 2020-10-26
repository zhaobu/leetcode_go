package main

import "leetcode/utils/zaplog"

func init() {
	zaplog.InitLog(&zaplog.Config{
		Logdir:   "./",
		LogName:  "main.log",
		Loglevel: "debug",
		Debug:    true},
	)
}

func main() {
	// Test1()
	// Test2()
	// Test3()
	// Test4()
	// Test5()
	Test6()
}
