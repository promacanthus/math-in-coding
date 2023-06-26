package recursion

import (
	"fmt"
	"sort"
)

var rewards = []int{1, 2, 5, 10} // 四种面额的纸币

// This is a recursive function that takes in a total reward and a slice of previous rewards
// and generates all possible combinations of rewards that add up to the total reward. If the
// total reward is 0, it prints out the current combination of rewards. If the total reward
// is negative, it returns and stops the recursion. Otherwise, it iterates through a list of
// rewards and creates a new combination of rewards by appending the current reward to the
// previous rewards slice. It then calls itself with the new combination and the remaining
// total reward after subtracting the current reward. This process continues until all
// possible combinations have been generated.
func getReward(totalReward int, result []int) {
	switch {
	case totalReward == 0:
		fmt.Println(result)
	case totalReward < 0:
		return
	default:
		for _, reward := range rewards {
			newRes := result
			newRes = append(newRes, reward)
			getReward(totalReward-reward, newRes)
		}
	}
}

// factorization takes an integer 'num' and a slice 'result'. It recursively finds the prime factorization of 'num'
// and appends the factors to 'result'. If 'num' is equal to 1, the function prints 'result'. If 'num' is not equal
// to 1, the function loops through all integers from 1 to 'num' and checks if the integer is a factor of 'num'.
// If it is, the function creates a new slice 'newRes' by appending the factor to 'result' and calls itself with 'num/i'
// and 'newRes'. If the factor is 1 and 'result' already contains 1, the loop continues without appending 1 to 'newRes'.
func factorization(num int, result []int) {
	if num < 1 {
		return
	}
	if num == 1 {
		if !containsOne(result) {
			result = append(result, 1)
		}
		fmt.Println(result)
	} else {
		var newResult []int
		for i := 1; i <= num; i++ {
			if num%i == 0 {
				if i == 1 && containsOne(result) {
					continue
				}
				newResult = make([]int, len(result))
				copy(newResult, result)
				newResult = append(newResult, i)
				factorization(num/i, newResult)
			}
		}
	}
}

// containsOne checks if the provided slice of integers contains the number 1. It does
// this by first sorting the slice using the Ints function from the sort package. It
// then uses the SearchInts function from the same package to search for the number 1
// within the sorted slice. If the function returns a value less than the length of the
// slice, it means that the number 1 was found and the function returns true. Otherwise, it returns false.
func containsOne(result []int) bool {
	res := result
	sort.Ints(res)
	return sort.SearchInts(res, 1) != len(res)
}

// MergeSort implements the Merge Sort algorithm to sort an input slice of integers.
// It first checks if the input slice is empty or nil, in which case it returns an
// empty slice. If the input slice has only one element, it is already sorted and
// returned as is. Otherwise, it recursively divides the input slice into two halves
// and sorts them separately using Merge Sort. Finally, the two sorted halves are
// merged together using the merge() function and returned as the sorted output.
// This algorithm has a time complexity of O(n log n) and is a popular choice
// for sorting large datasets.
func MergeSort(in []int) []int {
	if in == nil || len(in) == 0 {
		return []int{}
	}

	if len(in) == 1 {
		return in
	}

	mid := len(in) / 2
	left := in[0:mid]
	right := in[mid:]
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

// merge takes in two slices of integers, "left" and "right", and returns a merged
// slice of integers in sorted order. If either "left" or "right" is nil, it is
// initialized as an empty slice. It initializes an empty slice called "merged"
// and two variables "i" and "j" to keep track of the indices of the left and right
// slices, respectively. It iterates through the left and right slices, comparing
// the values at the current indices and appending the smaller value to "merged".
// Once one of the slices has been fully iterated through, the remaining values
// from the other slice are appended to "merged". Finally, it returns the merged slice.
func merge(left, right []int) []int {
	if left == nil {
		left = make([]int, 0)
	}
	if right == nil {
		right = make([]int, 0)
	}

	merged := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	if i < len(left) {
		merged = append(merged, left[i:]...)
	}
	if j < len(right) {
		merged = append(merged, right[j:]...)
	}
	return merged
}
