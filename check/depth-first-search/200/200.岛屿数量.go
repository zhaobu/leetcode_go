package main

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start

/*
解法1 dfs
1. 把二维数组看做是有4条边的图进行dfs遍历
2. 把遍历过的所有地方都变为0可以避免重复达到,可以少用visit数组
*/
func numIslands1(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	if m < 1 || n < 1 {
		return 0
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j)
				ret++
			}
		}
	}

	return ret
}

/*
解法2 bfs
1. 使用队列进行bfs遍历
2. 把遍历过的所有地方都变为0可以避免重复达到,可以少用visit数组
*/
func numIslands2(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	if m < 1 || n < 1 {
		return 0
	}

	type land struct {
		row int
		col int
	}
	ret := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				ret++

				grid[i][j] = '0'
				queue := []*land{{i, j}}

				for len(queue) > 0 {
					front := queue[0]
					queue = queue[1:]
					row, col := front.row, front.col
					if row-1 >= 0 && grid[row-1][col] == '1' {
						queue = append(queue, &land{row: row - 1, col: col})
						grid[row-1][col] = '0'
					}
					if row+1 < m && grid[row+1][col] == '1' {
						queue = append(queue, &land{row: row + 1, col: col})
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						queue = append(queue, &land{row: row, col: col - 1})
						grid[row][col-1] = '0'
					}
					if col+1 < n && grid[row][col+1] == '1' {
						queue = append(queue, &land{row: row, col: col + 1})
						grid[row][col+1] = '0'
					}
				}

			}
		}
	}

	return ret
}

/*
解法3 并查集
quickfind写法,超时
*/

func numIslands3(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	if m < 1 || n < 1 {
		return 0
	}

	//初始化并查集
	parent := make([]int, m*n)
	count := 0 //连通分量的个数

	// 查找v所属的集合(根节点)
	Find := func(v int) int {
		if v < 0 || v > len(parent) {
			panic("索引超出范围")
		}
		return parent[v]
	}

	//合并v1,v2所属的集合
	Uinon := func(v1, v2 int) {
		// Quick Find 的 union(v1, v2)：让 v1 所在集合的所有元素都指向 v2 的根节点

		p1 := Find(v1)
		p2 := Find(v2)
		for i, v := range parent {
			if v == p1 {
				parent[i] = p2
			}
		}
		count--
		return
	}

	//检查v1,v2是否属于同一集合
	IsSame := func(v1, v2 int) bool {
		return Find(v1) == Find(v2)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
			} else {
				parent[i*n+j] = -1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				v1, v2 := -1, i*n+j
				if i-1 >= 0 && grid[i-1][j] == '1' {
					v1 = (i-1)*n + j
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if i+1 < m && grid[i+1][j] == '1' {
					v1 = (i+1)*n + j
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					v1 = i*n + j - 1
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if j+1 < n && grid[i][j+1] == '1' {
					v1 = i*n + j + 1
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
			}
		}
	}

	return count
}

/*
解法4 并查集
quickunion写法
*/
func numIslands4(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	if m < 1 || n < 1 {
		return 0
	}

	//初始化并查集
	parent := make([]int, m*n)
	count := 0 //连通分量的个数

	// 查找v所属的集合(根节点)
	Find := func(v int) int {
		if v < 0 || v > len(parent) {
			panic("索引超出范围")
		}
		for v != parent[v] {
			v = parent[v]
		}
		return v
	}

	//合并v1,v2所属的集合
	Uinon := func(v1, v2 int) {
		// Quick Union 的 union(v1, v2)：让 v1 的根节点指向 v2 的根节点

		p1 := Find(v1)
		p2 := Find(v2)
		if p1 == p2 {
			return
		}
		parent[p1] = p2
		count--
		return
	}

	//检查v1,v2是否属于同一集合
	IsSame := func(v1, v2 int) bool {
		return Find(v1) == Find(v2)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
			} else {
				parent[i*n+j] = -1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				v1, v2 := -1, i*n+j
				if i-1 >= 0 && grid[i-1][j] == '1' {
					v1 = (i-1)*n + j
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if i+1 < m && grid[i+1][j] == '1' {
					v1 = (i+1)*n + j
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					v1 = i*n + j - 1
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if j+1 < n && grid[i][j+1] == '1' {
					v1 = i*n + j + 1
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
			}
		}
	}

	return count
}

/*
解法5 并查集
quickunion写法优化
优化方案
基于rank的优化
1. 用rank来记录每个连通分量树的高度
2. 较矮的树接到较高的树上,成为较高数的子树,此时树高度不变
3. 2棵树等高时,任意一颗接到另一个树上,另一颗数此时成了树的root节点,高度+1

其他优化
路径压缩:路径压缩使路径上的所有节点都指向根节点，所以实现成本稍高
路径分裂:使路径上的每个节点都指向其祖父节点（parent的parent）
路径减半:使路径上每隔一个节点就指向其祖父节点（parent的parent）

3种方案都能降低树的高度
路径分裂、路径减半的效率差不多，但都比路径压缩要好

*/
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])

	if m < 1 || n < 1 {
		return 0
	}

	parent := make([]int, m*n) //初始化并查集
	count := 0                 //连通分量的个数
	rank := make([]int, m*n)   //基于rank的优化(矮的树 嫁接到 高的树)

	// 查找v所属的集合(根节点),
	Find := func(v int) int {
		if v < 0 || v > len(parent) {
			panic("索引超出范围")
		}
		// 路径分裂:使路径上的每个节点都指向其祖父节点（parent的parent）
		for v != parent[v] {
			p := parent[v]
			parent[v] = parent[p]
			v = p
		}
		/*
		 路径减半:使路径上每隔一个节点就指向其祖父节点（parent的parent）
		 优化方案只能选一种
		*/
		// for v != parent[v] {
		// 	parent[v] = parent[parent[v]]
		// 	v = parent[v]
		// }
		return v
	}

	//合并v1,v2所属的集合
	Uinon := func(v1, v2 int) {
		// Quick Union 的 union(v1, v2)：让 v1 的根节点指向 v2 的根节点

		p1 := Find(v1)
		p2 := Find(v2)
		if p1 == p2 {
			return
		}
		/*
			1. 较矮的树接到较高的树上,成为较高数的子树,此时树高度不变
			2. 2棵树登高时,任意一颗接到另一个树上,另一颗数此时成了树的root节点,高度+1
		*/
		if rank[p1] < rank[p2] {
			parent[p1] = p2
		} else if rank[p1] > rank[p2] {
			parent[p2] = p1
		} else {
			parent[p1] = p2
			rank[p2]++
		}
		count--
		return
	}

	//检查v1,v2是否属于同一集合
	IsSame := func(v1, v2 int) bool {
		return Find(v1) == Find(v2)
	}

	/*
		初始化时要记录每个连通分量的高度1
	*/
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
				rank[i] = 1
			} else {
				parent[i*n+j] = -1
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				v1, v2 := -1, i*n+j
				if i-1 >= 0 && grid[i-1][j] == '1' {
					v1 = (i-1)*n + j
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if i+1 < m && grid[i+1][j] == '1' {
					v1 = (i+1)*n + j
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					v1 = i*n + j - 1
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
				if j+1 < n && grid[i][j+1] == '1' {
					v1 = i*n + j + 1
					if !IsSame(v1, v2) {
						Uinon(v1, v2)
					}
				}
			}
		}
	}

	return count
}

// @lc code=end
