package main

import (
	"fmt"
	"leetcode/utils/zaplog"
)

/*
[腾讯笔试题](https://www.yuque.com/docs/share/576ea5c7-e0e8-4654-be08-fb80196a8cc6?# 《腾讯》)
*/
func Test1() {
	// PrintN(3)
	for i := 3; i < 10; i++ {
		PrintN(i)
	}
}

func PrintN(n int) [][]int {
	zaplog.Infof("开始输出%d*%d", n, n)
	//构造n*n的二维数组
	nums := make([][]int, n)
	for i, _ := range nums {
		nums[i] = make([]int, n)
	}

	i, j := n-1, n-1       //当前输出的索引
	w := 0                 //0表示向上,1表示向左下,2表示向右
	max := (1 + n) * n / 2 //最大数字
	count := 1             //记录当前输出到的数字
	for count <= max {
		nums[i][j] = count
		count++
		if w == 0 {
			if i-1 >= 0 && nums[i-1][j] == 0 { //继续往上走
				i--
			} else { //往上走到最上
				i++
				j--
				w = 1
			}
		} else if w == 1 {
			if i+1 < n && j-1 >= 0 && nums[i+1][j-1] == 0 { //继续往左下走
				i++
				j--
			} else { //走到最左下
				j++
				w = 2
			}
		} else if w == 2 {
			if j+1 < n && nums[i][j+1] == 0 { //继续往右走
				j++
			} else { //走到最右边
				w = 0
				i--
			}
		}
	}
	str := "\n"
	for i, v := range nums {
		str += fmt.Sprintf("nums[%d]:%v\n", i, v)
	}
	zaplog.Infof("%s", str)
	zaplog.Infof("结束输出%d*%d", n, n)
	return nums
}

func getres(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	direction := 0
	idx := 1
	left := 0
	right := n - 1
	low := 0
	high := n - 1
	final := n * (n + 1) / 2
	fmt.Println(final)
	for idx <= final {
		switch direction % 3 {
		case 0: //下->上
			for i := high; i >= low; i-- {
				matrix[i][right] = idx
				idx += 1
			}
			right -= 1
			low += 1
		case 1: //右上->左下
			for i, j := low, right; i <= high && j >= left; i, j = i+1, j-1 {
				matrix[i][j] = idx
				idx += 1
			}
			low += 1
			left += 1
		case 2: //左->右
			for j := left; j <= right; j++ {
				matrix[high][j] = idx
				idx += 1
			}
			high -= 1
			left += 1
		}
		direction += 1
	}
	return matrix
}
