package coding

import "math/rand"

// ---------- 插入、删除和随机访问都是 O(1) 的容器 ----------
// 使用map作为所用来，key保存值，value保存值在数组中的下标
// 使用数组来保存实际的数据
type Container interface {
	Insert(int) bool
	Remove(int) bool
	GetRandom() int
}

var _ Container = &container{}

type container struct {
	index  map[int]int
	values []int
}

func NewContainer() Container {
	return &container{
		index:  make(map[int]int),
		values: make([]int, 0),
	}
}

func (c *container) Insert(value int) bool {
	if _, ok := c.index[value]; ok {
		return false
	}
	c.index[value] = len(c.values)
	c.values = append(c.values, value)
	return true
}

func (c *container) Remove(value int) bool {
	idx, ok := c.index[value]
	if !ok {
		return false
	}
	delete(c.index, value)
	// 也可以将待删除的数组和数组尾部的数字交换后直接删除数组尾部的数字，时间复杂度也是O(1)
	// 否则数组中的数字删除后，需要把后面的数字移动到前面，这样时间复杂度就是O(n)
	c.values = append(c.values[:idx], c.values[idx+1:]...)
	return true
}

func (c *container) GetRandom() int {
	i := rand.Intn(len(c.values))
	return c.values[i]
}

// ---------- 最近最少使用缓存（LRU） ----------
// 如下两个操作的时间复杂度是 O(1)：
//  1. get(key)
//  2. put(key,value)
//
// 常规的哈希表无法找到最近最少使用的，因此需要稍微改造一下。
// 把存入的元素按照访问的先后顺序存入链表中。每次访问一个元素，
// 该元素都被移到链表的尾部。这样，位于链表头部的元素就是最近最少使用的。
//  1. 把节点从原来的位置删除（需要知道节点的前一个链表，因此使用双链表，O(1）的时间复杂度）
//  2. 把节点添加到链表尾部
type DoubleNode struct {
	key   int
	value int
	prev  *DoubleNode
	next  *DoubleNode
}

func NewDoubleNode(key int, value int) *DoubleNode {
	return &DoubleNode{
		key:   key,
		value: value,
	}
}

// LRUCache (Least Recently Used)
type LRUCache struct {
	capacity int
	set      map[int]*DoubleNode
	// 首尾哨兵节点，简化操作
	head *DoubleNode
	tail *DoubleNode
}

func NewLRUCache(capacity int) *LRUCache {
	head := NewDoubleNode(-1, -1)
	tail := NewDoubleNode(-1, -1)
	head.next = tail
	tail.prev = head

	return &LRUCache{
		capacity: capacity,
		set:      make(map[int]*DoubleNode),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache) Put(key, value int) {
	if node, ok := c.set[key]; ok {
		node.value = value
		moveToTail(c.tail, node)
		return
	}

	node := NewDoubleNode(key, value)
	if len(c.set) == c.capacity {
		toBeDeleted := c.head.next
		deleteNode(toBeDeleted)
		delete(c.set, toBeDeleted.key)
	}
	c.set[key] = node
	insertToTail(c.tail, node)
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.set[key]
	if !ok {
		return -1
	}
	moveToTail(c.tail, node)
	return node.value
}

func deleteNode(node *DoubleNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func insertToTail(tail, node *DoubleNode) {
	tail.prev.next = node
	node.prev = tail.prev
	node.next = tail
	tail.prev = node
}

func moveToTail(tail, node *DoubleNode) {
	deleteNode(node)
	insertToTail(tail, node)
}
