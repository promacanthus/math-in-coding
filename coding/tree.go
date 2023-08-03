package coding

import (
	"math"
	"strconv"
	"strings"
)

// in order
func inOrder(root *TreeNode) []int {
	res := make([]int, 0)
	inOrderRecurse(root, &res)
	return res
}

func inOrderRecurse(root *TreeNode, res *[]int) {
	if root != nil {
		inOrderRecurse(root.left, res)
		*res = append(*res, root.value)
		inOrderRecurse(root.right, res)
	}
}

func inOrderByStack(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.value)
		cur = cur.right
	}
	return res
}

// pre order
func preOrder(root *TreeNode) []int {
	res := make([]int, 0)
	preOrderRecurse(root, &res)
	return res
}

func preOrderRecurse(root *TreeNode, res *[]int) {
	if root != nil {
		*res = append(*res, root.value)
		preOrderRecurse(root.left, res)
		preOrderRecurse(root.right, res)
	}
}

func preOrderByStack(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.value)
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.right
	}
	return res
}

// post order
func postOrder(root *TreeNode) []int {
	res := make([]int, 0)
	postOrderRecurse(root, &res)
	return res
}

func postOrderRecurse(root *TreeNode, res *[]int) {
	if root != nil {
		postOrderRecurse(root.left, res)
		postOrderRecurse(root.right, res)
		*res = append(*res, root.value)
	}
}

func postOrderByStack(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	cur := root
	var prev *TreeNode
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		if cur.right != nil && cur.right != prev {
			// 如果 cur 节点有右子节点并且右子节点不是前一个遍历的节点，
			// 则表示它有右子树并且右子树还没有遍历过，按照后序遍历的顺序，
			// 应该先遍历它的右子树，因此把变量cur指向它的右子节点。
			cur = cur.right
		} else {
			// 如果 cur 指向的节点没有右子树或它的右子树已经遍历过，
			// 则按照后序遍历的顺序，此时可以遍历这个节点，于是把它出栈并遍历它。
			// 接下来准备遍历下一个节点，于是把 prev 指向当前节点。
			// 下一个遍历的节点一定是它的父节点，而父节点之前已经存放到栈中，
			// 所以需要将 cur 重置为 null，这样下一次就可以将它的父节点出栈并遍历。
			stack = stack[:len(stack)-1]
			res = append(res, cur.value)
			prev = cur
			cur = nil
		}
	}
	return res
}

// ---------- 二叉树剪枝 ----------
// 如果一个节点要被删除，那么它的左右子节点也同时可以被删除，
// 使用后序遍历最符合需求，因为该节点的左右子节点肯定比它先被遍历。
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.left = pruneTree(root.left)
	root.right = pruneTree(root.right)
	if root.left == nil && root.right == nil && root.value == 0 {
		// return nil represents this node has been pruned
		return nil
	}
	return root
}

// ---------- 序列化和反序列化二叉树 -----------
// 以前序遍历的顺序遍历二叉树最适合序列化。如果采用前序遍历的顺序，
// 那么二叉树的根节点最先序列化到字符串中，然后是左子树，最后是右子树。
// 这样做的好处是在反序列化时最方便，从字符串中读出的第1个数值一定是根节点的值。
//
// 尽管空节点通常没有在图上画出来，但它们对树的结构是至关重要的。因此，应该把空节点序列化成一个特殊的字符串。
func serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := serialize(root.left)
	right := serialize(root.right)
	return strings.Join([]string{strconv.Itoa(root.value), left, right}, ",")
}

func deserialize(str string) *TreeNode {
	var idx int
	strs := strings.Split(str, ",")
	return deserializeDFS(strs, &idx)
}

func deserializeDFS(strs []string, idx *int) *TreeNode {
	str := strs[*idx]
	*idx++
	if str == "#" {
		return nil
	}
	node := NewTreeNode(strToInt(str))
	node.left = deserializeDFS(strs, idx)
	node.right = deserializeDFS(strs, idx)
	return node
}

func strToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// ---------- 从根节点到叶节点的路径数字之和 ----------
func sumNumber(root *TreeNode) int {
	return sumNumberDFS(root, 0)
}

func sumNumberDFS(root *TreeNode, path int) int {
	if root == nil {
		// 节点不符合要求，直接返回0
		return 0
	}
	path = path*10 + root.value
	if root.left == nil && root.right == nil {
		// 表示当前节点是叶子结点，整个数字组合完成，直接返回
		return path
	}
	return sumNumberDFS(root.left, path) + sumNumberDFS(root.right, path)
}

