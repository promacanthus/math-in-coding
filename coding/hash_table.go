package coding

import "math/rand"

// ---------- 插入、删除和随机访问都是O(1)的容器 ----------
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
	c.values = append(c.values[:idx], c.values[idx+1:]...)
	return true
}

func (c *container) GetRandom() int {
	i := rand.Intn(len(c.values))
	return c.values[i]
}
