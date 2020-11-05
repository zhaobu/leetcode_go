package Solution

/*
[leetcode官方 ](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/solution/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-de-zui--16/)
*/
type MinStack struct {
	nums    []int
	minNums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		nums:    []int{}, //栈
		minNums: []int{}, //记录入栈时每个元素对应的最小值
	}
}

func (this *MinStack) Push(x int) {
	//元素入栈
	this.nums = append(this.nums, x)
	//如果是第一个元素或者新入的元素x比之前的最小值还小
	if len(this.minNums) == 0 || x < this.minNums[len(this.minNums)-1] {
		this.minNums = append(this.minNums, x)
	} else { //新入的元素x不比之前的最小值小
		this.minNums = append(this.minNums, this.minNums[len(this.minNums)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.nums) < 1 {
		return
	}
	this.nums = this.nums[0 : len(this.nums)-1]
	this.minNums = this.minNums[0 : len(this.minNums)-1]
}

func (this *MinStack) Top() int {
	if len(this.nums) < 1 {
		return -1
	}
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.nums) < 1 {
		return -1
	}
	return this.minNums[len(this.minNums)-1]
}

/*
优化设计[-100000, 100000]限制
https://www.cxyxiaowu.com/2968.html
*/
type MinStack1 struct {
	nums []int
	min  int
}

/** initialize your data structure here. */
func Constructor1() MinStack1 {
	return MinStack1{
		nums: []int{}, //栈
		min:  -1,
	}
}

func (this *MinStack1) Push1(x int) {
	if len(this.nums) == 0 {
		this.min = x
		this.nums = append(this.nums, 0)
	} else {
		// 计算差值
		compare := x - this.min
		this.nums = append(this.nums, compare)
		// 如果差值小于0，显然 x 成为最小值，否则最小值不变
		if compare < 0 {
			this.min = x
		}
	}
}

func (this *MinStack1) Pop1() {
	if len(this.nums) < 1 {
		return
	}
	if this.nums[len(this.nums)-1] < 0 {
		this.min = this.min - this.nums[len(this.nums)-1]
	}
	this.nums = this.nums[0 : len(this.nums)-1]
}

func (this *MinStack1) Top1() int {
	if len(this.nums) < 1 {
		return -1
	}
	return this.nums[len(this.nums)-1] + this.min
}

func (this *MinStack1) GetMin1() int {
	if len(this.nums) < 1 {
		return -1
	}
	return this.min
}
