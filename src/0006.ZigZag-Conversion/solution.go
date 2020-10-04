package Solution

/*
//[leetcode官方 方法一：按行排序]
*/
func Convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n := min(len(s), numRows)
	rows := make([][]byte, n)
	dwon := false //true表示向下,false表示向上
	for i, j := 0, 0; i < len(s); i++ {
		rows[j] = append(rows[j], s[i])
		if j == 0 || j == n-1 {
			dwon = !dwon
		}
		if dwon {
			j++
		} else {
			j--
		}
	}
	newRows := make([]byte, 0, len(s))
	for _, v := range rows {
		newRows = append(newRows, v...)
	}
	return string(newRows)
}

/*
//[leetcode官方 方法二：按行访问]
*/
func Convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	t := 2*numRows - 2 //周期
	res := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += t { //每次加一个周期
			//本周期内的第一个字符
			res += string(s[i+j])
			//本周期内的第二个字符
			if i != 0 && i != numRows-1 && j+t-i < n {
				res += string(s[j+t-i])
			}
		}
	}
	return res
}

// 0   4   8
// 1 3 5 7 9
// 2   6   10

// 0     6      12
// 1   5 7   11 13
// 2 4   8 10   14
// 3     9      15

// 0       8
// 1     7
// 2   6
// 3 5
// 4
