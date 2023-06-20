package iterator

import "math"

// GetWheats calculates the total number of wheat grains in a grid of a given size.
// It starts with one grain in the first grid and doubles the number of grains in each subsequent
// grid until it reaches the given grid size. The sum of all the grains in the grid is then returned.
func GetWheats(grid int) int64 {
	var (
		sum          int64
		wheatsInGrid int64
	)

	wheatsInGrid = 1
	sum += wheatsInGrid
	for i := 2; i <= grid; i++ {
		wheatsInGrid *= 2
		sum += wheatsInGrid
	}
	return sum
}

// GetSquareRoot takes in an integer 'root' and returns the square root of 'root' using binary search.
// It takes in two more parameters 'maxTry' and 'threshold' which determine the maximum number of iterations
// and the acceptable error threshold respectively. If the 'root' is less than 1, the function returns -1.
// The function then initializes a minimum and maximum value for the binary search and iterates through
// the search process for 'maxTry' times. It calculates the square of the middle value and compares it to the 'root'.
// If the difference between the square of the middle value and the 'root' is less than or equal to the 'threshold',
// the function returns the middle value. If the square of the middle value is greater than the 'root',
// the maximum value for the search is updated to the middle value, otherwise, the minimum value is updated to the middle value.
// If the function exhausts all the iterations and does not find a square root within the acceptable threshold, it returns 0.
func GetSquareRoot(root, maxTry int, threshold float64) int {
	if root < 1 {
		return -1
	}

	min := 1
	max := root
	for i := 0; i < maxTry; i++ {
		mid := (min + max) / 2
		square := mid * mid
		delta := math.Abs(float64(square)/float64(root) - 1)
		if delta <= threshold {
			return mid
		}

		if square > root {
			max = mid
		} else {
			min = mid
		}
	}
	return 0
}

// SearchString implements binary search to find a target string in a sorted array of strings.
// It takes in two parameters, the array of strings and the target string. If the array is empty,
// it returns -1 indicating that the target string is not found. Otherwise, it initializes two pointers,
// left and right, to the start and end of the array respectively. It then enters a loop where it calculates
// the middle index and compares the target string with the string at that index. If they match, it returns
// the index. If the string at the middle index is greater than the target string, it means the target string
// is in the left half of the array, so it updates the right pointer. Otherwise, it updates the left pointer.
// The loop continues until the left and right pointers meet, indicating that the target string is not found
// in the array. The function then returns -1.
//
//	like go standard library : sort.SearchStrings(a []string,x string) int (https://pkg.go.dev/sort#SearchStrings)
func SearchString(dir []string, target string) int {
	if len(dir) == 0 {
		return -1
	}

	left := 0
	right := len(dir) - 1
	for left < right {
		mid := left + (right-left)/2
		if dir[mid] == target {
			return mid
		}
		if dir[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
