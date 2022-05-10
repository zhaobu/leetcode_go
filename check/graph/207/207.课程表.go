package main

/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start

/*
解法1 拓扑排序,队列实现(图的广度优先遍历)
拓扑排序算法过程：
1. 选择图中一个入度为0的点，记录下来
2. 在图中删除该点和所有以它为起点的边
3. 重复1和2，直到图为空或没有入度为0的点
*/
func canFinish1(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses) //邻接表存储图的边
	/*
		indeg[i]=k表示以i为终点的边有k条
	*/
	indeg := make([]int, numCourses) //每个点的入度
	//初始化边
	for _, v := range prerequisites {
		from, to := v[1], v[0]
		edges[from] = append(edges[from], to)
		indeg[to]++
	}
	// fmt.Printf("edges=%+v,indeg=%+v\n", edges, indeg)

	queue := []int{}
	//一次性将所有入度为0的点全部入队
	for i, v := range indeg {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		//删除入度为0的点
		head := queue[0]
		queue = queue[1:]
		numCourses--
		//删除所有以该点为起点的所有边,并判断删除边后的点入度是否为0
		for _, to := range edges[head] {
			indeg[to]--
			if indeg[to] == 0 {
				queue = append(queue, to)
			}
		}
	}
	return numCourses == 0
}

/*
解法2 拓扑排序,栈实现(图的深度优先遍历)
*/
func canFinish3(numCourses int, prerequisites [][]int) bool {

	edges := make([][]int, numCourses) //邻接表存储图
	/*
		0: 「未搜索」：我们还没有搜索到这个节点；
		1: 「搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成）；
		2: 「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。
	*/
	visited := make([]int, numCourses) //记录点的访问状态
	circle := false                    //是否存在环

	ret := make([]int, 0, numCourses)

	var dfs func(u int)
	dfs = func(u int) {
		// 将当前搜索的节点 u 标记为「搜索中」，
		visited[u] = 1
		//遍历该节点的每一个相邻节点 v
		for _, v := range edges[u] {
			if visited[v] == 0 { //如果 v 为「未搜索」，那么我们开始搜索 v，待搜索完成回溯到 u
				dfs(v)
				if circle { //如果已经判断出有环,就提前结束
					return
				}
			} else if visited[v] == 1 { //如果 v 为「搜索中」，说明就找到了图中的一个环，不存在拓扑排序
				circle = true
				return
			}
			/*
				else if  visited[v] == 2{
					如果 v 为「已完成」，那么说明 v 已经在栈中了,
					而 u 还不在栈中，不用进行任何操作
				}
			*/
		}
		//当 u 的所有相邻节点都为「已完成」时，我们将 u 放入栈中，并将其标记为「已完成」。
		visited[u] = 2
		ret = append(ret, u)
	}

	//初始化边
	for _, v := range prerequisites {
		start, end := v[1], v[0]
		edges[end] = append(edges[end], start)
	}

	//遍历每一个点
	for i := 0; i < numCourses && !circle; i++ {
		//如果visited[i]==0 说明该点是一个起点
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return !circle
}

// @lc code=end
