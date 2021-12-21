package Solution

import (
	"fmt"
	"leetcode/check"
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

// // 解法2:类似动态规划
// func largestRectangleArea2(heights []int) int {
// 	if len(heights) < 1 {
// 		return 0
// 	}
// 	leftLess := make([]int, len(heights))  //记录i左侧且最近的小于其高度的柱子的下标
// 	rightLess := make([]int, len(heights)) //记录i右侧且最近的小于其高度的柱子的下标
// 	leftLess[0] = -1
// 	rightLess[len(heights)-1] = len(heights)

// 	for i := 1; i < len(heights); i++ { //从左往右依次计算leftLess
// 		left := i - 1                                  //从i的左边第一个开始比较
// 		for left >= 0 && heights[left] >= heights[i] { //跳过i的左边所有>=heights[i]的位置
// 			/*
// 			   1. 走到这里说明left位置的高度比i位置的要大,应该找比heights[left]
// 			   更小的位置去和heights[i]比较
// 			   2. 比heights[left]更小的第一个位置就是leftLess[left]
// 			   3. 因为是从左往右计算的leftLess,所以小于i的所有leftLess值已知
// 			*/
// 			left = leftLess[left]
// 		}
// 		leftLess[i] = left
// 	}

// 	for i := len(heights) - 2; i >= 0; i-- { //从右往左依次计算rightLess
// 		right := i + 1 //从i的右边第一个位置开始比较
// 		for right < len(heights) && heights[right] >= heights[i] {
// 			right = rightLess[right]
// 		}
// 		rightLess[i] = right
// 	}

// 	//计算最大面积
// 	maxArea := 0
// 	for i, height := range heights {
// 		curArea := (rightLess[i] - leftLess[i] - 1) * height
// 		if curArea > maxArea {
// 			maxArea = curArea
// 		}
// 	}
// 	return maxArea
// }

// //单调递增栈2
// func largestRectangleArea3(heights []int) int {
// 	if len(heights) < 1 {
// 		return 0
// 	}
// 	ascStack := make([]int, 0, len(heights)) //单调递增栈,栈中存储的是heights的下标,使用时通过heights来获取高度
// 	maxArea := 0                             //记录最大面积
// 	heights = append(heights, 0)             //在heights最后插入一个比谁都小的0,从而在0入栈后会计算完所有的情况
// 	right := 0                               //计算maxArea时右边界
// 	left := 0                                //计算maxArea时左边界

// 	for i, height := range heights {
// 		// 如果新的元素较小，那就一直把栈内元素弹出来，直到栈顶比新元素小
// 		for len(ascStack) > 0 && heights[ascStack[len(ascStack)-1]] >= height {
// 			stackTop := ascStack[len(ascStack)-1] //(栈顶)出栈元素
// 			right = i                             //当元素出栈时，说明这个新元素是出栈元素向后找第一个比其小的元素
// 			ascStack = ascStack[:len(ascStack)-1] // 出栈
// 			if len(ascStack) > 0 {                //当元素出栈后，说明新栈顶元素是出栈元素向前找第一个比其小的元素
// 				left = ascStack[len(ascStack)-1]
// 			} else {
// 				left = -1
// 			}
// 			//计算以heights[stackTop]为高度的面积
// 			curArea := (right - left - 1) * heights[stackTop]
// 			if curArea > maxArea {
// 				maxArea = curArea
// 			}
// 		}
// 		ascStack = append(ascStack, i) //如果新的元素比栈顶元素大，就入栈
// 	}
// 	return maxArea
// }

// //单调递增栈
// func largestRectangleArea(heights []int) int {
// 	if len(heights) < 1 {
// 		return 0
// 	}
// 	leftLess := make([]int, len(heights))    //记录i左侧且最近的小于其高度的柱子的下标
// 	rightLess := make([]int, len(heights))   //记录i右侧且最近的小于其高度的柱子的下标
// 	ascStack := make([]int, 0, len(heights)) //单调递增栈,栈中存储的是heights的下标,使用时通过heights来获取高度
// 	for i, height := range heights {
// 		// 如果新的元素较小，那就一直把栈内元素弹出来，直到栈顶比新元素小
// 		for len(ascStack) > 0 && heights[ascStack[len(ascStack)-1]] >= height {
// 			stackTop := ascStack[len(ascStack)-1] //(栈顶)出栈元素
// 			rightLess[stackTop] = i               // 新元素是出栈元素向后找第一个比其小的元素
// 			ascStack = ascStack[:len(ascStack)-1] // 出栈
// 		}
// 		//当栈顶元素小于当前元素时,栈顶元素就是当前元素的左边第一个比其小的元素
// 		if len(ascStack) > 0 {
// 			leftLess[i] = ascStack[len(ascStack)-1]
// 		} else {
// 			leftLess[i] = -1
// 		}
// 		ascStack = append(ascStack, i) //如果新的元素比栈顶元素大，就入栈
// 	}
// 	//栈里面剩余的元素需要全部弹出,因为他们还没有求rightLess,
// 	for len(ascStack) > 0 {
// 		stackTop := ascStack[len(ascStack)-1] //(栈顶)出栈元素
// 		rightLess[stackTop] = len(heights)    // 新元素是出栈元素向后找第一个比其小的元素
// 		ascStack = ascStack[:len(ascStack)-1] // 出栈
// 	}
// 	//计算最大面积
// 	maxArea := 0
// 	for i, height := range heights {
// 		curArea := (rightLess[i] - leftLess[i] - 1) * height
// 		if curArea > maxArea {
// 			maxArea = curArea
// 		}
// 	}
// 	return maxArea
// }

func (s *LargestRectangleArea) TestlargestRectangleArea() {
	fmt.Printf("测试执行中 ********* TestlargestRectangleArea\n")
	heights := [][]int{
		{6, 7, 5, 2, 4, 5, 9, 3},
		// {2, 1, 5, 6, 2, 3},
		// {2, 4},
	}
	for _, v := range heights {
		s.Equal(check.LargestRectangleArea1(v), check.LargestRectangleArea2(v), "%s和%s比较:数组的值为:%+v", "LargestRectangleArea1", "LargestRectangleArea2", v)
		s.Equal(check.LargestRectangleArea1(v), check.LargestRectangleArea3(v), "%s和%s比较:数组的值为:%+v", "LargestRectangleArea1", "LargestRectangleArea3", v)
		s.Equal(check.LargestRectangleArea1(v), check.LargestRectangleArea(v), "%s和%s比较:数组的值为:%+v", "LargestRectangleArea1", "LargestRectangleArea", v)
	}
}

func TestSuit(t *testing.T) {
	//将运行MyTestSuit中所有名为TestXxx的方法
	suite.Run(t, new(LargestRectangleArea))
}