// ---------- 向下的路径节点值之和 ----------
func pathSum(root *TreeNode, sum int) int {
	set := make(map[int]int)
	// key 累加的节点值之和
	// value 累加值出现的次数
	set[0] = 1
	return pathSumDFS(root, set, sum, 0)
}

// pathSum 表示路径中节点值的累加和
func pathSumDFS(root *TreeNode, set map[int]int, sum, pathSum int) int {
	if root == nil {
		return 0
	}

	pathSum += root.value
	count := set[pathSum-sum]
	set[pathSum]++
	count += pathSumDFS(root.left, set, sum, pathSum)
	count += pathSumDFS(root.right, set, sum, pathSum)
	set[pathSum] = set[pathSum] - 1
	return count
}

// ---------- 节点值之和最大的路径 ----------
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	maxPathSumDFS(root, &maxSum)
	return maxSum
}

func maxPathSumDFS(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	maxSumLeft := math.MinInt
	left := int(math.Max(0, float64(maxPathSumDFS(root.left, &maxSumLeft))))

	maxSumRight := math.MinInt
	right := int(math.Max(0, float64(maxPathSumDFS(root.right, &maxSumRight))))

	// 求出左右子树中路径节点值之和的最大值
	*maxSum = int(math.Max(float64(maxSumLeft), float64(maxSumRight)))
	// 求出经过根节点的路径节点值之和的最大值
	*maxSum = int(math.Max(float64(*maxSum), float64(root.value+left+right)))

	// 函数的返回值是经过当前节点 root 并前往其左子树或右子树的路径的节点值之和的最大值。
	// 它的父节点要根据这个返回值求路径的节点值之和。
	// 由于同时经过左右子树的路径不能经过父节点，因此返回值是变量 left 与 right 的较大值加上当前节点 root 的值。
	return root.value + int(math.Max(float64(left), float64(right)))
}

func searchBST(root *TreeNode, value int) *TreeNode {
	cur := root
	for cur != nil {
		if cur.value == value {
			break
		}
		if value < cur.value {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return cur
}

// ----------- 展平二叉搜索树 ----------
// 变量 prev 表示前一个遍历到的节点。在遍历到当前节点 cur 时，
// 把变量 prev 的右子节点的指针指向 cur，并将 cur指向左子节点的指针设置为空
// 展平之后的二叉搜索树的根节点是值最小的节点，因此也是中序遍历第1个被遍历到的节点
func increasingBST(root *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	cur := root
	var prev, first *TreeNode

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.right = cur
		} else {
			// cur 是最左边的节点，
			// 也就是新的树的根节点
			first = cur
		}
		prev = cur
		cur.left = nil // 当前节点的左节点置为空
		cur = cur.right
	}
	return first
}

// ---------- 二叉搜索树的下一个节点 ----------
func inodeSuccessorV1(root, p *TreeNode) *TreeNode {
	stack := make([]*TreeNode, 0)
	cur := root
	var found bool
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if found {
			break
		} else if p == cur {
			found = true
		}
		cur = cur.right
	}
	return cur
}

// 下一个节点的值，肯定不会小于p节点的值，而且是所有大于或等于p节点的值中最小的一个。
// 根据二叉搜索树超找的思路，比较每一个节点的值和p节点的值。
func inodeSuccessorV2(root, p *TreeNode) *TreeNode {
	cur := root
	var res *TreeNode
	for cur != nil {
		if cur.value > p.value {
			// 当前节点大于目标节点，当前节点可能是目标节点的下一个节点
			// 先记录一下结果，继续往左边走
			res = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return res
}

// ---------- 所有大于或等于节点的值之和 ----------
// 对于二叉搜索树：
//   - 中序遍历得到的是递增的结果，即所有小于当前节点的节点都已经访问过了
//   - 那么反向的中序遍历，即先访问右子树，在访问根节点，在访问左子树得到的就是递减的结果，
//     所有大于当前节点的节点都已经访问过了。
func convertBST(root *TreeNode) *TreeNode {
	cur := root
	stack := make([]*TreeNode, 0)
	var sum int
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.right
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum += cur.value // 大于等于当前节点的值，因此节点自己也要算上
		cur.value = sum
		cur = cur.left
	}
	return root
}

type iterator interface {
	next() int
	hasNext() bool
}

// ---------- 二叉搜索树迭代器 ----------
type BSTIterator struct {
	cur   *TreeNode
	stack []*TreeNode
}

// NewBSTIterator 输入二叉搜索树的根节点初始化该迭代器
func NewBSTIterator(root *TreeNode) *BSTIterator {
	return &BSTIterator{
		cur:   root,
		stack: make([]*TreeNode, 0),
	}
}

// next 返回二叉搜索树中下一个最小的节点的值
func (it *BSTIterator) next() int {
	for it.cur != nil {
		it.stack = append(it.stack, it.cur)
		it.cur = it.cur.left
	}
	it.cur = it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	res := it.cur.value
	it.cur = it.cur.right
	return res
}

// hasNext 返回二叉搜索树是否还有下一个节点
func (it *BSTIterator) hasNext() bool {
	return it.cur != nil || len(it.stack) != 0
}

// ---------- 二叉搜索树种两个节点的值之和 ----------
// 使用哈希表保存一下访问过的节点。
func findTargetV1(root *TreeNode, k int) bool {
	set := make(map[int]bool)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if set[k-cur.value] {
			return true
		} else {
			set[cur.value] = true
		}
		cur = cur.right
	}
	return false
}

