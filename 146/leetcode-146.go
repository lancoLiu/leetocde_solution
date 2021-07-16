package main

// LRU (最近最少使用) 缓存机制

//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[   [2],    [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]

//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]

/**
LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间


在 O(1) 时间复杂度内完成Get和Put这两种操作
*/

//读取之后需要更新存储
//如果满了需要删除之后再存储

//要让PUT和GET都是O(1),我们可以总结出cache必须得查找快，插入快，删除快，顺序区分
//哈希表查找快但是无顺序
//链表插入快但是查找慢
type LRUCache struct {
	Size     int
	Capacity int
	//原因
	Cache      map[int]*LinkNode
	Head, Tail *LinkNode
}

type LinkNode struct {
	Key, Val   int
	prev, next *LinkNode
}

func initLinkNode(key, val int) *LinkNode {
	return &LinkNode{
		Key: key,
		Val: val,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		Size:     0,
		Capacity: capacity,
		Cache:    map[int]*LinkNode{},
		Head:     initLinkNode(0, 0),
		Tail:     initLinkNode(0, 0),
	}
	l.Head.next = l.Tail
	l.Tail.next = l.Head
	return l
}

func (this *LRUCache) addToHead(node *LinkNode) {
	//当前节点的前驱节点指向head
	node.prev = this.Head
	//当前节点的后驱节点指向head节点的后驱节点
	node.next = this.Head.next
	//head节点的后驱节点的前驱节点指向node
	this.Head.next.prev = node
	//head节点的后驱节点指向node
	this.Head.next = node

}

/**
删除当前节点
*/
func (this *LRUCache) removeNode(node *LinkNode) {
	node.next.prev = node.prev
	node.prev.next = node.next
}

/**
移动到头节点
*/
func (this *LRUCache) moveToHead(node *LinkNode) {
	this.removeNode(node)
	this.addToHead(node)

}

/**
移除尾节点
*/
func (this *LRUCache) removeTail() *LinkNode {
	//尾节点
	node := this.Tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Cache[key]; !ok {
		return -1
	}
	node := this.Cache[key]
	this.moveToHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Cache[key]; !ok {
		node := initLinkNode(key, value)
		this.Cache[key] = node
		this.addToHead(node)
		this.Size++
		if this.Size > this.Capacity {
			removed := this.removeTail()
			delete(this.Cache, removed.Key)
			this.Size--

		}
	} else {
		node := this.Cache[key]
		node.Val = value
		this.moveToHead(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
