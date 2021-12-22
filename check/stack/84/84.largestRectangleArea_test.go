package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// func init() {
// 	zaplog.InitLog(&zaplog.Config{
// 		Logdir:   "./",
// 		LogName:  "main.log",
// 		Loglevel: "debug",
// 		Debug:    true},
// 	)
// }

type LargestRectangleArea struct {
	suite.Suite
}

//实现了SetupAllSuite接口,在套件中所有测试开始运行前调用这个方法,对应的是TearDownAllSuite
func (s *LargestRectangleArea) SetupSuite() {
	fmt.Printf("所有测试开始运行前调用SetupSuite这个方法. \n")
}

//实现了TearDownSuite接口,在套件中所有测试运行完成后调用这个方法
func (s *LargestRectangleArea) TearDownSuite() {
	fmt.Printf("所有测试运行完成后调用TearDownSuite这个方法. \n")
}

//在每个测试运行前调用，接受套件名和测试名作为参数
func (s *LargestRectangleArea) BeforeTest(suiteName, testName string) {
	fmt.Printf("在每个测试运行前调用BeforeTest，接受套件名和测试名作为参数. suite:%s test:%s\n", suiteName, testName)
}

//在每个测试运行后调用，接受套件名和测试名作为参数
func (s *LargestRectangleArea) AfterTest(suiteName, testName string) {
	fmt.Printf("在每个测试运行后调用AfterTest，接受套件名和测试名作为参数. suite:%s test:%s\n", suiteName, testName)
}

func (s *LargestRectangleArea) TestlargestRectangleArea() {
	fmt.Printf("测试执行中 ********* TestlargestRectangleArea\n")
	heights := [][]int{
		{6, 7, 5, 2, 4, 5, 9, 3},
		{2, 1, 5, 6, 2, 3},
		{2, 4},
	}
	for _, v := range heights {
		s.Equal(largestRectangleArea1(v), largestRectangleArea2(v), "%s和%s比较:数组的值为:%+v", "largestRectangleArea1", "largestRectangleArea2", v)
		s.Equal(largestRectangleArea1(v), largestRectangleArea3(v), "%s和%s比较:数组的值为:%+v", "largestRectangleArea1", "largestRectangleArea3", v)
		s.Equal(largestRectangleArea1(v), largestRectangleArea(v), "%s和%s比较:数组的值为:%+v", "largestRectangleArea1", "LargestRectangleArea", v)
	}
}

func TestSuit(t *testing.T) {
	//将运行MyTestSuit中所有名为TestXxx的方法
	suite.Run(t, new(LargestRectangleArea))
}
