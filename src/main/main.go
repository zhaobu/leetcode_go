package main

// func init() {
// 	zaplog.InitLog(&zaplog.Config{
// 		Logdir:   "./",
// 		LogName:  "main.log",
// 		Loglevel: "debug",
// 		Debug:    true},
// 	)
// }

type A struct {
	s string
}

// 这是上面提到的 "在方法内把局部变量指针返回" 的情况
func foo(s string) bool {
	a := new(A)
	a.s = s
	return true //返回局部变量a,在C语言中妥妥野指针，但在go则ok，但a会逃逸到堆
}

func main() {
	foo("a")
}
