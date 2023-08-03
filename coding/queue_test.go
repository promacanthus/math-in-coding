package coding

import (
	"reflect"
	"testing"
)

func TestMovingAverage_Next(t *testing.T) {
	tests := []struct {
		name   string
		fields *MovingAverage
		args   []int
		want   []float64
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: NewMovingAverage(3),
			args:   []int{1, 2, 3, 4},
			want:   []float64{1, 1.5, 2, 3},
		},
	}
	for _, tt := range tests {
		for i, arg := range tt.args {
			t.Run(tt.name, func(t *testing.T) {
				if got := tt.fields.Next(arg); got != tt.want[i] {
					t.Errorf("MovingAverage.Next() = %v, want %v", got, tt.want[i])
				}
			})
		}
	}
}

func TestRecentAverage_ping(t *testing.T) {
	tests := []struct {
		name   string
		fields *RecentAverage
		args   []int
		want   []int
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: RecentCounter(3000),
			args:   []int{1, 10, 3001, 3002},
			want:   []int{1, 2, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, time := range tt.args {
				if got := tt.fields.ping(time); got != tt.want[i] {
					t.Errorf("RecentAverage.ping() = %v, want %v", got, tt.want[i])
				}
			}

		})
	}
}

func TestBFS(t *testing.T) {
	eight := NewTreeNode(8)
	six := NewTreeNode(6)
	ten := NewTreeNode(10)
	eight.AddChild(six, ten)
	six.AddChild(NewTreeNode(5), NewTreeNode(7))
	ten.AddChild(NewTreeNode(9), NewTreeNode(11))

	type args struct {
		root *TreeNode
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
				root: eight,
			},
			want: []int{8, 6, 10, 5, 7, 9, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BFS(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BFS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCBTInserter_Insert(t *testing.T) {
	one := NewTreeNode(1)
	two := NewTreeNode(2)
	three := NewTreeNode(3)
	four := NewTreeNode(4)
	one.AddChild(two, three)
	two.AddChild(four, NewTreeNode(5))
	three.AddChild(NewTreeNode(6), nil)

	cbt := NewCBTInserter(one)

	tests := []struct {
		name   string
		fields *CBTInserter
		args   []int
		want   []*TreeNode
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: cbt,
			args:   []int{7, 8, 9},
			want:   []*TreeNode{three, four, four},
		},
	}
	for _, tt := range tests {
		for i, arg := range tt.args {
			t.Run(tt.name, func(t *testing.T) {
				if got := cbt.Insert(arg); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("CBTInserter.Insert() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestLargestValues(t *testing.T) {
	three := NewTreeNode(3)
	four := NewTreeNode(4)
	two := NewTreeNode(2)
	three.AddChild(four, two)
	four.AddChild(NewTreeNode(5), NewTreeNode(1))
	two.AddChild(nil, NewTreeNode(9))

	type args struct {
		root *TreeNode
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
				root: three,
			},
			want: []int{3, 4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestValuesV1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LargestValuesV1() = %v, want %v", got, tt.want)
			}
			if got := LargestValuesV2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LargestValuesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findBottomLeftValue(t *testing.T) {
	eight := NewTreeNode(8)
	six := NewTreeNode(6)
	ten := NewTreeNode(10)
	eight.AddChild(six, ten)
	six.AddChild(NewTreeNode(5), NewTreeNode(7))
	ten.AddChild(NewTreeNode(9), NewTreeNode(11))

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
				root: eight,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rightSideView(t *testing.T) {
	eight := NewTreeNode(8)
	six := NewTreeNode(6)
	ten := NewTreeNode(10)
	eight.AddChild(six, ten)
	six.AddChild(NewTreeNode(5), NewTreeNode(7))
	type args struct {
		root *TreeNode
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
				root: eight,
			},
			want: []int{8, 10, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}

