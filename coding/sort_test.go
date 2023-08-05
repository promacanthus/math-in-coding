package coding

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				intervals: [][]int{{1, 3}, {4, 5}, {8, 10}, {2, 6}, {9, 12}, {15, 18}},
			},
			want: [][]int{{1, 6}, {8, 12}, {15, 18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{2, 3, 4, 2, 3, 2, 1},
			},
			want: []int{1, 2, 2, 2, 3, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_relativeSortArray(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				arr1: []int{2, 3, 3, 7, 3, 9, 2, 1, 7, 2},
				arr2: []int{3, 2, 1},
			},
			want: []int{3, 3, 3, 2, 2, 2, 1, 7, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := relativeSortArray(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relativeSortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{3, 1, 2, 4, 5, 5, 6},
				k:    3,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				nums: []int{4, 1, 5, 6, 2, 7, 8, 3},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortArray() = %v, want %v", got, tt.want)
			}
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortList(t *testing.T) {
	three := NewNode(3)
	five := NewNode(5)
	one := NewNode(1)
	four := NewNode(4)
	two := NewNode(2)
	six := NewNode(6)

	three.next = five
	five.next = one
	one.next = four
	four.next = two
	two.next = six

	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: three,
			},
			want: one,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSortedLinkedList(t *testing.T) {
	l1 := NewLinkList([]int{1, 4, 7})
	l2 := NewLinkList([]int{2, 5, 8})
	l3 := NewLinkList([]int{3, 6, 9})

	type args struct {
		linkedLists []*LinkedList
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				linkedLists: []*LinkedList{l1, l2, l3},
			},
			want: l1.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSortedLinkedList(tt.args.linkedLists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedLinkedList() = %v, want %v", got, tt.want)
			}
			if got := mergeSortedLinkedLists(tt.args.linkedLists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSortedLinkedLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
