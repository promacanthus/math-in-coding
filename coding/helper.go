package coding

import "math"

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func IntsMin(nums ...int) int {
	min := math.MaxInt
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func IntsMax(nums ...int) int {
	max := math.MinInt
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func InitDP(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return dp
}

func BooleanPointer(b bool) *bool {
	return &b
}
