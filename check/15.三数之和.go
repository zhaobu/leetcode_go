/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (34.00%)
 * Likes:    4078
 * Dislikes: 0
 * Total Accepted:    729.9K
 * Total Submissions: 2.1M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0
 * 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：nums = [0]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0
 * -10^5
 *
 *
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 5)
	if nums == nil || len(nums) < 3 {
		return res
	}
	//先对数组进行排序
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; { //依次找到2个数之和为-c的2个元素
		i, j := k+1, len(nums)-1
		for i < j {
			sum := nums[k] + nums[i] + nums[j]
			if sum == 0 {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				// i++和j--必须至少执行一次
				i++
				j--
				//跳过所有和nums[i相等的]
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				//跳过所有和nums[j相等的]
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else if sum < 0 {
				i++ //i++必须至少执行一次
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else {
				j-- // j--必须至少执行一次
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
		k++
		/*
			k++必须至少执行一次,如果不在这里写k++,那也可以
			放在最外层for循环,不过下面的for循环就不一样
		*/
		for k < len(nums)-2 && nums[k] == nums[k-1] {
			k++
		}
	}
	return res
}

// @lc code=end

