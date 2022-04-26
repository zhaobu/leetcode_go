package main

/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start

/*
解法1 拓扑排序,队列实现(图的广度优先遍历)
拓扑排序算法过程：
1. 选择图中一个入度为0的点，记录下来
2. 在图中删除该点和所有以它为起点的边
3. 重复1和2，直到图为空或没有入度为0的点
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	/*
		edge[i]=[a,b,c]表示以i为起点的边有(i,a),(i,b),(i,c)
			edge = [[2]    #表示0->2,以0为起点的边为(0,2)
			    [2]     #表示1->2,以1为起点的边为(1,2)
			    [3, 4]  #表示2->3,2->4,以2为起点的边为(3,4)
			    []      #表示没有以3为起点的边
			    []]     #表示没有以4为起点的边
	*/
	edge := make([][]int, numCourses) //邻接表存储图的边
	/*
		indeg[i]=k表示以i为终点的边有k条
	*/
	indeg := make([]int, numCourses) //每个点的入度
	ret := []int{}
	//初始化边
	for _, v := range prerequisites {
		start, end := v[1], v[0]
		edge[start] = append(edge[start], end)
		indeg[end]++
	}
	// fmt.Printf("edge=%+v,indeg=%+v\n", edge, indeg)

	queue := []int{}
	//一次性将所有入度为0的点全部入队
	for i := range indeg {
		if indeg[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		//删除入度为0的点
		head := queue[0]
		queue = queue[1:]
		ret = append(ret, head)
		//删除所有以该点为起点的所有边,并判断删除边后的点入度是否为0
		for endIdx := 0; endIdx < len(edge[head]); endIdx++ {
			end := edge[head][endIdx]
			indeg[end]--
			if indeg[end] == 0 {
				queue = append(queue, end)
			}
		}
	}
	if len(ret) == numCourses {
		return ret
	}
	return []int{}
}

/*
解法2 拓扑排序,栈实现(图的深度优先遍历)
1. 在后序遍历的位置记录访问的节点,因为是深度优先,所以最先记录的就是最后一个课程
2. 后序遍历的这一特点很重要，之所以拓扑排序的基础是后序遍历，
是因为一个任务必须等到它依赖的所有任务都完成之后才能开始开始执行
*/

func findOrder2(numCourses int, prerequisites [][]int) []int {

	edge := make([][]int, numCourses) //邻接表存储图
	/*
		1. visited 记录哪些节点被遍历过，而 onPath 记录当前递归堆栈中有哪些节点
		2. 类比贪吃蛇游戏，visited 记录蛇经过过的格子，而 onPath 仅仅记录蛇身。
		onPath 用于判断是否成环，类比当贪吃蛇自己咬到自己（成环）的场景
	*/
	visited := make([]bool, numCourses) //记录哪些节点被遍历过,防止走回头路
	onPath := make([]bool, numCourses)  //记录当前递归堆栈中有哪些节点
	circle := false                     //是否存在环

	ret := make([]int, 0, numCourses)

	var dfs func(u int)
	dfs = func(u int) {
		if onPath[u] {
			circle = true
		}
		if visited[u] || circle {
			return
		}
		visited[u] = true
		onPath[u] = true
		//遍历该节点的每一个相邻节点 v
		for _, v := range edge[u] {
			dfs(v)
		}
		//回溯
		onPath[u] = false
		ret = append(ret, u)
	}

	//初始化边
	for _, v := range prerequisites {
		start, end := v[1], v[0]
		edge[end] = append(edge[end], start)
	}

	/*
		图中并不是所有节点都相连，所以要用一个 for 循环将所有节点都作为起点调用一次 DFS 搜索算法
	*/
	for i := 0; i < numCourses && !circle; i++ {
		dfs(i)
	}

	if circle { //存在环
		return []int{}
	}

	/*
		1. 如果初始化时是初始化的逆邻接表,就不用反转
		2. 因为是深度优先,所以越是靠近栈底的越是最后才能选修的,越是靠近栈顶的越是没有前置课程的
		3. 栈底是最后一门课,栈顶是第一门课,反转后即为一种结果

	*/
	// for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
	// 	ret[i], ret[j] = ret[j], ret[i]
	// }

	return ret
}

/*
解法3 拓扑排序,栈实现(图的深度优先遍历)
和方法2思想类似,只不过优化visited和onPath为一个变量
*/
func findOrder3(numCourses int, prerequisites [][]int) []int {

	edge := make([][]int, numCourses) //邻接表存储图
	/*
		0: 「未搜索」：我们还没有搜索到这个节点；
		1: 「搜索中」：我们搜索过这个节点，但还没有回溯到该节点，即该节点还没有入栈，还有相邻的节点没有搜索完成）；
		2: 「已完成」：我们搜索过并且回溯过这个节点，即该节点已经入栈，并且所有该节点的相邻节点都出现在栈的更底部的位置，满足拓扑排序的要求。
	*/
	visited := make([]byte, numCourses) //记录点的访问状态
	circle := false                     //是否存在环

	ret := []int{}

	var dfs func(u int)
	dfs = func(u int) {
		if visited[u] != 0 || circle {
			if visited[u] == 1 {
				circle = true
			}
			return
		}
		// 将当前搜索的节点 u 标记为「搜索中」，
		visited[u] = 1
		//遍历该节点的每一个相邻节点 v
		for _, v := range edge[u] {
			dfs(v)
		}
		//当 u 的所有相邻节点都为「已完成」时，我们将 u 放入栈中，并将其标记为「已完成」。
		visited[u] = 2
		ret = append(ret, u)
	}

	//初始化边
	for _, v := range prerequisites {
		start, end := v[1], v[0]
		edge[end] = append(edge[end], start)
	}

	/*
		图中并不是所有节点都相连，所以要用一个 for 循环将所有节点都作为起点调用一次 DFS 搜索算法
	*/
	for i := 0; i < numCourses && !circle; i++ {
		dfs(i)
	}

	if circle { //存在环
		return []int{}
	}

	/*
		1. 如果初始化时是初始化的逆邻接表,就不用反转
		2. 因为是深度优先,所以越是靠近栈底的越是最后才能选修的,越是靠近栈顶的越是没有前置课程的
		3. 栈底是最后一门课,栈顶是第一门课,反转后即为一种结果

	*/
	// for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
	// 	ret[i], ret[j] = ret[j], ret[i]
	// }

	return ret
}

// @lc code=end
