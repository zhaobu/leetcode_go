package main

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start

/*
解法1 快慢指针
1. 看成是链表,数组下标i是当前节点,nums[i]是下一个节点
如果成环,按照这种方式遍历链表会无限循环
2. 可以考虑用141. 环形链表的解法来找出环的入口
3. 当存在多个重复值,也就是多个环时,只会在最小环中无限循环,
也就是重复元素第2次出现的位置形成的环.不会走到后面更大的环
*/
func findDuplicate1(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	// fmt.Printf("slow=%d,fast=%d\n", slow, fast)
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
		// fmt.Printf("slow=%d,fast=%d\n", slow, fast)
	}
	// fmt.Printf("slow=%d,fast=%d\n", slow, fast)
	help := 0
	for help != slow {
		// fmt.Printf("slow=%d,help=%d\n", slow, help)
		help = nums[help]
		slow = nums[slow]
	}

	return slow
}

/*
解法2 二分查找
cnt[i]表示nums数组中<=i的个数,假设重复元素为target,target的取值范围为[1,n].
重复元素相当于用target替换了其他元素.
假如把整个数组排序,整个数组nusm总共有n+1个数,可以把数分为三部分
1. < target的在target的左边
2. > target的在target的右边
3. = target的在中间部分
比如:
[1,2,3,4,4,5,6,7] nums
[0,1,2,3,4,5,6,7] 下标
[0,1,2,3,5,6,7,8] cnt[i]

可以观察cnt[i]和下标的值得出结论:
当 i < target 时,cnt[i]=i
当 i >= target 时,cnt[i]>i

另一种情况:
[2,2,2,3] nums
[0,1,2,3] 下标
[0,0,3,4] cnt[i]

可以观察cnt[i]和下标的值得出结论:
当 i < target 时,cnt[i]<=i
当 i >= target 时,cnt[i]>i
所以可以二分查找第一个cnt[i]>i的下标
二分查找的初始范围当然就是[1,n]
*/

func findDuplicate2(nums []int) int {
	left, right := 1, len(nums)-1

	getCnt := func(n int) (cnt int) {
		for i := range nums {
			if nums[i] <= n {
				cnt++
			}
		}
		return cnt
	}

	for left <= right {
		mid := left + (right-left)>>1
		if getCnt(mid) <= mid { //cnt[i]<=i
			left = mid + 1
		} else { //cnt[i]>i
			/*
				1. 当cnt[i]>i时,说明target还有可能在i的左边
				2. 当i已经是最左边或者cnt[i-1]已经不能满足cnt[i-1]>i-1而是
				cnt[i-1]<=i-1时,说明当前的i就是要求的target
			*/
			if mid == 1 || getCnt(mid-1) <= mid-1 {
				return mid
			}
			right = mid - 1
		}
	}
	return -1
}

/*
解法3 位运算
*/
func findDuplicate(nums []int) int {
	n := len(nums) - 1
	ret := int32(0) //一定要用int32位数字,因为后面要进行移位操作
	for i := 0; i < 32; i++ {
		//计算[1,n]所有数字第i位为1的个数
		cntN := 0
		for j := 1; j <= n; j++ {
			if (1<<i)&j > 0 {
				cntN++
			}
		}
		//计算nums数组中所有数字第i位为1的个数
		cntNum := 0
		for j := range nums {
			if (1<<i)&nums[j] > 0 {
				cntNum++
			}
		}
		// fmt.Printf("cntNum=%d,cntN=%d\n", cntNum, cntN)
		//重复元素会让二进制数第i位上为1的个数变多
		if cntNum > cntN {
			ret |= (1 << int32(i))
		}
	}

	return int(ret)
}

// @lc code=end
