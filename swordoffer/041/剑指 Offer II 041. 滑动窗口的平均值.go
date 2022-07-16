package main

/*
 *
 *
 * 剑指 Offer II 041. 滑动窗口的平均值
 */

// @lc code=start

/*
 解法1
*/
type MovingAverage struct {
	arr      []int
	cap, sum int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		arr: make([]int, 0, size),
		cap: size,
		sum: 0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.arr) < this.cap {
		this.arr = append(this.arr, val)
		this.sum += val
	} else {
		this.sum -= this.arr[0]
		this.sum += val
		this.arr = append(this.arr[1:], val)
	}
	//  fmt.Printf("sum=%d, len=%d \n", this.sum, this.len)
	return float64(this.sum) / float64(len(this.arr))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

// @lc code=end
