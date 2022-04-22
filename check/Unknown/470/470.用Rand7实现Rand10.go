package main /*
 * @lc app=leetcode.cn id=470 lang=golang
 *
 * [470] 用 Rand7() 实现 Rand10()
 */

// @lc code=start

/*
解法1
https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/mo-neng-gou-zao-fa-du-li-sui-ji-shi-jian-9xpz/

1. 先拒绝采样取得[1,6]之间的等概率随机数,
2. 再拒绝采样取得[1,5]之间的等概率随机数
3. 第1步判断奇偶,又可以把[1,6]划分为{1,3,5}和{2,4,6}的0.5的概率
4. 然后利用第3步的0.5概率,结合第2步的[1,5]就可以求出一个[1,10]之间等概率的随机数求解方法
*/
func rand101() int {
	first, second := 0, 0
	/*
		第一个数取[1,6]
	*/
	for {
		first = rand7()
		if first <= 6 {
			break
		}
	}

	/*
		第二个数取[1,5]
	*/
	for {
		second = rand7()
		if second <= 5 {
			break
		}
	}
	//然后再判断first的奇偶
	if first&1 == 0 {
		second += 5
	}
	return second
}

/*
解法2
https://leetcode-cn.com/problems/implement-rand10-using-rand7/solution/cong-zui-ji-chu-de-jiang-qi-ru-he-zuo-dao-jun-yun-/

前置知识点1: 如何通过小范围求大范围
已知 rand_N() 可以等概率的生成[1, N]范围的随机数
那么：
(rand_X() - 1) × Y + rand_Y() ==> 可以等概率的生成[1, X * Y]范围的随机数
即实现了 rand_XY()

前置知识点2: 如何通过大范围求小范围
已知rand4()会均匀产生[1,4]的随机数,如何求rand2()
rand4() % 2 + 1 = ?
   1 % 2    + 1 = 2
   2 % 2    + 1 = 1
   3 % 2    + 1 = 2
   4 % 2    + 1 = 1
只要rand_N()中N是2的倍数，就都可以用来实现rand2()，反之，若N不是2的倍数，则产生的结果不是等概率的。比如：
rand6() % 2 + 1 = ?
   1 % 2    + 1 = 2
   2 % 2    + 1 = 1
   3 % 2    + 1 = 2
   4 % 2    + 1 = 1
   5 % 2    + 1 = 2
   6 % 2    + 1 = 1

rand5() % 2 + 1 = ?
   1 % 2    + 1 = 2
   2 % 2    + 1 = 1
   3 % 2    + 1 = 2
   4 % 2    + 1 = 1
   5 % 2    + 1 = 2
*/
func rand102() int {
	for {

		num := (rand7()-1)*7 + rand7() //rand49
		if num <= 40 {                 //拒绝采样[1,40]
			return num%10 + 1
		}
		num -= 40                 //rand9
		num = (num-1)*7 + rand7() //rand63

		if num <= 60 { //拒绝采样[1,60]
			return num%10 + 1
		}

		num -= 60                 //rand3
		num = (num-1)*7 + rand7() //rand21
		if num <= 20 {
			return num%10 + 1
		}
	}
}

/*
解法3 解法2的改写
*/
func rand10() int {
	for {
		first := rand7() //第一个随机数
		firstRand := 7   //第一个随机数的范围
		for firstRand > 1 {
			num := (first-1)*7 + rand7()   //计算随机数
			curRand := firstRand * 7       //本轮随机数的范围[1,curRand]
			maxRand := (curRand / 10) * 10 //计算拒绝采样的范围
			if num <= maxRand {            //拒绝采样[1,maxRand]
				return num%10 + 1
			}
			first = num - maxRand         //计算下一轮的第一个随机数
			firstRand = curRand - maxRand //计算下一轮第一个随机数的范围
		}
	}
}

// @lc code=end
