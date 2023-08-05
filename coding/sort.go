package coding

import (
	"container/heap"
	"math"
	"math/rand"
	"sort"
)

// ---------- 合并区间 ----------
func merge(intervals [][]int) [][]int {
	// sort the intervals by interval origin
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	for i := 0; i < len(intervals); {
		tmp := make([]int, 2)
		tmp[0] = intervals[i][0]
		tmp[1] = intervals[i][1]
		j := i + 1
		for ; j < len(intervals) && intervals[j][0] <= tmp[1]; j++ {
			tmp[1] = int(math.Max(float64(tmp[1]), float64(intervals[j][1])))
		}
		res = append(res, tmp)
		i = j
	}
	return res
}

func sortArray(nums []int) []int {
	min, max := math.MaxInt, math.MinInt
	for _, num := range nums {
		min = int(math.Min(float64(num), float64(min)))
		max = int(math.Max(float64(num), float64(max)))
	}
	counts := make([]int, max-min+1)
	for _, num := range nums {
		counts[num-min]++
	}

	res := make([]int, len(nums))
	for i, num := 0, min; num <= max; num++ {
		for counts[num-min] > 0 {
			res[i] = num
			i++
			counts[num-min]--
		}
	}
	return res
}

// ---------- 数组相对排序 ----------
func relativeSortArray(arr1, arr2 []int) []int {
	counts := make([]int, 1001)
	for _, num := range arr1 {
		counts[num]++
	}
	i := 0
	for _, num := range arr2 {
		for counts[num] > 0 {
			arr1[i] = num
			i++
			counts[num]--
		}
	}
	for num := 0; num < len(counts); num++ {
		for counts[num] > 0 {
			arr1[i] = num
			i++
			counts[num]--
		}
	}
	return arr1
}

func quickSort(nums []int, start, end int) {
	if end > start {
		pivot := partition(nums, start, end)
		quickSort(nums, start, pivot-1)
		quickSort(nums, pivot+1, end)
	}
}

func partition(nums []int, start, end int) int {
	// randomly select a pivot
	random := rand.Intn(end-start+1) + start
	// swap the pivot and the end
	nums[random], nums[end] = nums[end], nums[random]
	// initialize p1 and p2
	p1, p2 := start-1, start
	for ; p2 < end; p2++ {
		// move p2 to find the small one less than the pivot
		if nums[p2] < nums[end] {
			p1++
			// swap p1 and p2
			nums[p2], nums[p1] = nums[p1], nums[p2]
		}
	}
	p1++
	// move back the pivot
	nums[p1], nums[end] = nums[end], nums[p1]
	return p1
}

// 数组中第K大的数字
func findKthLargest(nums []int, k int) int {
	// target is the index of the kth largest element of a sorted array.
	target := len(nums) - k

	start, end := 0, len(nums)-1
	index := partition(nums, start, end)
	// use binary search to check if the index is equal to the target.
	for index != target {
		if index > target {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(nums, start, end)
	}
	return nums[index]
}

func mergeSortArray(nums []int) []int {
	length := len(nums)
	src, dst := nums, make([]int, length)
	// seg represents the elements in the array that will be merged in every iteration.
	for seg := 1; seg < length; seg += seg {
		// start is the index of the first element to be merged in the iteration.
		for start := 0; start < length; start += seg * 2 {
			mid := int(math.Min(float64(start+seg), float64(length)))
			end := int(math.Min(float64(start+seg*2), float64(length)))
			// i is the index of the first array.
			// j is the index of the second array.
			for i, j, k := start, mid, start; i < mid || j < end; k++ {
				if j == end || // retrieves all elements in the two candidate arrays.
					(i < mid && src[i] < src[j]) {
					dst[k] = src[i]
					i++
				} else {
					dst[k] = src[j]
					j++
				}
			}
		}
		tmp := src
		src = dst // dst is the sorted array.
		dst = tmp
	}
	return src
}

func mergeSort(nums []int) []int {
	dst := make([]int, len(nums))
	ms(nums, dst, 0, len(nums))
	return dst
}

func ms(src, dst []int, start, end int) {
	if start+1 >= end {
		return
	}

	mid := (start + end) / 2
	ms(src, dst, start, mid)
	ms(src, dst, mid, end)

	for i, j, k := start, mid, start; i < mid || j < end; k++ {
		if j == end || i < mid && src[i] < src[j] {
			dst[k] = src[i]
			i++
		} else {
			dst[k] = src[j]
			j++
		}
	}
}

// ---------- 链表排序 ----------
func sortList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	head1 := head
	head2 := split(head)
	head1 = sortList(head1)
	head2 = sortList(head2)
	return mergeList(head1, head2)
}

func split(head *Node) *Node {
	fast, slow := head.next, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	second := slow.next
	slow.next = nil
	return second
}

func mergeList(head1 *Node, head2 *Node) *Node {
	dummy := NewNode(0)
	cur := dummy
	for head1 != nil && head2 != nil {
		if head1.value <= head2.value {
			cur.next = head1
			head1 = head1.next
		} else {
			cur.next = head2
			head2 = head2.next
		}
		cur = cur.next
	}
	if head1 != nil {
		cur.next = head1
	}
	if head2 != nil {
		cur.next = head2
	}
	return dummy.next
}

// ---------- 合并排序链表 ----------
func mergeSortedLinkedList(linkedLists []*LinkedList) *Node {
	mh := &minHeapForNode{}
	for _, l := range linkedLists {
		cur := l.head
		for cur != nil {
			heap.Push(mh, cur)
			cur = cur.next
		}
	}
	dummy := NewNode(0)
	cur := dummy
	for mh.Len() != 0 {
		cur.next = heap.Pop(mh).(*Node)
		cur = cur.next
	}
	return dummy.next
}

type minHeapForNode []*Node

func (h *minHeapForNode) Len() int           { return len(*h) }
func (h *minHeapForNode) Less(i, j int) bool { return (*h)[i].value < (*h)[j].value }
func (h *minHeapForNode) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeapForNode) Push(x any)         { *h = append(*h, x.(*Node)) }
func (h *minHeapForNode) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func mergeSortedLinkedLists(lists []*LinkedList) *Node {
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists))
}

func mergeLists(lists []*LinkedList, start, end int) *Node {
	if start+1 == end {
		return lists[start].head
	}
	mid := (start + end) / 2
	head1 := mergeLists(lists, start, mid)
	head2 := mergeLists(lists, mid, end)
	return mergeList(head1, head2)
}
