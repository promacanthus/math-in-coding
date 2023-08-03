package coding

import (
	"reflect"
	"testing"
)

func Test_pruneTree(t *testing.T) {
	one := NewTreeNode(1)
	two := NewTreeNode(0)
	three := NewTreeNode(0)
	one.AddChild(two, three)
	two.AddChild(NewTreeNode(0), NewTreeNode(0))
	three.AddChild(NewTreeNode(0), NewTreeNode(1))

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: one,
			},
			want: one,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pruneTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pruneTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateSerAndDesTree() (*TreeNode, *TreeNode) {
	treeOneOne := NewTreeNode(6)
	treeOneTwo := NewTreeNode(6)
	treeOneThree := NewTreeNode(6)
	treeOneOne.AddChild(treeOneTwo, treeOneThree)
	treeOneTwo.AddChild(NewTreeNode(6), NewTreeNode(6))

	treeTwoOne := NewTreeNode(6)
	treeTwoTwo := NewTreeNode(6)
	treeTwoThree := NewTreeNode(6)
	treeTwoOne.AddChild(treeTwoTwo, treeTwoThree)
	treeTwoThree.AddChild(NewTreeNode(6), NewTreeNode(6))

	return treeOneOne, treeTwoOne
}

func Test_serialize(t *testing.T) {
	tn, tn2 := generateSerAndDesTree()
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: tn,
			},
			want: "6,6,6,#,#,6,#,#,6,#,#",
		},
		{
			name: "",
			args: args{
				root: tn2,
			},
			want: "6,6,#,#,6,6,#,#,6,#,#",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deserialize(t *testing.T) {
	tn, tn2 := generateSerAndDesTree()
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				str: "6,6,6,#,#,6,#,#,6,#,#",
			},
			want: tn,
		},
		{
			name: "",
			args: args{
				str: "6,6,#,#,6,6,#,#,6,#,#",
			},
			want: tn2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserialize(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumNumber(t *testing.T) {
	three := NewTreeNode(3)
	nine := NewTreeNode(9)
	zero := NewTreeNode(0)
	three.AddChild(nine, zero)
	nine.AddChild(NewTreeNode(5), NewTreeNode(1))
	zero.AddChild(nil, NewTreeNode(2))

	type args struct {
		root *TreeNode
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
				root: three,
			},
			want: 1088,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumber(tt.args.root); got != tt.want {
				t.Errorf("sumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pathSum(t *testing.T) {
	five := NewTreeNode(5)
	two := NewTreeNode(2)
	four := NewTreeNode(4)
	five.AddChild(two, four)
	two.AddChild(NewTreeNode(1), NewTreeNode(6))
	four.AddChild(NewTreeNode(3), NewTreeNode(7))
	type args struct {
		root *TreeNode
		sum  int
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
				root: five,
				sum:  8,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.sum); got != tt.want {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxPathSum(t *testing.T) {
	negativeNine := NewTreeNode(-9)
	twenty := NewTreeNode(20)
	fifty := NewTreeNode(15)
	negativeNine.AddChild(NewTreeNode(4), twenty)
	twenty.AddChild(fifty, NewTreeNode(7))
	fifty.AddChild(NewTreeNode(-3), nil)

	type args struct {
		root *TreeNode
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
				root: negativeNine,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_increasingBST(t *testing.T) {
	four := NewTreeNode(4)
	two := NewTreeNode(2)
	five := NewTreeNode(5)
	one := NewTreeNode(1)
	four.AddChild(two, five)
	two.AddChild(one, NewTreeNode(3))
	five.AddChild(nil, NewTreeNode(6))

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: four,
			},
			want: one,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inodeSuccessor(t *testing.T) {
	eight := NewTreeNode(8)
	six := NewTreeNode(6)
	ten := NewTreeNode(10)
	nine := NewTreeNode(9)
	eight.AddChild(six, ten)
	six.AddChild(NewTreeNode(5), NewTreeNode(7))
	ten.AddChild(nine, NewTreeNode(11))
	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				root: eight,
				p:    eight,
			},
			want: nine,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inodeSuccessorV1(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inodeSuccessorV1() = %v, want %v", got, tt.want)
			}
			if got := inodeSuccessorV2(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inodeSuccessorV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBSTIterator(t *testing.T) {
	two := NewTreeNode(2)
	two.AddChild(NewTreeNode(1), NewTreeNode(3))

	it := NewBSTIterator(two)
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
		{
			name: "next",
			want: 1,
		},
		{
			name: "next",
			want: 2,
		},
		{
			name: "next",
			want: 3,
		},
		{
			name: "next",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if it.hasNext() {
				if got := it.next(); got != tt.want {
					t.Errorf("BSTIterator.next() = %v, want %v", got, tt.want)
				}
			} else {
				t.Logf("BSTIterator.hasNext() = %v", false)
			}
		})
	}
}

func Test_findTarget(t *testing.T) {
	eight := NewTreeNode(8)
	six := NewTreeNode(6)
	ten := NewTreeNode(10)
	eight.AddChild(six, ten)
	six.AddChild(NewTreeNode(5), NewTreeNode(7))
	ten.AddChild(NewTreeNode(9), NewTreeNode(11))

	type args struct {
		root *TreeNode
		k    int
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
				root: eight,
				k:    12,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				root: eight,
				k:    22,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetV1(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTargetV1() = %v, want %v", got, tt.want)
			}
			if got := findTargetV2(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTargetV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containersNearbyAlmostDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
		t    int
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
				nums: []int{1, 2, 3, 1},
				k:    3,
				t:    0,
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 5, 9, 1, 5, 9},
				k:    2,
				t:    3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containersNearbyAlmostDuplicateV1(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containersNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}

			if got := containersNearbyAlmostDuplicateV3(tt.args.nums, tt.args.k, tt.args.t); got != tt.want {
				t.Errorf("containersNearbyAlmostDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCalendar_book(t *testing.T) {
	c := NewMyCalendar()
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		c    *MyCalendar
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			c:    c,
			args: args{
				start: 10,
				end:   20,
			},
			want: true,
		},
		{
			name: "",
			c:    c,
			args: args{
				start: 15,
				end:   25,
			},
			want: false,
		},
		{
			name: "",
			c:    c,
			args: args{
				start: 20,
				end:   30,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.book(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MyCalendar.book() = %v, want %v", got, tt.want)
			}
		})
	}
}
