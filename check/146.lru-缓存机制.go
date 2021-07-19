/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 3000
)

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
	used     int //
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	// 删除节点当前位置
	node.prev.next = node.next
	node.next.prev = node.prev

	//把该节点移动到链表头部
	node.next = this.head
	this.head.prev = node
	this.head = node
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	//判断
	node, ok := this.cache[key]
	if !ok {
		node = &CacheNode{
			next:  this.head,
			key:   key,
			value: value,
		}
		this.head.prev = node
		this.head = node
		if this.tail == nil {
			this.tail = node

		}
		return
	}
	//找到了就修改值并且放到头部

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end




