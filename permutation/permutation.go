package permutation

import "fmt"

var (
	qwHorseTime = map[string]float64{
		"q1": 1,
		"q2": 2,
		"q3": 3,
	}

	tjHorseTime = map[string]float64{
		"t1": 1.5,
		"t2": 2.5,
		"t3": 3.5,
	}

	qwHorses = []string{"q1", "q2", "q3"}
)

// permutation takes in an array of horse names and an empty result array. It recursively generates all possible permutations
// of the horse names and prints them out. It also compares each permutation with a predefined array of desired horse names
// and prints out any matches. To generate the permutations, the function iterates through each horse name in the input array
// and creates a new result array with that horse name appended. It then creates a new array of remaining horse names by
// removing the selected horse name from the input array. It is called recursively with the new result and remaining horse
// arrays. Overall, it is useful for generating and comparing all possible orderings of a set of horse names.
func permutation(horses []string, result []string) {
	if len(horses) == 0 {
		fmt.Println(result)
		compare(result, qwHorses)
		return
	}

	var (
		newResult []string
		restHorse []string
	)
	for i := 0; i < len(horses); i++ {
		newResult = make([]string, len(result))
		copy(newResult, result)
		newResult = append(newResult, horses[i])

		restHorse = make([]string, len(horses))
		copy(restHorse, horses)
		restHorse = append(restHorse[:i], restHorse[i+1:]...)
		permutation(restHorse, newResult)
	}
}

// compare takes in two slices of strings representing the horses selected by two players, TJ and QW.
// It then compares the times of each horse from the respective players' selections and keeps track
// of how many times TJ's horses had faster times than QW's horses. If TJ's horses had faster times
// more than half of the time, the function prints "TJ Win", otherwise it prints "QW Win". It is
// useful for determining the winner of a horse race between two players.
func compare(tj, qw []string) {
	count := 0
	for i := 0; i < len(tj); i++ {
		if tjHorseTime[tj[i]] < qwHorseTime[qw[i]] {
			count++
		}
	}
	if count > len(tj)/2 {
		fmt.Println("TJ Win")
	} else {
		fmt.Println("QW Win")
	}
}

// GeneratePassword takes in two string slices - input and result. It generates all possible combinations of the
// elements in input and prints them out. It does this by iterating through the input slice and appending
// each element to a copy of the result slice, then recursively calling itself with the new result slice.
// This process continues until the length of the result slice is equal to the length of the input slice.
// At this point, all possible combinations have been generated and the function prints out the final result.
func GeneratePassword(characters, generatedPassword []string) {
	if len(characters) == 0 {
		return
	}
	if len(generatedPassword) == len(characters) {
		fmt.Println(generatedPassword)
		return
	}
	var newGeneratedPassword []string
	for i := 0; i < len(characters); i++ {
		newGeneratedPassword = make([]string, len(generatedPassword)+1)
		copy(newGeneratedPassword, generatedPassword)
		newGeneratedPassword[len(generatedPassword)] = characters[i]
		GeneratePassword(characters, newGeneratedPassword)
	}
}
