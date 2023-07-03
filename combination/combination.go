package combination

import "fmt"

// combine takes in three parameters: an integer "k", a slice of strings "teams", and another slice of strings "result".
// The purpose of this function is to generate all possible combinations of "k" elements from the "teams" slice and print them out.
// The function first checks if the length of the "result" slice is equal to "k". If so, it means that a combination has been formed,
// and the function prints out the combination and then returns. If the length of "result" is not equal to "k", the function proceeds
// to generate the combinations. Inside a for loop, the function creates a new slice called "newResult" with a length equal to the length
// of "result" plus one. It then copies the elements from the "result" slice to the "newResult" slice using the "copy" function. Finally,
// it assigns the value of the current element in the "teams" slice to the last index of the "newResult" slice. It then creates a new slice
// called "restTeams" by slicing the "teams" slice starting from the index after the current element. Finally, it calls itself recursively,
// passing in the value of "k", the "restTeams" slice, and the "newResult" slice as arguments. This recursive call generates the combinations
// for the remaining elements in the "teams" slice. Overall, it uses a recursive approach to generate all possible combinations of "k" elements
// from the "teams" slice and prints them out.
func combine(k int, teams, result []string) {
	if len(result) == k {
		fmt.Println(result)
		return
	}

	var newResult []string
	for i := 0; i < len(teams); i++ {
		newResult = make([]string, len(result)+1)
		copy(newResult, result)
		combine(k, teams[i+1:], newResult)
	}
}
