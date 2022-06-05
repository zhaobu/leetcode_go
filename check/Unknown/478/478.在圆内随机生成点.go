package main

import (
	"math"
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=478 lang=golang
 *
 * [478] 在圆内随机生成点
 */

// @lc code=start
type Solution struct {
	radius           float64
	xCenter, yCenter float64
}

func Constructor(radius float64, xCenter float64, yCenter float64) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{
		radius:  radius,
		xCenter: xCenter,
		yCenter: yCenter,
	}
}

/*
解法1 拒绝采样
*/
func (this *Solution) RandPoint1() []float64 {

	x := rand.Float64()*2 - 1 //[-1, 1)
	y := rand.Float64()*2 - 1 //[-1, 1)
	for x*x+y*y >= 1 {
		x, y = rand.Float64()*2-1, rand.Float64()*2-1
	}
	return []float64{this.xCenter + x*this.radius, this.yCenter + y*this.radius}
}

/*
解法2 数学
*/
func (this *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64())
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{this.xCenter + r*cos*this.radius, this.yCenter + r*sin*this.radius}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, xCenter, yCenter);
 * param_1 := obj.RandPoint();
 */
// @lc code=end
