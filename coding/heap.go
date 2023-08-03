package coding

import "container/heap"

type MinIntHeap []int

// Len is the number of elements in the collection.
func (h *MinIntHeap) Len() int {
	return len(*h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h *MinIntHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Swap swaps the elements with indexes i and j.
func (h *MinIntHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinIntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() any {
	old := *h
	n := len(*h)
	res := (*h)[n-1]
	*h = (old)[:n-1]
	return res
}

// ----------- 数据流的第k大数字 ----------
type KthLargest struct {
	heap *MinIntHeap
	size int
}

func NewKthLargest(nums []int, k int) *KthLargest {
	h := &KthLargest{
		heap: new(MinIntHeap),
		size: k,
	}

	for _, num := range nums {
		h.add(num)
	}
	return h
}

func (h *KthLargest) add(num int) int {
	// The heap is not full.
	if h.heap.Len() < h.size {
		heap.Push(h.heap, num)
		return (*h.heap)[0]
	}

	// The number is less than the root of the heap.
	if h.heap.Len() == h.size && num <= (*h.heap)[0] {
		return (*h.heap)[0]
	}

	// The number is greater than the root of the heap.
	// Push it into the heap and remove the root of the heap.
	heap.Push(h.heap, num)
	heap.Remove(h.heap, 0)
	return (*h.heap)[0]
}

// ---------- 出现频率最高的k个数字 ----------
func topKFrequent(nums []int, k int) []int {
	set := make(map[int]int)
	for _, num := range nums {
		set[num]++
	}

	top := &frequentHeap{
		items: []Item{},
		size:  k,
	}

	for k, v := range set {
		item := Item{
			Key:       k,
			Frequency: v,
		}
		if top.Len() < top.size {
			heap.Push(top, item)
		} else {
			if item.Frequency > top.items[0].Frequency {
				heap.Push(top, item)
				heap.Pop(top)
			}
		}
	}
	var res []int
	for _, item := range top.items {
		res = append(res, item.Key)
	}
	return res
}

type Item struct {
	Key       int
	Frequency int
}

type frequentHeap struct {
	items []Item
	size  int
}

// Len is the number of elements in the collection.
func (h *frequentHeap) Len() int {
	return len(h.items)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h *frequentHeap) Less(i int, j int) bool {
	return h.items[i].Frequency < h.items[j].Frequency
}

// Swap swaps the elements with indexes i and j.
func (h *frequentHeap) Swap(i int, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *frequentHeap) Push(x any) {
	h.items = append(h.items, x.(Item))
}

func (h *frequentHeap) Pop() any {
	old := h.items
	n := len(old)
	res := old[n-1]
	h.items = old[:n-1]
	return res
}

// ---------- 和最小的k个数对 -----------
func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	input1 := nums1[:k]
	input2 := nums2[:k]
	maxHeap := new(MaxHeap)
	for _, v1 := range input1 {
		for _, v2 := range input2 {
			sum := v1 + v2
			if maxHeap.Len() < k {
				heap.Push(maxHeap, &item{sum: sum, pair: []int{v1, v2}})
			} else if sum < (*maxHeap)[0].sum {
				heap.Push(maxHeap, &item{sum: sum, pair: []int{v1, v2}})
				heap.Pop(maxHeap)
			}
		}
	}

	res := make([][]int, 0)
	for _, i := range *maxHeap {
		res = append(res, i.pair)
	}
	return res
}

type item struct {
	sum  int
	pair []int
}

type MaxHeap []*item

// Len is the number of elements in the collection.
func (h *MaxHeap) Len() int {
	return len(*h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h *MaxHeap) Less(i int, j int) bool {
	return (*h)[i].sum > (*h)[j].sum
}

// Swap swaps the elements with indexes i and j.
func (h *MaxHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(*item))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*h = old[:n-1]
	return item
}
