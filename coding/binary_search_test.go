package coding

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
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
				nums:   []int{1, 3, 6, 8},
				target: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 6, 8},
				target: 3,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsertV1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsertV1() = %v, want %v", got, tt.want)
			}
			if got := searchInsertV2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsertV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_peakIndexMountainArray(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 3, 5, 4, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexMountainArrayV1(tt.args.nums); got != tt.want {
				t.Errorf("peakIndexMountainArrayV1() = %v, want %v", got, tt.want)
			}
			if got := peakIndexMountainArrayV2(tt.args.nums); got != tt.want {
				t.Errorf("peakIndexMountainArrayV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 1, 2, 2, 3, 4, 4, 5, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNonDuplicateV1(tt.args.nums); got != tt.want {
				t.Errorf("singleNonDuplicateV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pickIndex(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{1, 2, 3, 4},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickIndexV1(tt.args.nums); got != tt.want {
				t.Errorf("pickIndexV1() = %v, want %v", got, tt.want)
			}
			if got := pickIndexV2(tt.args.nums); got != tt.want {
				t.Errorf("pickIndexV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mySqrt(t *testing.T) {
	type args struct {
		n int
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
				n: 4,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 18,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.n); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minEatingSpeed(t *testing.T) {
	type args struct {
		piles []int
		H     int
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
				piles: []int{3, 6, 7, 11},
				H:     8,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.args.piles, tt.args.H); got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
