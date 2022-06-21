package main

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=497 lang=golang
 *
 * [497] 非重叠矩形中的随机点
 */

// @lc code=start
type Solution struct {
	rects [][]int
}

func Constructor(rects [][]int) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		rects: rects,
	}
}

func (this *Solution) Pick() []int {
	//先随机矩形下标
	index := rand.Intn(len(this.rects))
	//随机x
	rects := this.rects[index]
	xWidth := rects[2] - rects[0] + 1
	yWidth := rects[3] - rects[1] + 1

	return []int{
		rand.Intn(xWidth) + rects[0],
		rand.Intn(yWidth) + rects[1],
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
// @lc code=end
