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
type LRUCache struct {
}

func Constructor(capacity int) LRUCache {

}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
