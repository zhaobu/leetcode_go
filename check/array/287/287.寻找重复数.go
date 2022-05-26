package main

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start

/*
解法1 快慢指针
*/
func findDuplicate(nums []int) int {
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

// @lc code=end
