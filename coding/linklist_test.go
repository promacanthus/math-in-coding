package coding

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	ll := NewLinkList([]int{1, 2, 3, 4, 5, 6})

	type args struct {
		head *Node
		n    int
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
				head: ll.head,
				n:    2,
			},
			want: ll.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_detectCycle(t *testing.T) {
	ll := NewLinkList([]int{1, 2, 3, 4, 5, 6})
	three := ll.getNodeByValue(3)
	six := ll.getNodeByValue(6)
	six.next = three

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
				head: ll.head,
			},
			want: three,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycleV1(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycleV1() = %v, want %v", got, tt.want)
			}
			if got := detectCycleV2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycleV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntersectionNode(t *testing.T) {
	l1 := NewLinkList([]int{1, 2, 3})
	l2 := NewLinkList([]int{7, 8})
	l3 := NewLinkList([]int{4, 5, 6})
	three := l1.getNodeByValue(3)
	eight := l2.getNodeByValue(8)
	three.next = l3.head
	eight.next = l3.head

	type args struct {
		a *Node
		b *Node
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
				a: l1.head,
				b: l2.head,
			},
			want: l3.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNodeV1(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNodeV1() = %v, want %v", got, tt.want)
			}
			if got := getIntersectionNodeV2(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNodeV2() = %v, want %v", got, tt.want)
			}
			if got := getIntersectionNodeV4(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNodeV4() = %v, want %v", got, tt.want)
			}
			// 这个解法会形成环，最后一个测试
			if got := getIntersectionNodeV3(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNodeV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reserverList(t *testing.T) {
	ll := NewLinkList([]int{1, 2, 3, 4, 5})
	five := ll.getNodeByValue(5)

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
				head: ll.head,
			},
			want: five,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reserveList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reserverList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addTwoNumbers(t *testing.T) {
	llA := NewLinkList([]int{9, 8, 4})
	llB := NewLinkList([]int{1, 8})
	res := NewLinkList([]int{1, 0, 0, 2})

	type args struct {
		a *Node
		b *Node
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
				a: llA.head,
				b: llB.head,
			},
			want: res.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reorderList(t *testing.T) {
	ll := NewLinkList([]int{1, 2, 3, 4, 5, 6})
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
				head: ll.head,
			},
			want: ll.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeLinkedList(t *testing.T) {
	ll := NewLinkList([]int{1, 2, 3, 3, 2, 1})
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: ll.head,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeLinkedList(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flatten(t *testing.T) {
	l1 := NewDCLinkedList([]int{1, 2, 3, 4})
	l2 := NewDCLinkedList([]int{5, 6, 7})
	l3 := NewDCLinkedList([]int{8, 9})
	l1.AddChild(2, l2.head)
	l2.AddChild(6, l3.head)

	type args struct {
		head *DCNode
	}
	tests := []struct {
		name string
		args args
		want *DCNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				head: l1.head,
			},
			want: l1.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatten(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	ll := NewLinkList([]int{1, 2, 3, 5, 6, 7})
	seven := ll.getNodeByValue(7)
	seven.next = ll.head

	type args struct {
		head  *Node
		value int
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
				head:  ll.head,
				value: 4,
			},
			want: ll.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.head, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
