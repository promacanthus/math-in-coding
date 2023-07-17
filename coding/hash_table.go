package coding

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

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

// ---------- 有效的变位词 ----------
// 只有小写英文字母的情况。
func isAnagramV1(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := make([]int, 26)
	for _, v := range s1 {
		count[v-'a']++
	}
	for _, v := range s2 {
		if count[v-'a'] == 0 {
			return false
		}
		count[v-'a']--
	}
	return true
}

// 所有字符的情况，ASCII 不够，需要 Unicode。
func isAnagramV2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	set := make(map[rune]int)
	for _, v := range s1 {
		set[v]++
	}
	for _, v := range s2 {
		if set[v] == 0 {
			return false
		}
		set[v]--
	}
	return true
}

// ---------- 变位词组 ----------
// 将变位词的每一个单词都映射成一个质数，然后每个单词的每一个字母相乘得到一个乘积，
// 根据乘法交换律，变位词的乘积是一样的，根据乘积将这些单词分组即可。
// 由于每个字母都映射到一个质数，因此不互为变位词的两个单词一定会映射到不同的数字。
//
// 时间复杂度，O(n*m),有 n 个单词，每个单词有 m 个字母。
// NOTICE: 如果单词比较长的话，这个方法可能会溢出的。
func groupAnagramV1(strs []string) [][]string {
	primeNumber := generatePrimes(26)

	groups := make(map[int][]string)
	for _, str := range strs {
		wordHast := 1
		for _, s := range str {
			wordHast *= primeNumber[s-'a']
		}
		groups[wordHast] = append(groups[wordHast], str)
	}
	res := make([][]string, 0)
	for _, group := range groups {
		res = append(res, group)
	}
	return sortResult(res)
}

// 质数是大于 1 的自然数，它不能被任何比 1 和它本身小的自然数整除。
// 换句话说，质数只能被 1 和它本身整除。
// 如果某个数不是质数，那么一定存在一个因子小于等于它的平方根（如，3*3=9）
func isPrime(n int) bool {
	if n < 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 前 26 个质数:
// 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
// 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101
func generatePrimes(n int) []int {
	var primes []int
	for i := 2; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// 把一组变位词映射到同一个单词。因为，变位词的单词出现的次数相同。
// 将字母中的单词排序之后就会变成同一个字符串。
// 时间复杂度，O(n*mlogm)：
// 有 n 个单词，每个单词有 m 个字母，对每个单词排序 O(mlogm)
// NOTCIE: 这里虽然效率低了一些，单数不会出现溢出的情况。
func groupAnagramV2(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, str := range strs {
		byteArray := []byte(str)
		sort.Slice(byteArray, func(i, j int) bool {
			return byteArray[i] < byteArray[j]
		})
		groups[string(byteArray)] = append(groups[string(byteArray)], str)
	}
	var res [][]string
	for _, group := range groups {
		res = append(res, group)
	}
	return sortResult(res)
}

// range map is random.
func sortResult(res [][]string) [][]string {
	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})
	return res
}

// ---------- 外星语言是否排序 ----------
// 构建一个索引，把每一个字母和对应的值建立好映射关系。
func isAlienSorted(words []string, order string) bool {
	orderIdx := make([]int, len(order))
	for i := 0; i < len(order); i++ {
		orderIdx[order[i]-'a'] = i
	}
	for i := 0; i < len(words)-1; i++ {
		if !isSorted(words[i], words[i+1], orderIdx) {
			return false
		}
	}
	return true
}

// 一次比较两个单词中每一个字母，找到第一个不同的字母，排在前面的字母，第一个不同的单词也应该排在前面。
// 如果没有找到不同的字母，短的单词排序的时候靠前。
func isSorted(word1, word2 string, orderIdx []int) bool {
	i := 0
	for ; i < len(word1) && i < len(word2); i++ {
		if orderIdx[word1[i]-'a'] < orderIdx[word2[i]-'a'] {
			return true
		}
		if orderIdx[word1[i]-'a'] > orderIdx[word2[i]-'a'] {
			return false
		}
	}
	// 说明上面的两个单词没有不同的字母，判断是否前一个单词更短。
	return i == len(word1)
}

// ---------- 最小时间差 ----------
// 排列组合，两两比较。
// 时间复杂度: O(n^2)
func findMinimalDifferenceV1(times []string) time.Duration {
	minDiff := time.Duration(time.Hour * 24)
	for i := 0; i < len(times); i++ {
		for j := i + 1; j < len(times); j++ {
			interval := findInterval(times[i], times[j])
			absInterval := interval.Abs()
			if absInterval < minDiff {
				minDiff = absInterval
			}
		}
	}
	return minDiff
}

// 先排序在遍历一遍比较前后两个的差值。
// 时间复杂度: O(nlogn)
func findMinimalDifferenceV2(times []string) time.Duration {
	sort.Slice(times, func(i, j int) bool {
		return findInterval(times[i], times[j]) < 0
	})
	minDiff := time.Duration(time.Hour * 24)
	for i := 0; i < len(times)-1; i++ {
		interval := findInterval(times[i], times[i+1])
		absInterval := interval.Abs()
		if absInterval < minDiff {
			minDiff = absInterval
		}
	}
	return minDiff
}

// 使用哈希表将24小时的每一分钟就保存下来。
// 时间复杂度: O(n) 对于固定常数循环次数，时间复杂度是 O(1)
// 空间复杂度：O(1) 对于固定常数大小的空间，空间复杂度是 O(1)
func findMinimalDifferenceV3(times []string) time.Duration {
	interval := (time.Hour * 24).Minutes()
	if len(times) > int(interval) {
		// 如果输入的时间点大于 1440，
		// 肯定存在重复的时间点，最小时间间隔就是 0
		return time.Duration(0)
	}

	timeSet := make([]bool, int(interval))
	for _, timeString := range times {
		t, _ := time.Parse(layout, timeString)
		if timeSet[t.Minute()] {
			// 如果时间点已经存在于哈希表中，
			// 最小时间间隔就是 0
			return time.Duration(0)
		}
		timeSet[t.Minute()] = true
	}

	minDiff := interval
	prev := float64(-1)
	first := interval
	last := float64(-1)
	for i := 0; i < len(timeSet); i++ {
		if timeSet[i] {
			if prev >= 0 {
				minDiff = math.Min(float64(i)-prev, minDiff)
			}
			prev = float64(i)
			first = math.Min(float64(i), first)
			last = math.Max(float64(i), float64(last))
		}
	}
	// 把第一个时间加上24小时(1440分)，表示第二天的同一个时间点，
	// 求出它和最后一个时间点之间的差值。
	minDiff = math.Min(first+float64(len(timeSet))-last, minDiff)
	return time.Duration(minDiff * float64(time.Minute))
}

const layout = "15:04"

func findInterval(time1, time2 string) time.Duration {
	t1, _ := time.Parse(layout, time1)
	t2, _ := time.Parse(layout, time2)
	return t1.Sub(t2)
}
