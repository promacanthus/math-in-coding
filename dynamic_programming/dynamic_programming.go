package dynamicprogramming

import "math"

// getStrDistance takes two input strings, "a" and "b", and returns an integer representing the minimum number of operations required.
// If either of the input strings is empty, the function returns -1. The algorithm uses a dynamic programming approach to calculate
// the minimum distance.
//
// It creates a 2D array, "d", of size (len(a)+1) x (len(b)+1), where each cell represents the distance between substrings of "a" and "b".
// The first two for loops initialize the base cases of the dynamic programming matrix.
//
//	1. The value in cell (i,0) represents the distance between the substring of "a" from index 0 to i and an empty string,
//	which is equal to the length of the substring.
//	2. The value in cell (0,j) represents the distance between an empty string and the substring of "b" from index 0 to j.
//
// The nested for loops iterate through each character in both strings. For each character, the algorithm calculates the minimum
// distance by considering three possible operations:
//
//	1. appending the character from "b" to the substring of "a",
//	2. appending the character from "a" to the substring of "b",
//	3. replacing the character in "a" with the character in "b".
//
// The minimum distance is then stored in cell (i+1, j+1) of the matrix.
// Finally, the function returns the value in the bottom-right cell of the matrix, which represents the minimum distance between the entire
// strings "a" and "b". Overall, the code efficiently calculates the Levenshtein distance between two strings using dynamic programming.
func getStrDistance(a, b string) int {
	if a == "" || b == "" {
		return -1
	}

	d := make([][]int, len(a)+1)
	for i := range d {
		d[i] = make([]int, len(b)+1)
	}
	for i := 0; i <= len(a); i++ {
		d[i][0] = i
	}
	for j := 0; j <= len(b); j++ {
		d[0][j] = j
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			r := 0
			if a[i] != b[j] {
				r = 1
			}
			firstAppend := d[i][j] + 1
			secondAppend := d[i+1][j] + 1
			replace := d[i][j] + r
			min := math.Min(float64(firstAppend), float64(secondAppend))
			min = math.Min(min, float64(replace))
			d[i+1][j+1] = int(min)
		}
	}
	return d[len(a)][len(b)]
}
