package main

/*
 * @lc app=leetcode.cn id=227 lang=golang
 *
 * [227] 基本计算器 II
 */

// @lc code=start

/*
解法1 双栈通用解法(因为此题,字符不包含(),所以更简单)
中缀表达式计算值的常规思路
1. 用2个栈nums和ops分别存储操作数和操作符,注意数字可能大于9,ops中不存储)
2. 如果当前字符是数字, 根据记录的上一个字符的类型,计算出实际的数字,并且入栈nums
3. 如果当前字符是空格,直接跳过
4. 如果当前字符是(,直接入栈ops
5.
5.1 如果当前字符是+,-,*,/等计算符号,如果ops是空,则直接入栈ops
5.2 如果ops不为空,并且ops栈顶是也是计算符号,并且优先级不比当前计算符号优先级低,则
ops出栈1次,nums出栈2次,进行一次计算,计算结果入栈nums.
5.3 重复多次,直到ops栈为空,或者栈顶元素为(,或者优先级小于当前计算符号
比如: 假设我们当前已经扫描到 2 + 1 了（此时栈内的操作为 + ）
如果再遇到一个+,-, 则我们可以进行1次计算
如果再遇到一个/*, 因为优先级更高,不需要计算

6. 当遍历完中缀表达式后,此时ops栈中可能有计算符号,则需要进行多次calOnce,直到ops栈为空
ops中不可能还有(,因为遇到)时已经把(弹出了
*/
func calculate(s string) int {
	nums := []int{} //存储数字
	ops := []byte{} //存储 (,+,-,*,/

	/*
		1. 出栈一个操作数,出栈2个数字,并进行计算,然后计算结果入栈
		2. 调用该函数时,必须保证ops不为空,且栈顶元素为计算符号
	*/
	calOnce := func() {
		// if ops[len(ops)-1] == '(' {
		// 	return
		// }
		ans := 0
		n1, n2 := nums[len(nums)-2], nums[len(nums)-1]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]
		nums = nums[:len(nums)-2]

		if op == '+' {
			ans = n1 + n2
		} else if op == '-' {
			ans = n1 - n2
		} else if op == '*' {
			ans = n1 * n2
		} else {
			ans = n1 / n2
		}
		nums = append(nums, ans)
	}

	getType := func(c byte) int {
		if c == ' ' {
			return 0
		} else if c >= '0' && c <= '9' {
			return 1
		} else if c == '+' || c == '-' {
			return 2
		} else if c == '*' || c == '/' {
			return 3
		} else if c == '(' {
			return 4
		} else {
			return 5
		}
	}

	preType := 0 //记录上一个字符的类型,用来计算大于9的数字
	for i := range s {
		curType := getType(s[i])
		if curType == 1 {
			num := int(s[i] - '0')
			if preType == 1 {
				nums[len(nums)-1] = nums[len(nums)-1]*10 + num
			} else {
				nums = append(nums, num)
			}
		} else if curType == 4 {
			ops = append(ops, s[i])
		} else if curType == 2 || curType == 3 {
			/*
				1. 出栈一个操作符和2个操作数进行计算,直到操作数栈顶的优先级低于当前操作数
				1. 注意此时有可能ops栈顶是(,但是calOnce里面有校验,所以可以不用管
			*/
			for len(ops) > 0 && getType(ops[len(ops)-1]) != 4 && getType(ops[len(ops)-1]) >= curType {
				calOnce()
			}
			ops = append(ops, s[i])
		} else if curType == 5 {
			//出栈一个操作符和2个操作数进行计算,直到到栈顶是'('
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				calOnce()
			}
			ops = ops[:len(ops)-1] //弹出栈顶的(
		}
		preType = curType
	}
	//遍历完中缀表达式后,操作数栈可能不为空,并且操作数栈中不可能含有(
	for len(ops) > 0 {
		calOnce()
	}
	return nums[0]
}

// @lc code=end
