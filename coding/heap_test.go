package coding

import (
	"reflect"
	"testing"
)

func TestKthLargest_add(t *testing.T) {
	tests := []struct {
		name   string
		fields KthLargest
		nums   []int
		k      int
		inputs []int
		wants  []int
	}{
		// TODO: Add test cases.
		{
			name:   "",
			nums:   []int{4, 5, 8, 2},
			k:      3,
			inputs: []int{3, 5},
			wants:  []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewKthLargest(tt.nums, tt.k)
			for i, input := range tt.inputs {
				if got := h.add(input); got != tt.wants[i] {
					t.Errorf("KthLargest.add() = %v, want %v", got, tt.wants[i])
				}
			}

		})
	}
}

func Test_topKFrequent(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				nums: []int{1, 2, 2, 1, 3, 1},
				k:    2,
			},
			want: []int{2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
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
				nums1: []int{1, 5, 13, 21},
				nums2: []int{2, 4, 9, 15},
				k:     3,
			},
			want: [][]int{{5, 2}, {1, 2}, {1, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
