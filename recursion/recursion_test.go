package recursion

import (
	"reflect"
	"testing"
)

func Test_getReward(t *testing.T) {
	type args struct {
		totalReward int
		result      []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				totalReward: 10,
				result:      []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getReward(tt.args.totalReward, tt.args.result)
		})
	}
}

func Test_factorization(t *testing.T) {
	type args struct {
		num    int
		result []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				num:    8,
				result: []int{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			factorization(tt.args.num, tt.args.result)
		})
	}
}

func TestMergeSort(t *testing.T) {
	type args struct {
		in []int
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
				in: []int{3, 8, 9, 7, 0, 6, 5, 1, 2, 4},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
