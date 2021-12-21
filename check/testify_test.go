package check

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ArraryTestSuit struct {
	suite.Suite
	testCount int
}

//实现了SetupAllSuite接口,在套件中所有测试开始运行前调用这个方法,对应的是TearDownAllSuite
func (s *ArraryTestSuit) SetupSuite() {
	fmt.Printf("所有测试开始运行前调用SetupSuite这个方法. \n")
}

//实现了TearDownSuite接口,在套件中所有测试运行完成后调用这个方法
func (s *ArraryTestSuit) TearDownSuite() {
	fmt.Printf("所有测试运行完成后调用TearDownSuite这个方法. \n")
}

//实现了SetupTestSuite接口,在套件中每个测试执行前都会调用这个方法,对应的是TearDownTestSuite
func (s *ArraryTestSuit) SetupTest() {
	fmt.Printf("每个测试执行前都会调用SetupTest这个方法. test count:%d\n", s.testCount)
}

//实现了TearDownTest接口,在套件中每个测试执行后都会调用这个方法
func (s *ArraryTestSuit) TearDownTest() {
	s.testCount++
	fmt.Printf("每个测试执行后都会调用TearDownTest这个方法. test count:%d\n", s.testCount)
}

//在每个测试运行前调用，接受套件名和测试名作为参数
func (s *ArraryTestSuit) BeforeTest(suiteName, testName string) {
	fmt.Printf("在每个测试运行前调用BeforeTest，接受套件名和测试名作为参数. suite:%s test:%s\n", suiteName, testName)
}

//在每个测试运行后调用，接受套件名和测试名作为参数
func (s *ArraryTestSuit) AfterTest(suiteName, testName string) {
	fmt.Printf("在每个测试运行后调用AfterTest，接受套件名和测试名作为参数. suite:%s test:%s\n", suiteName, testName)
}

func (s *ArraryTestSuit) TestExample1() {
	fmt.Printf("测试执行中 ********* TestExample1\n")
}

func (s *ArraryTestSuit) TestExample2() {
	fmt.Printf("测试执行中 ********* TestExample2\n")
}

func TestSuit(t *testing.T) {
	//将运行MyTestSuit中所有名为TestXxx的方法
	suite.Run(t, new(ArraryTestSuit))
}
