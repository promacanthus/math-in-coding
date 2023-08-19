package coding

import (
	"testing"
)

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
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
				cost: []int{1, 100, 1, 1, 100, 1},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairsTopDown(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairsTopDown() = %v, want %v", got, tt.want)
			}
			if got := minCostClimbingStairsBottomUp(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairsBottomUp() = %v, want %v", got, tt.want)
			}
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robTopDown(t *testing.T) {
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
				nums: []int{2, 3, 4, 5, 3},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robTopDown(tt.args.nums); got != tt.want {
				t.Errorf("robTopDown() = %v, want %v", got, tt.want)
			}
			if got := robBottomUp(tt.args.nums); got != tt.want {
				t.Errorf("robBottomUp() = %v, want %v", got, tt.want)
			}
			if got := robBottomUpOptimized(tt.args.nums); got != tt.want {
				t.Errorf("robBottomUpOptimized() = %v, want %v", got, tt.want)
			}
			if got := robWithTwoDP(tt.args.nums); got != tt.want {
				t.Errorf("robUseTwoDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rob(t *testing.T) {
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
				nums: []int{2, 3, 4, 5, 3},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
			if got := robCycle(tt.args.nums); got != tt.want {
				t.Errorf("robCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostTopDown(t *testing.T) {
	type args struct {
		cost [][]int
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
				cost: [][]int{{17, 2, 16}, {15, 14, 5}, {13, 3, 1}},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostTopDown(tt.args.cost); got != tt.want {
				t.Errorf("minCostTopDown() = %v, want %v", got, tt.want)
			}
			if got := minCostBottomUp(tt.args.cost); got != tt.want {
				t.Errorf("minCostBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFlipsMonoIncrTopDown(t *testing.T) {
	type args struct {
		str string
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
				str: "00110",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlipsMonoIncrTopDown(tt.args.str); got != tt.want {
				t.Errorf("minFlipsMonoIncrTopDown() = %v, want %v", got, tt.want)
			}
			if got := minFlipsMonoIncrBottomUp(tt.args.str); got != tt.want {
				t.Errorf("minFlipsMonoIncrBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFibonacci(t *testing.T) {
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
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFibonacci(tt.args.nums); got != tt.want {
				t.Errorf("findFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCut(t *testing.T) {
	type args struct {
		s string
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
				s: "aaba",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCut(tt.args.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		s1 string
		s2 string
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
				s1: "abcde",
				s2: "badfe",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				s1: "abcd",
				s2: "badfe",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
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
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
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
				s: "appplep",
				t: "apple",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniquePaths(t *testing.T) {
	type args struct {
		n int
		m int
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
				n: 3,
				m: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsTopDown(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("uniquePathsTopDown() = %v, want %v", got, tt.want)
			}
			if got := uniquePathsBottomUp(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("uniquePathsBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minPathSum(t *testing.T) {
	type args struct {
		matrix [][]int
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
				matrix: [][]int{
					{1, 3, 1},
					{2, 5, 2},
					{3, 4, 1},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSumBottomUp(tt.args.matrix); got != tt.want {
				t.Errorf("minPathSumBottomUp() = %v, want %v", got, tt.want)
			}
			if got := minPathSumTopDown(tt.args.matrix); got != tt.want {
				t.Errorf("minPathSumTopDown() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
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
				triangle: [][]int{
					{2},
					{3, 4},
					{6, 5, 7},
					{4, 1, 8, 3},
				},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
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
				nums: []int{3, 4, 1},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findTargetSumWays(t *testing.T) {
	type args struct {
		nums []int
		s    int
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
				nums: []int{2, 2, 2},
				s:    2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWays(tt.args.nums, tt.args.s); got != tt.want {
				t.Errorf("findTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
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
				coins:  []int{1, 3, 9, 10},
				target: 15,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.target); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_permutationSum(t *testing.T) {
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
				nums:   []int{1, 2, 3},
				target: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutationSum(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("permutationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
