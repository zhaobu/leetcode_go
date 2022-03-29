package main

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

// @lc code=start
type Node struct {
	Key  int
	Val  int
	Pre  *Node
	Next *Node
}
type LRUCache struct {
	bottom   *Node
	top      *Node
	capacity int
	cache    map[int]*Node
}

func Constructor(capacity int) LRUCache {
	top, bottom := &Node{}, &Node{}
	top.Next = bottom
	bottom.Pre = top
	return LRUCache{
		bottom:   bottom,
		top:      top,
		capacity: capacity,
		cache:    make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	old, ok := this.cache[key]
	if !ok {
		return -1
	}
	//如果访问的不是top
	if this.top.Next != old {
		//移动到最近访问
		//删除当前访问的节点
		old.Pre.Next = old.Next
		old.Next.Pre = old.Pre
		//插入到最近访问
		this.moveToTop(old)
	}
	return old.Val
}

func (this *LRUCache) Put(key int, value int) {
	old, ok := this.cache[key]
	if ok {
		old.Val = value
		//移动到最近访问
		//top, 1, 2, 3,  bottom
		//删除当前访问的节点
		old.Pre.Next = old.Next
		old.Next.Pre = old.Pre
		//插入到最近访问
		this.moveToTop(old)
		return
	}
	if len(this.cache) == this.capacity { //达到最大容量
		//删除最久未访问的
		delete(this.cache, this.bottom.Pre.Key)
		this.bottom.Pre.Pre.Next = this.bottom
		this.bottom.Pre = this.bottom.Pre.Pre
	}
	// 构建新节点
	new := &Node{
		Key: key,
		Val: value,
	}
	this.cache[key] = new
	//移动到最近访问
	this.moveToTop(new)
}

func (this *LRUCache) moveToTop(n *Node) {
	this.top.Next.Pre = n
	n.Next = this.top.Next
	this.top.Next = n
	n.Pre = this.top
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
