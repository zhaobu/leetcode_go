package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	fmt.Printf("%+v \n", change(matrix))
}

/*
matrix =
[

[1,2,3,4],
[5,6,7,8],
[9,10,11,12]

]

输出：[1,2,3,4,8,12,11,10,9,5,6,7]
*/
func change(matrix [][]int) (ret []int) {
	n := len(matrix)
	if n == 0 {
		return ret
	}
	m := len(matrix[0])

	dirs := [][]int{
		{0, 1},
		{1, 0},
		{0, -1}, //左
		{-1, 0}, //上
	}

	visited := make([][]bool, n)
	for i := range matrix {
		visited[i] = make([]bool, m)
	}

	count := m * n
	curDir := 0
	i, j := 0, 0
	for count > 0 {
		ret = append(ret, matrix[i][j])
		visited[i][j] = true
		count--
		i1, j1 := i+dirs[curDir][0], j+dirs[curDir][1]
		if i1 < 0 || i1 >= n || j1 < 0 || j1 >= m || visited[i1][j1] {
			curDir = (curDir + 1) % 4
		}
	}

	return ret
}
