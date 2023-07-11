package coding

import (
	"math"
)

// Single Linked List
type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

type LinkedList struct {
	head *Node
}

func NewLinkList(values []int) *LinkedList {
	ll := &LinkedList{}
	for _, v := range values {
		ll.AddNode(v)
	}
	return ll
}

func (l *LinkedList) AddNode(value int) {
	newNode := NewNode(value)
	if l.head == nil {
		l.head = newNode
		return
	}

	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

func (ll *LinkedList) getNodeByValue(value int) *Node {
	node := ll.head
	for node != nil && node.value != value {
		node = node.next
	}
	return node
}

// Double Linked List
type DNode struct {
	value int
	prev  *DNode
	next  *DNode
}

// Double Linked List with a chain
type DCNode struct {
	value int
	prev  *DCNode
	next  *DCNode
	child *DCNode
}

func NewDNode(value int) *DCNode {
	return &DCNode{value: value}
}

type DCLinkedList struct {
	head *DCNode
}

func NewDCLinkedList(value []int) *DCLinkedList {
	ll := &DCLinkedList{}
	for _, v := range value {
		ll.AddNode(v)
	}
	return ll
}

func (l *DCLinkedList) AddNode(value int) {
	newNode := NewDNode(value)
	if l.head == nil {
		l.head = newNode
		return
	}

	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
	newNode.prev = node
}

func (l *DCLinkedList) getNodeByValue(value int) *DCNode {
	node := l.head
	for node != nil {
		if node.value == value {
			return node
		}
		node = node.next
	}
	return nil
}

func (l *DCLinkedList) AddChild(value int, child *DCNode) {
	node := l.getNodeByValue(value)
	node.child = child
}

// ---------- 删除倒数第 k 个节点 ----------
// 双指针，两个指针之间的间隔是 k-1
// 使用哨兵节点，简化待删除的节点是 head 的情况
func removeNthFromEnd(head *Node, k int) *Node {
	dummy := &Node{}
	dummy.next = head

	p1 := dummy.next
	p2 := dummy.next
	for i := 0; i < k; i++ {
		p1 = p1.next
	}
	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	p2.next = p2.next.next
	return dummy.next
}

// ---------- 链表中环的入口节点 ----------
// 判断链表是否有环：
//
//	快慢指针可以判断链表中是否有环，如果两个指针相遇，则表明链表中存在环。
//
// 计算链表中环的长度：
//
//	两个指针之所以会相遇是因为快的指针绕环一圈追上慢的指针，因此它们相遇的节点一定是在环中。
//	从这个相遇的节点出发一边继续向前移动一边计数，当再次回到这个节点时就可以得到环中节点的数目。
//
// 寻找环的入口点：
//
//	先定义两个指针 P1 和 P2，指向链表的头节点。如果链表中的环有 n 个节点，P1 先在链表中向前移动 n 步，
//	然后 P1、P2 以相同的速度向前移动,P2 指向环的入口节点时，P1 已经围绕环走了一圈又回到了入口节点。
//
// 当 P1 重新走到环的入口点时，也表示已经走到了整个链表的尾部。
// 假设链表长度是 x+n，其中 n 是环的长度，那么入口点就在 x 的位置，
// P1 比 P2 先走 n 步，然后 P1 和 P2 一起移动，同时走 x 步时，P1 来到环的入口点，P2 来到链表尾部及环的入口点，
// 此时 P1 和 P2 相遇，这个点就是环的入口点。
func detectCycleV1(head *Node) *Node {
	node := getNodeInLoop(head)
	if node == nil {
		return node
	}

	count := 1
	for n := node.next; n != node; n = n.next {
		count++
	}

	fast := head
	for i := 0; i < count; i++ {
		fast = fast.next
	}

	slow := head
	for fast != slow {
		fast = fast.next
		slow = slow.next
	}
	return slow
}

// 如果链表中有环，快慢两个指针一定会在环中的某个节点相遇。
// - 慢的指针一次走一步，在相遇时慢的指针一共走了 k 步。
// - 快的指针一次走两步，在相遇时快的指针一共走了 2k 步。
// 因此，到相遇时快的指针比慢的指针多走了 k 步。另外，两个指针相遇时快的指针比慢的指针在环中多转了若干圈。
// 两个指针相遇时快的指针多走的步数 k 一定是环中节点的数目的整数倍，此时慢的指针走过的步数 k 也是环中节点数的整数倍。
// 所以此时慢指针的位置就相当于知道了环的长度的整数倍。
//
// 1. 让一个指针指向相遇的节点，该指针的位置是之前慢的指针走了 k 步到达的位置。
// 2. 让另一个指针指向链表的头节点，然后两个指针以相同的速度一起朝着指向下一个节点的指针移动，
// 当后面的指针到达环的入口节点时，前面的指针比它多走了 k 步，而 k 是环中节点的数目的整数倍，
// 相当于前面的指针在环中转了 k 圈后也到达环的入口节点，两个指针正好相遇，两个指针相遇的节点正好是环的入口节点。
func detectCycleV2(head *Node) *Node {
	node := getNodeInLoop(head)
	if node == nil {
		return node
	}

	p := head
	for p != node {
		p = p.next
		node = node.next
	}
	return node
}

func getNodeInLoop(head *Node) *Node {
	if head == nil || head.next == nil {
		return nil
	}

	slow := head.next // 走第一步
	fast := slow.next // 走第两步
	for fast != nil && slow != nil {
		if fast == slow {
			return slow
		}
		slow = slow.next
		fast = fast.next // 先走一步
		if fast.next != nil {
			fast = fast.next // 再走一步
		}
	}
	return nil
}

// ---------- 两个链表的第1个重合节点 ----------
// 空间换时间
// 用一个哈希表保存 其中一条链，然后遍历另一条链，存在的第一个节点就是重合的第一个节点。
// 时间复杂度：O(n+m)，空间复杂度：O(n)
func getIntersectionNodeV1(a, b *Node) *Node {
	set := make(map[*Node]struct{})
	nodeA := a
	for nodeA != nil {
		set[nodeA] = struct{}{}
		nodeA = nodeA.next
	}
	nodeB := b
	for nodeB != nil {
		if _, ok := set[nodeB]; ok {
			return nodeB
		} else {
			nodeB = nodeB.next
		}
	}
	return nil
}

// 空间换时间
// 使用两个栈，从后往前遍历两个链表，遇到的一个相同的节点之前的那个节点，就是两个链表重合的第一个节点。
// 时间复杂度：O(n+m+k)，空间复杂度：O(n+m)
func getIntersectionNodeV2(a, b *Node) *Node {
	stackA := make([]*Node, 0)
	stackB := make([]*Node, 0)
	nodeA := a
	nodeB := b
	for nodeA != nil {
		stackA = append(stackA, nodeA)
		nodeA = nodeA.next
	}
	for nodeB != nil {
		stackB = append(stackB, nodeB)
		nodeB = nodeB.next
	}

	var curA, curB, lastA *Node
	for {
		curA = stackA[len(stackA)-1]
		stackA = stackA[:len(stackA)-1]
		curB = stackB[len(stackB)-1]
		stackB = stackB[:len(stackB)-1]
		if curA == curB {
			lastA = curA
		} else {
			return lastA
		}
	}
}

// 形成环，寻找环入口
// 因为两个两边有重合的部分，所以遍历第一个链表到尾节点，然后指向第二个链表的头结点，
// 那么就以第二个链表形成了一个环，这个环的入口就是两个链表的第一个重合点。
// 时间复杂度：O(n+z*k+x)，空间复杂度：O(1)
func getIntersectionNodeV3(a, b *Node) *Node {
	lastNode := getLastElement(a)
	lastNode.next = b
	return detectCycleV2(a)
}

func getLastElement(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

// 双指针
// 先分别计算出两个链表的长度，然后将第一个指针先把长的那个链表多出来的节点先走完，
// 然后同时走两个链表，遇到的第一个相同的节点就是结果。
// 时间复杂度：O(n+m)，空间复杂度：O(1)
func getIntersectionNodeV4(a, b *Node) *Node {
	longer, shorter, delta := getLongerAndShorter(a, b)
	node1 := longer
	for i := 0; i < delta; i++ {
		node1 = node1.next
	}
	node2 := shorter
	for node1 != node2 {
		node1 = node1.next
		node2 = node2.next
	}
	return node1
}

func getLongerAndShorter(a, b *Node) (longer, shorter *Node, delta int) {
	lenA := getLinkedListLength(a)
	lenB := getLinkedListLength(b)
	delta = int(math.Abs(float64(lenA) - float64(lenB)))
	if lenA > lenB {
		longer = a
		shorter = b
	} else {
		longer = b
		shorter = a
	}
	return
}

func getLinkedListLength(head *Node) int {
	var length int
	for head != nil {
		length++
		head = head.next
	}
	return length
}

// ---------- 反转链表 ----------
// a<-b<-c...<-i<-j  k<-...
// 在调整节点 j 的 next 指针时：
// - 需要知道节点 j 本身
// - 还需要知道节点 j 的前一个节点 i，因为需要把节点 j 的 next 指针指向节点 i
// - 还需要事先保存节点 j 的下一个节点 k，以防止链表断开。
// 因此，在遍历链表逐个反转每个节点的 next 指针时需要用到 3 个指针，分别指向当前遍历到的节点、它的前一个节点及后一个节点。
// 时间复杂度：O(n)，空间复杂度：O(1)
func reserveList(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

// ---------- 链表中的数字相加 ----------
// 所有整数相加的问题都需要考虑溢出的情况。
// 将两个链表反转后，从链表的头结点开始相加就相当于两个整数从个位数开始相加。
// 同时需要考虑进位的问题。
func addTwoNumbers(a, b *Node) *Node {
	reservedA := reserveList(a)
	reservedB := reserveList(b)
	res := addTwoList(reservedA, reservedB)
	return reserveList(res)
}

func addTwoList(a, b *Node) *Node {
	dummy := NewNode(0)
	sumNode := dummy
	carry := 0
	for a != nil || b != nil {
		sum := getNodeValue(a) + getNodeValue(b) + carry
		carry, sum = getCarryAndSum(sum)
		node := NewNode(sum)
		sumNode.next = node
		sumNode = sumNode.next
		a = getNextNode(a)
		b = getNextNode(b)
	}
	if carry > 0 {
		sumNode.next = NewNode(carry)
	}
	return dummy.next
}

func getNodeValue(node *Node) int {
	if node == nil {
		return 0
	}
	return node.value
}

func getCarryAndSum(sum int) (int, int) {
	if sum >= 10 {
		return 1, sum - 10
	}
	return 0, sum
}

func getNextNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	return node.next
}

// ---------- 重排链表 ----------
// 1. 寻找链表重点，快慢双指针，快指针到尾节点时，慢指针到前半段的尾节点，慢指针的下一个节点就是后半段的头结点
// 2. 反转后半段链表
// 3. 将两个链表从头到尾依次交替串联起来，使用哨兵节点简化操作
// 一个值得注意的问题是，链表的节点总数既可能是奇数也可能是偶数。当链表的节点总数是奇数时，就要确保链表的前半段比后半段多一个节点。
func reorderList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	mid := getMiddleNode(head)
	tmp := mid.next
	mid.next = nil
	return linkNodes(head, reserveList(tmp))
}

func getMiddleNode(head *Node) *Node {
	dummy := NewNode(0)
	dummy.next = head
	slow := dummy
	fast := dummy
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next
		if fast.next != nil {
			fast = fast.next
		}
	}
	return slow
}