// 利用二叉搜索树的特性，使用双指针。
// 双指针判断排序数组中是否存在和为K的值，一个在左边，一个在右边。
// 1. 两数之和大于K，右边的往左
// 2. 两数之和小于K，左边的向右。
// 二叉搜索树也可以看成是一个排序数组。
func findTargetV2(root *TreeNode, k int) bool {
	iterator := NewBSTIterator(root)
	iteratorReversed := NewBSTIteratorReversed(root)
	left := iterator.next()
	right := iteratorReversed.next()
	for left != right {
		if left+right == k {
			return true
		}
		if left+right > k {
			right = iteratorReversed.next()
		} else {
			left = iterator.next()
		}
	}
	return false
}

type BSTIteratorReversed struct {
	cur   *TreeNode
	stack []*TreeNode
}

func NewBSTIteratorReversed(root *TreeNode) *BSTIteratorReversed {
	return &BSTIteratorReversed{
		cur:   root,
		stack: make([]*TreeNode, 0),
	}
}

func (it *BSTIteratorReversed) next() int {
	for it.cur != nil {
		it.stack = append(it.stack, it.cur)
		it.cur = it.cur.right
	}
	it.cur = it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	res := it.cur.value
	it.cur = it.cur.left
	return res
}

func (it *BSTIteratorReversed) hasNext() bool {
	return it.cur != nil || len(it.stack) != 0
}

// ---------- 值和下标之差都在给定的范围内 ----------
func containersNearbyAlmostDuplicateV1(nums []int, k, t int) bool {
	for i := 0; i < len(nums)-k; i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i > k {
				break
			}
			if int(math.Abs(float64(nums[i]-nums[j]))) <= t {
				return true
			}
		}
	}
	return false
}

func containersNearbyAlmostDuplicateV3(nums []int, k, t int) bool {
	buckets := make(map[int]int)
	bucketSize := t + 1
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		id := getBucketID(num, bucketSize)
		_, ok1 := buckets[id]
		v2, ok2 := buckets[id-1]
		v3, ok3 := buckets[id+1]
		if ok1 || ok2 && v2+t >= num || ok3 && v3-t <= num {
			return true
		}
		buckets[id] = num
		if i >= k {
			delete(buckets, getBucketID(nums[i-k], bucketSize))
		}
	}
	return false
}

func getBucketID(num, bucketSize int) int {
	if num >= 0 {
		return num / bucketSize
	} else {
		return (num + 1) / (bucketSize - 1)
	}
}

// ---------- 日程表 ------------
type MyCalendar struct {
	index  []int
	events map[int]int
}

func NewMyCalendar() *MyCalendar {
	return &MyCalendar{
		index:  make([]int, 25),
		events: make(map[int]int),
	}
}

// 如果待添加的事项占用的时间区间是[m，n），就需要找出开始时间小于m的所有事项中开始最晚的一个，
// 以及开始时间大于m的所有事项中开始最早的一个。如果待添加的事项和这两个事项都没有重叠，那么该事项可以添加在日程表中。
func (c *MyCalendar) book(start, end int) bool {
	var preStart, nextStart int
	for i, j := 1, 24; i < 25 && j > 0; {
		if v := c.index[i]; v != 0 && v > start {
			nextStart = i
		}
		if v := c.index[j]; v != 0 && v < start {
			preStart = j
		}
		i++
		j--
	}
	if preStart == 0 && nextStart == 0 || // empty calendar
		nextStart == 0 && c.events[preStart] <= start || // empty next calendar and preEnd is before start
		preStart == 0 && nextStart >= end || // empty pre start calendar and next start is after end
		c.events[preStart] <= start && nextStart >= end { // preEnd is before start and next start is after end
		c.index[start] = start
		c.events[start] = end
		return true
	}
	return false
}
