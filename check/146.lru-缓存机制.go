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

	hnext *CacheNode // 将结点串在散列表的拉链中
}

// @lc code=start
type LRUCache struct {
	node []CacheNode // 散列表,存放拉链

	head *CacheNode // lru head node
	tail *CacheNode // lru tail node

	capacity int //
	used     int //
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		node:     make([]CacheNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.tail == nil {
		return -1
	}

	//通过散列表，可以很快在缓存中找到一个数据
	if tmp := this.searchNode(key); tmp != nil {
		this.moveToTail(tmp)
		return tmp.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if tmp := this.searchNode(key); tmp != nil { //插入数据时已在lru中
		tmp.value = value
		this.moveToTail(tmp)
		return
	}
	//插入数据不在lru中
	this.addNode(key, value)

	//且lru已满
	if this.used > this.capacity {
		this.delNode()
	}
}

func (this *LRUCache) addNode(key int, value int) {
	newNode := &CacheNode{
		key:   key,
		value: value,
	}

	tmp := &this.node[hash(key)] //找到新加入的节点hash到的拉链的头部
	//把node插入拉链的队头后面
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	this.used++

	if this.tail == nil { //如果缓存队列为空.直接把头和尾都指向新节点
		this.tail, this.head = newNode, newNode
		return
	}

	//如果缓存队列不为空,新加入的节点要放到队尾
	this.tail.next = newNode
	newNode.prev = this.tail
	this.tail = newNode
}

//删除队头节点(最先加入的节点)
func (this *LRUCache) delNode() {
	if this.head == nil {
		return
	}

	prev := &this.node[hash(this.head.key)] //找到头节点所在拉链的头部
	tmp := prev.hnext

	for tmp != nil && tmp.key != this.head.key {
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil { //拉链中没找到头节点没有找到
		return
	}
	//找到了,删除拉链中的头节点
	prev.hnext = tmp.hnext

	//删除缓存队列中的头节点
	this.head = this.head.next
	this.head.prev = nil
	this.used--
}

func (this *LRUCache) searchNode(key int) *CacheNode {
	if this.tail == nil {
		return nil
	}

	// 查找
	tmp := this.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (this *LRUCache) moveToTail(node *CacheNode) {
	if this.tail == node { //node已经是尾结点
		return
	}
	if this.head == node { //要移动的是头节点
		this.head = node.next
		this.head.prev = nil
	} else { //node是中间节点
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	//把node加入当前尾结点后面
	node.next = nil
	this.tail.next = node
	node.prev = this.tail

	//把链表尾指针指向node
	this.tail = node
}

func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end




