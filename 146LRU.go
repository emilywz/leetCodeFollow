package main

//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。
// 它应该支持以下操作： 获取数据 get 和 写入数据 put 。
//
//获取数据 get(key) - 如果关键字 (key) 存在于缓存中，
// 则获取关键字的值（总是正数），否则返回 -1。
//写入数据 put(key, value) - 如果关键字已经存在，则变更其数据值；
// 如果关键字不存在，则插入该组「关键字/值」。
// 当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//进阶:
//你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//示例:
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得关键字 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得关键字 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4



//type LRUCache struct {
//	Capacity int
//	HashMap  map[int]*Node
//	head     *Node
//	end      *Node
//}
//
////双向链表实现
//type Node struct {
//	Key   int
//	Value int
//	pre   *Node
//	next  *Node
//}
//
//func Constructor(capacity int) LRUCache {
//	lru := LRUCache{Capacity: capacity}
//	lru.HashMap = make(map[int]*Node, capacity)
//	return lru
//}
//
//func (this *LRUCache) Get(key int) int {
//	if v, ok := this.HashMap[key]; ok {
//		this.refreshNode(v)
//		return v.Value
//	} else {
//		return -1
//	}
//}
//
//func (this *LRUCache) Put(key int, value int) {
//	if v, ok := this.HashMap[key]; !ok {
//		if len(this.HashMap) >= this.Capacity {
//			//移除旧节点
//			oldKey := this.removeNode(this.head)
//			delete(this.HashMap, oldKey)
//		}
//		node := Node{Key: key, Value: value}
//		this.addNode(&node)
//		this.HashMap[key] = &node
//	} else {
//		v.Value = value
//		this.refreshNode(v)
//	}
//}
//
//func (this *LRUCache) refreshNode(node *Node) {
//	if node == this.end {
//		return
//	}
//	this.removeNode(node)
//	this.addNode(node)
//}
//
//func (this *LRUCache) addNode(node *Node) {
//	if this.end != nil {
//		this.end.next = node
//		node.pre = this.end
//		node.next = nil
//	}
//	this.end = node
//	if this.end == nil {
//		this.head = node
//	}
//}
//
//func (this *LRUCache) removeNode(node *Node) int {
//	if node == this.end {
//		this.end = this.end.pre
//		this.end.next = nil
//	} else if node == this.head {
//		this.head = this.head.next
//		this.head.pre = nil
//	} else {
//		node.pre.next = node.next
//		node.next.pre = node.pre
//	}
//	return node.Key
//}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