func linkNodes(a, b *Node) *Node {
	dummy := NewNode(0)
	head := dummy
	for a != nil && b != nil {
		head.next = a
		head = head.next
		a = a.next

		head.next = b
		head = head.next
		b = b.next
	}
	return dummy.next
}

// ---------- 回文链表 ----------
// 链表前半段和后半段的反转是相同的。
// 不管链表的节点总数是奇数还是偶数，慢指针都指向链表前半段的最后一个节点。
func isPalindromeLinkedList(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	mid := getMiddleNode(head)
	tmp := mid.next
	mid.next = nil
	tail := reserveList(tmp)

	nodeA := head
	nodeB := tail
	for nodeA != nil && nodeB != nil {
		if nodeA.value != nodeB.value {
			return false
		}
		nodeA = nodeA.next
		nodeB = nodeB.next
	}
	return true
}

// ---------- 展平多级双向链表 ----------
// 递归的处理每一个子链表，需要知道子链表的尾节点，所以递归函数的返回值是尾节点。
func flatten(head *DCNode) *DCNode {
	flattenGetTail(head)
	return head
}

func flattenGetTail(head *DCNode) *DCNode {
	var tail *DCNode
	node := head
	for node != nil {
		// 获取当前节点的下一个节点
		next := node.next
		var child, childTail *DCNode
		if node.child != nil {
			// 获取子链表的头结点
			child = node.child
			// 获取子链表的尾节点
			childTail = flattenGetTail(child)
			// 将子链表插入到当前节点 node 和 next 之间
			// node --> child...tail --> next
			node.child = nil
			node.next = child
			child.prev = node
			childTail.next = next
			if next != nil {
				next.prev = childTail
			}
			// 获取当前层级中链表的尾节点
			tail = childTail
		} else {
			// 没有子链表，则当前节点就是当前层链表的尾节点
			tail = node
		}
		// 继续遍历当前节点
		node = next
	}
	return tail
}

