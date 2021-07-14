/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start

type CacheNode struct {
	next  *CacheNode
	key   int
	value int
}

// @lc code=start
type LRUCache struct {
	Data     *CacheNode
	Length   int
	Capacity int
}

func Constructor(capacity int) LRUCache {
	cache := &LRUCache{
		Data:     nil,
		Length:   0,
		Capacity: capacity,
	}
	return *cache
}

func (this *LRUCache) Get(key int) int {
	var (
		cur  *CacheNode = this.Data
		last *CacheNode = cur //保存用于找到的key的上一个节点,方便删除
	)

	//查找key
	for cur != nil {
		if cur.key == key {
			if last != cur { //表示访问到的不是第一个节点
				last.next = cur.next //将找到的节点从原来的位置删除
				// 然后再插入到链表的头部
				cur.next = this.Data
				this.Data = cur
			}
			return cur.value
		}
		last = cur
		cur = cur.next
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	var (
		cur      *CacheNode = this.Data
		last     *CacheNode //保存用于找到的key的上一个节点,方便删除
		tailLast *CacheNode //尾节点的前一个节点
	)

	//查找key
	for cur != nil {
		if cur.key == key {
			//找到了后就修改该值
			cur.value = value
			if last != cur && last != nil { //表示访问到的不是第一个节点,&& last != nil过滤掉容量>1,反复put同一个key时length为1的情况
				last.next = cur.next //将找到的节点从原来的位置删除
				// 然后再插入到链表的头部
				cur.next = this.Data
				this.Data = cur
			}
			return
		}
		tailLast = last
		last = cur
		cur = cur.next
	}
	//未找到就直接移动到头部
	if this.Length == this.Capacity { //满了就先删除最后一个节点
		if tailLast != nil { //容量>1的情况
			tailLast.next = nil
		} else { //容量=1的情况
			this.Data = nil
			this.Length = 0
		}
	}
	//将新的数据结点插入链表的头部
	node := &CacheNode{
		next:  this.Data,
		key:   key,
		value: value,
	}
	this.Data = node
	if this.Length < this.Capacity {
		this.Length++
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end




