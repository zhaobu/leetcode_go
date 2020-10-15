package Solution

/*
 */
func LengthOfLastWord1(s string) int {
	n := len(s)
	if 0 == n {
		return 0
	}
	//从后往前遍历j2,j1分别表示找到的单词末尾和单词的前面的位置
	j1, j2 := -1, -1
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			//如果j2 == -1表示找打了单词开头,标记开头并跳出循环
			if j2 != -1 {
				j1 = i
				break
			}
			continue
		}
		//找到第一个空格,标记末尾
		if j2 == -1 {
			j2 = i
		}
	}
	return j2 - j1
}

func LengthOfLastWord2(s string) int {
	n := len(s)
	if 0 == n {
		return 0
	}
	count := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count > 0 {
				break
			}
			continue
		}
		count++
	}
	return count
}
