package main

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start

/*
解法1 动态规划
dp[i]表示以s[i]结尾时最长有效括号子串的长度
每个字符共有2种情况,每个字符和前一个字符共有4种情况
1: ((
2: )(
以(结尾的2种情况不用考虑,dp[i]=0
3: ()
直接dp[i]=dp[i-2]+2
4: ))
直接看dp[i-1]
如果dp[i-1]=0,则dp[i]=0
如果dp[i-1]>0,则要看s[i-dp[i-1]-1]是否为(
	如果不是说明不能和)配对,dp[i]=0
	如果是,则已经能得到的最长有效括号子串的长度为dp[i-1]+2,
	另外还要加上(之前的部分,还可能会有有效的子串,长度为dp[i-2-dp[i-1]]
	所以综合来看,就是dp[i]=dp[i-1]+2+dp[i-2-dp[i-1]]
*/
func longestValidParentheses1(s string) int {
	/*
	   "(("
	   ")("
	   "()()"
	   "(())"
	*/
	m := len(s)
	if m < 2 {
		return 0
	}

	//dp[i]表示以s[i]结尾时最长有效括号子串的长度
	dp := make([]int, m)

	ret := 0
	for i := 1; i < m; i++ {
		dp[i] = 0
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = 2     //至少已经有了一对()
				if i-2 >= 0 { //加上i-2前面的部分
					dp[i] += dp[i-2]
				}
			} else { //s[i-1] == ')'
				if dp[i-1] > 0 { //s[i-1]能配对才考虑s[i]
					left := i - dp[i-1] - 1
					if left >= 0 && s[left] == '(' {
						dp[i] = dp[i-1] + 2
						if left-1 >= 0 {
							dp[i] += dp[left-1]
						}
					}
				}
			}
			if dp[i] > ret {
				ret = dp[i]
			}
		}
	}

	return ret
}

/*
解法2 栈
*/
func longestValidParentheses2(s string) int {
	/*
	   "(("
	   ")("
	   "()()"
	   "(())"
	*/
	m := len(s)
	if m < 2 {
		return 0
	}
	ret := 0
	stack := make([]int, 0, m) //栈底表示最后一个没有被匹配的右括号的下标
	stack = append(stack, -1)  //初始化时入栈1个-1,保持栈底定义的统一
	for i := 0; i < m; i++ {
		if s[i] == '(' { //左括号直接入栈
			stack = append(stack, i)
		} else { //s[i] == ')'
			stack = stack[:len(stack)-1] //弹出栈顶元素,与当前右括号匹配
			if len(stack) == 0 {
				/*
					1. 说明出栈前只剩下栈底元素
					2. 栈中最多只有一个')'的下标
					3. 遇到的s[i]也是')',栈顶也是')',所以应该更新')'的下标
				*/
				stack = append(stack, i)
			} else {
				/*
					1. 遇到的s[i]也是')',栈顶是'(',所以弹出栈顶后,
					i-stack[len(stack)-1]就是以s[i]结尾的最长匹配字符串
				*/
				top := stack[len(stack)-1]
				cur := i - top
				if cur > ret {
					ret = cur
				}
			}
		}
	}

	return ret
}

/*
解法3 双指针
1. 从左向右遍历s用left,right分别统计'('和')'的个数
如果left和right相等,left+right就是当前最长的有效子串长度
如果right>left,left,right就都置零
1. 从右向左遍历s用left,right分别统计'('和')'的个数
如果left和right相等,left+right就是当前最长的有效子串长度
如果left>right,left,right就都置零

两边遍历后,就能求出最长有效字符串的长度
*/
func longestValidParentheses(s string) int {
	/*
	   "(("
	   ")("
	   "()()"
	   "(())"
	*/
	m := len(s)
	if m < 2 {
		return 0
	}

	ret := 0
	left, right := 0, 0

	for i := 0; i < m; i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if left == right {
			cur := left + right
			if cur > ret {
				ret = cur
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := m - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if left == right {
			cur := left + right
			if cur > ret {
				ret = cur
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return ret
}

// @lc code=end
