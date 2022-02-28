package main

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */

// @lc code=start

/*
解法1
1. 两个单词，设为 A 和 B，就有六种操作方法, 但是有些操作方式是等价的.
对单词 A 删除一个字符和对单词 B 插入一个字符是等价的
对单词 B 删除一个字符和对单词 A 插入一个字符也是等价的
对单词 A 替换一个字符和对单词 B 替换一个字符是等价的
所以实际上本质不同的操作共有3种:
+ 在单词 A 中插入一个字符；
+ 在单词 B 中插入一个字符；
+ 修改单词 A 的一个字符。

对于这3种操作,可以分为4种情况:
1. 在单词 A 中插入一个字符:
先删除 word1[0,i) 的最后一个字符得到 word1[0,i-1) ,
然后由 word1[0,i-1) 转换为 word2[0,j).
dp[i][j] = 1 + dp[i-1][j]

2. 在单词 B 中插入一个字符:
先由 word1[0,i) 转换为 word2[0,j-1),
然后在最后插入字符word2[j-1],得到word2[0,j)
dp[i][j] = 1 + dp[i][j-1]

3. 修改单词 A 的一个字符:
3.1 如果word1[i-1] == word2[j-1]
由word1[0,i-1)转换为word2[0,j-1)后不用再做任何操作
dp[i][j] := dp[i-1][j-1]
3.2 如果word1[i-1] != word2[j-1]
先由word1[0,i-1)转换为word2[0,j-1)然后再将word1[i-1]替换为word2[j-1]
dp[i][j] := dp[i-1][j-1] + 1

所以dp[i][j]求的就是上面3种情况中的最小值
*/
func minDistance1(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 < 1 {
		return n2
	}
	if n2 < 1 {
		return n1
	}
	/*
	   dp[i][j] 表示 word1[0,i) 和 word2[0,j) 的编辑距离
	*/
	dp := make([][]int, n1+1)
	// fmt.Printf("0, dp=%+v\n", dp)

	//初始化二位数组的每一行,并且初始化dp[i][0],也就是从word1到空串word2的编辑距离
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][0] = i
	}
	// fmt.Printf("1, dp=%+v\n", dp)

	//初始化dp[0][j],也就是从空串word1到word2的编辑距离
	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}
	// fmt.Printf("2, dp=%+v\n", dp)

	var min func(a, b, c int) int
	min = func(a, b, c int) int {
		if a > b {
			if b > c {
				return c
			}
			return b
		}
		if a > c {
			return c
		}
		return a
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			/*
			   先删除 word1[0,i) 的最后一个字符得到 word1[0,i-1) ,
			   然后由 word1[0,i-1) 转换为 word2[0,j).
			*/
			dp1 := dp[i-1][j] + 1

			/*
			   先由word1[0,i)转换为word2[0,j-1),
			   然后在最后插入字符word2[j-1],得到word2[0,j)
			*/
			dp2 := dp[i][j-1] + 1

			/*
			   如果word1[i-1] == word2[j-1],由word1[0,i-1)转换为
			   word2[0,j-1)后不用再做任何操作
			*/
			dp3 := dp[i-1][j-1]
			/*
			   如果word1[i-1] != word2[j-1],先由word1[0,i-1)转换为word2[0,j-1)
			   然后再将word1[i-1]替换为word2[j-1]
			*/
			if word1[i-1] != word2[j-1] {
				dp3 += 1
			}
			dp[i][j] = min(dp1, dp2, dp3)
		}
	}

	// fmt.Printf("3, dp=%+v\n", dp)
	return dp[n1][n2]
}

/*
解法2
递归,该方法超时,但是可以辅助理解动态规划
if s1[i] == s2[j]:
    啥都别做（skip）
    i, j 同时向前移动
else:
    三选一：
        插入（insert）
        删除（delete）
        替换（replace）
*/
func minDistance2(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 < 1 {
		return n2
	}
	if n2 < 1 {
		return n1
	}

	var min func(a, b, c int) int
	min = func(a, b, c int) int {
		if a > b {
			if b > c {
				return c
			}
			return b
		}
		if a > c {
			return c
		}
		return a
	}

	s1, s2 := []byte(word1), []byte(word2)

	// 定义：dp(i, j) 返回 s1[0..i] 和 s2[0..j] 的最小编辑距离
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 { //当 n1=0 时,i=-1,此时表示word1是空串
			return j + 1
		}
		if j == -1 { //当 n2=0 时,j=-1,此时表示word2是空串
			return i + 1
		}

		/*
			本来就相等，不需要任何操作
			s1[0..i] 和 s2[0..j] 的最小编辑距离等于
			s1[0..i-1] 和 s2[0..j-1] 的最小编辑距离
			也就是说 dp(i, j) 等于 dp(i-1, j-1)
		*/
		if s1[i] == s2[j] {
			return dp(i-1, j-1)
		}

		insertAction := dp(i, j-1) + 1    //直接在 s1[i] 插入一个和 s2[j] 一样的字符,那么 s2[j] 就被匹配了，前移 j，继续跟 i 对比,操作数加一
		deleteAction := dp(i-1, j) + 1    //直接把 s[i] 这个字符删掉. 前移 i，继续跟 j 对比, 操作数加一
		replaceAction := dp(i-1, j-1) + 1 //直接把 s1[i] 替换成 s2[j]，这样它俩就匹配了. 同时前移 i，j 继续对比 , 操作数加一

		return min(insertAction, deleteAction, replaceAction)
	}
	return dp(n1-1, n2-1)
}

/*
解法3
动态规划,相当于用 DP table 优化递归,唯一不同的是，DP table 是自底向上求解，递归解法是自顶向下求解：
*/

func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	if n1 < 1 {
		return n2
	}
	if n2 < 1 {
		return n1
	}

	var min func(a, b, c int) int
	min = func(a, b, c int) int {
		if a > b {
			if b > c {
				return c
			}
			return b
		}
		if a > c {
			return c
		}
		return a
	}

	// 定义：s1[0..i] 和 s2[0..j] 的最小编辑距离是 dp[i-1][j-1]
	dp := make([][]int, n1+1)

	//初始化dp[i][0]
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
		dp[i][0] = i
	}
	// fmt.Printf("0, dp=%+v\n", dp)

	for j := 0; j <= n2; j++ {
		dp[0][j] = j
	}
	// fmt.Printf("1, dp=%+v\n", dp)

	//自底向上求解
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i][j-1]+1,   //插入
					dp[i-1][j]+1,   //删除
					dp[i-1][j-1]+1, //替换
				)
			}
		}
	}

	return dp[n1][n2]
}

// @lc code=end
