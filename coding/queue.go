package coding

import (
	"math"
)

// ---------- 滑动窗口的平均值 ----------

type MovingAverage struct {
	capacity int   // the size of the queue
	sum      int   // the sum of the queue
	values   []int // the values of the queue
}

func NewMovingAverage(size int) *MovingAverage {
	return &MovingAverage{
		capacity: size,
		sum:      0,
		values:   make([]int, 0),
	}
}

func (m *MovingAverage) Size() int { return len(m.values) }

func (m *MovingAverage) Capacity() int { return m.capacity }

func (m *MovingAverage) Next(val int) float64 {
	if m.Size() >= m.Capacity() {
		m.sum -= m.values[0]
		m.values = m.values[1:]
	}
	m.sum += val
	m.values = append(m.values, val)
	return float64(m.sum) / float64(m.Size())
}

// ---------- 最近请求次数 ----------
type RecentAverage struct {
	times []int
	size  int
}

func RecentCounter(size int) *RecentAverage {
	return &RecentAverage{
		times: make([]int, 0),
		size:  size,
	}
}

func (r *RecentAverage) Len() int { return len(r.times) }

func (r *RecentAverage) ping(time int) int {
	r.times = append(r.times, time)
	for time > r.times[0]+r.size {
		// 如果当前要入队的值，比队首的元素加上整个队列的容量还要大，
		// 那么，当前元素入队后，队首的元素就要出队了。
		r.times = r.times[1:]
	}
	return r.Len()
}

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		value: val,
	}
}

func (n *TreeNode) AddChild(left *TreeNode, right *TreeNode) {
	if left != nil {
		n.left = left
	}
	if right != nil {
		n.right = right
	}
}

func BFS(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.value)
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
	}
	return res
}

// ---------- 在完全二叉树中添加节点 ----------
// 完全二叉树中第n层的节点数：2^(n-1)。
// 按照广度优先的顺序，找出第1个左子树或者右子树为空的节点，
// 空着的位置就是待插入节点的位置。
// 效率优化，每次插入节点的时候不需要从根节点开始遍历，记住上面待插入节点的父节点，
// 然后判断它的左右子树是否已经满，如果没有满直接插入，否则插入广度优先搜索的下一个节点。
type CBTInserter struct {
	queue []*TreeNode
	root  *TreeNode
}

func NewCBTInserter(root *TreeNode) *CBTInserter {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for queue[0].left != nil && queue[0].right != nil {
		// 初始化的时候，遍历到第一个左或右子节点为空节点，
		// 这个节点就是待插入的节点。
		node := queue[0]
		queue = queue[1:]
		queue = append(queue, node.left)
		queue = append(queue, node.right)
	}
	return &CBTInserter{
		queue: queue,
		root:  root,
	}
}

func (t *CBTInserter) Insert(val int) *TreeNode {
	parent := t.queue[0]
	node := NewTreeNode(val)
	if parent.left == nil {
		parent.left = node
	} else {
		parent.right = node
		// parent 节点的左右子节点已经满了，
		// 待插入节点应该是广度优先搜索的下一个节点。
		t.queue = t.queue[1:]
		t.queue = append(t.queue, parent.left)
		t.queue = append(t.queue, parent.right)
	}
	return parent
}

func (t *CBTInserter) Root() *TreeNode {
	return t.root
}

// ---------- 二叉树中每层的最大值 ----------
// 使用变量保存当前层和下一层的节点数，这样就能确定当前层是否遍历完成。
func LargestValuesV1(root *TreeNode) []int {
	var curCount, nextCount int
	var res []int
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	curCount++
	max := math.MinInt
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		curCount--
		max = int(math.Max(float64(max), float64(node.value)))

		if node.left != nil {
			queue = append(queue, node.left)
			nextCount++
		}
		if node.right != nil {
			queue = append(queue, node.right)
			nextCount++
		}

		if curCount == 0 {
			res = append(res, max)
			max = math.MinInt

			curCount = nextCount
			nextCount = 0
		}
	}
	return res
}

// 用两个队列实现，把不同层的节点放入不同的队列中。
// queue1 : 存放当前层的节点
// queue2 : 存放下一层的节点
func LargestValuesV2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var queue1, queue2 []*TreeNode
	queue1 = append(queue1, root)
	max := math.MinInt
	for len(queue1) > 0 {
		node := queue1[0]
		queue1 = queue1[1:]
		max = int(math.Max(float64(max), float64(node.value)))

		if node.left != nil {
			queue2 = append(queue2, node.left)
		}
		if node.right != nil {
			queue2 = append(queue2, node.right)
		}

		if len(queue1) == 0 {
			res = append(res, max)
			max = math.MinInt
			queue1 = queue2
			queue2 = make([]*TreeNode, 0)
		}
	}
	return res
}

// ---------- 二叉树最低层最左边的树 ----------
// 最低层最左边的节点，就是最后一层的第一个节点。
func findBottomLeftValue(root *TreeNode) int {
	var q1, q2 []*TreeNode
	q1 = append(q1, root)
	res := root.value
	for len(q1) > 0 {
		node := q1[0]
		q1 = q1[1:]

		if node.left != nil {
			q2 = append(q2, node.left)
		}
		if node.right != nil {
			q2 = append(q2, node.right)
		}

		if len(q1) == 0 {
			q1 = q2
			q2 = make([]*TreeNode, 0)
			if len(q1) != 0 {
				res = q1[0].value
			}
		}
	}
	return res
}

// ---------- 二叉树的右侧视图 -----------
// 右侧试图，即看到每一层最右边的一个节点。
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var q1, q2 []*TreeNode
	q1 = append(q1, root)
	for len(q1) > 0 {
		node := q1[0]
		q1 = q1[1:]
		cur := node.value

		if node.left != nil {
			q2 = append(q2, node.left)
		}
		if node.right != nil {
			q2 = append(q2, node.right)
		}

		if len(q1) == 0 {
			res = append(res, cur)
			q1 = q2
			q2 = make([]*TreeNode, 0)
		}
	}
	return res
}