// ---------- 排序的循环链表 ----------
// 插入新节点有两个特殊的情况：
// 1. 新节点的值比最大值还大
// 2. 新节点的值比最小值还小
// 那么新节点都将被插入到最大值和最小值之间。
//
// 因此新节点插入的规则就是找到两个节点，如果前一个节点比新节点小，后一个节点比新节点大，
// 那么就把新节点插入到这两个节点之间，否则将新节点插入到最大值和最小值之间。
//
// 边界条件：
// 当链表为空时，插入的节点是第一个节点，next指针指向自己。
// 当链表只有一个节点是，两个节点互为对方的 next 指针。
func insert(head *Node, value int) *Node {
	newNode := NewNode(value)
	if head == nil { // 链表为空
		head = newNode
		head.next = head
	} else if head.next == nil { // 链表只有一个节点
		head.next = newNode
		newNode.next = head
	} else {
		insertCore(head, newNode)
	}
	return head
}

func insertCore(head *Node, node *Node) {
	cur := head
	next := head.next
	biggest := head
	// 寻找待插入节点介于两者之间的位置
	for !(cur.value <= node.value && next.value >= node.value) && next != head {
		cur = next
		next = next.next
		if cur.value >= biggest.value {
			biggest = cur
		}
	}
	if cur.value <= node.value && next.value >= node.value {
		cur.next = node
		node.next = next
	} else {
		// 插入最大值之后
		node.next = biggest.next
		biggest.next = node
	}
}
