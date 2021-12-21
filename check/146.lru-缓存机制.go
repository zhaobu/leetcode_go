package check

/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start

// 使用双向链表存储缓存数据
type CacheNode struct {
	// 前驱和后继指针将结点串在双向链表中
	prev *CacheNode
	next *CacheNode

	key   int // lru key
	value int // lru value
}

// @lc code=start
type LRUCache struct {
	head *CacheNode // lru head node
	tail *CacheNode // lru tail node

	cache map[int]*CacheNode

	capacity int //
	size     int //
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		head:     &CacheNode{},
		tail:     &CacheNode{},
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*CacheNode, capacity),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// 删除节点当前位置
	this.deleteNode(node)
	//把该节点移动到链表头部
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	//判断
	node, ok := this.cache[key]
	if !ok {
		this.size++
		if this.size > this.capacity { //超出容量要删除最后一个节点
			delKey := this.tail.prev.key
			this.tail.prev.prev.next = this.tail
			this.tail.prev = this.tail.prev.prev
			delete(this.cache, delKey)
			this.size--
		}
		//新节点加入到头部
		this.addToHead(&CacheNode{key: key, value: value})
		return
	}
	//找到了就修改值并且放到头部
	node.value = value

	// 删除节点当前位置
	this.deleteNode(node)
	//把该节点移动到链表头部
	this.moveToHead(node)
	return
}

func (this *LRUCache) deleteNode(node *CacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *CacheNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) addToHead(node *CacheNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	this.cache[node.key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
