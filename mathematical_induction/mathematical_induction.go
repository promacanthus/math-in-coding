package mathematical_induction

import "math"

type Result struct {
	currentWheat int64
	totalWheat   int64
}

// Prove is used to prove the mathematical proposition that the total number of wheat grains on a chessboard,
// when each square is filled with double the number of grains as the previous square, is equal to 2^n - 1,
// where n is the number of squares on the chessboard. It takes in an integer k and a result struct, which
// contains the current and total number of wheat grains. It recursively calls itself with k-1 until k=1,
// where it checks if the proposition holds true for n=1. If it does, it sets the current and total number of
// wheat grains to 1 and returns true. If not, it returns false. If the proposition holds true for n=k-1, the
// function calculates the current and total number of wheat grains for n=k and checks if the proposition holds
// true for n=k. If both the previous and current propositions hold true, the function returns true. Otherwise,
// it returns false. Overall, the function is used to prove a mathematical proposition using recursion and mathematical calculations.
func Prove(k int, result *Result) bool {
	// 证明 n=1 命题成立
	if k == 1 {
		if (math.Pow(2, 1) - 1) == 1 {
			result.currentWheat = 1
			result.totalWheat = 1
			return true
		} else {
			return false
		}
	} else {
		// 如果 n=k-1 命题成立，证明 n=k 命题也成立
		previous := Prove(k-1, result)
		result.currentWheat *= 2
		result.totalWheat += result.currentWheat
		current := false
		if (math.Pow(2, float64(k)) - 1) == float64(result.totalWheat) {
			current = true
		}
		if previous && current {
			return true
		}
		return false
	}
}
