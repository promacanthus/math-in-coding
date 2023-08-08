package coding

import (
	"sort"
	"strconv"
)

// ---------- 所有子集 ----------
func subsets(nums []int) [][]int {
	var res [][]int
	if len(nums) == 0 {
		return res
	}
	subsetHelper(nums, 0, make([]int, 0), &res)
	return res
}

func subsetHelper(nums []int, idx int, subset []int, res *[][]int) {
	if idx == len(nums) {
		// if index is equal to the number of elements,
		// all elements has been picked,
		// the subset is the result to append to the result.
		*res = append(*res, subset)
	} else if idx < len(nums) {
		subsetHelper(nums, idx+1, subset, res)                    // don't select the element
		subsetHelper(nums, idx+1, append(subset, nums[idx]), res) // select the element
	}
}

// ---------- 包含K个元素的组合 ----------
func combine(n, k int) [][]int {
	var (
		sub []int
		res [][]int
	)
	combineHelper(n, k, 1, sub, &res)
	return res
}

func combineHelper(n, k, idx int, sub []int, res *[][]int) {
	if len(sub) == k {
		*res = append(*res, sub)
	} else if idx <= n {
		combineHelper(n, k, idx+1, sub, res)
		combineHelper(n, k, idx+1, append(sub, idx), res)
	}
}

// ---------- 允许重复选择元素的组合 -----------
func combinationSum(nums []int, sum int) [][]int {
	var (
		res [][]int
		tmp []int
	)
	csHelper(nums, sum, 0, tmp, &res)
	return res
}

func csHelper(nums []int, target, idx int, tmpRes []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, tmpRes)
	} else if target > 0 && idx < len(nums) {
		csHelper(nums, target, idx+1, tmpRes, res)
		csHelper(nums, target-nums[idx], idx, append(tmpRes, nums[idx]), res)
	}
}

// ---------- 包含重复元素集合的组合 -----------
func combinationSumWithRedundant(nums []int, sum int) [][]int {
	sort.Ints(nums)
	var (
		res    [][]int
		tmpRes []int
	)
	cswrHelper(nums, sum, 0, tmpRes, &res)
	return res
}

func cswrHelper(nums []int, sum int, idx int, tmpRes []int, res *[][]int) {
	if sum == 0 {
		*res = append(*res, tmpRes)
	} else if sum > 0 && idx < len(nums) {
		cswrHelper(nums, sum, getNextIndex(nums, idx), tmpRes, res) // skip the same value
		cswrHelper(nums, sum-nums[idx], idx+1, append(tmpRes, nums[idx]), res)
	}
}

func getNextIndex(nums []int, idx int) int {
	next := idx
	for next < len(nums) && nums[next] == nums[idx] {
		next++
	}
	return next
}

// ---------- 没有重复元素集合的全排列 ----------
func permute(nums []int) [][]int {
	var res [][]int
	permuteHelper(nums, 0, &res)
	return res
}

func permuteHelper(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		// Found a permutation, add it to the result
		*res = append(*res, append([]int{}, nums...))
	} else {
		for i := start; i < len(nums); i++ {
			// Swap current element with the element at index 'start'
			nums[start], nums[i] = nums[i], nums[start]
			// Recursively generate permutations for the remaining elements
			permuteHelper(nums, start+1, res)
			// Undo the swap (backtrack)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}

// ---------- 包含重复元素集合的全排列 ----------
func permuteUnique(nums []int) [][]int {
	var res [][]int
	permuteUniqueHelper(nums, 0, &res)
	return res
}

func permuteUniqueHelper(nums []int, start int, res *[][]int) {
	if start == len(nums) {
		*res = append(*res, append([]int{}, nums...))
	} else {
		set := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if !set[nums[i]] {
				set[nums[i]] = true
				nums[start], nums[i] = nums[i], nums[start]
				permuteUniqueHelper(nums, start+1, res)
				nums[start], nums[i] = nums[i], nums[start]
			}
		}
	}
}

// ---------- 生成匹配的括号 ----------
func generateParenthesis(n int) []string {
	var res []string
	parenthesisHelper(n, n, "", &res)
	return res
}

// - left represents the numbers of `(` to be generated
// - right represents the number of `)` to be generated
func parenthesisHelper(left, right int, tmp string, res *[]string) {
	if left == 0 && right == 0 {
		// all left and right are generated.
		*res = append(*res, tmp)
	}
	if left > 0 {
		parenthesisHelper(left-1, right, tmp+"(", res)
	}
	if left < right { // has generated left is greater than right
		parenthesisHelper(left, right-1, tmp+")", res)
	}
}

// ---------- 分割回文子字符串 -----------
func partitionString(s string) [][]string {
	var res [][]string
	partitionHelper(s, nil, &res)
	return res
}

func partitionHelper(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		// Found a valid partition, add it to the result
		*res = append(*res, append([]string(nil), path...))
		return
	}
	for i := 1; i <= len(s); i++ {
		subStr := s[:i]
		if isPalindromeString(subStr) {
			// Include the current substring in the partition
			partitionHelper(s[i:], append(path, subStr), res)
		}
	}
}

func isPalindromeString(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// ---------- 恢复IP地址 ----------
func restoreIPAddress(str string) []string {
	var res []string
	restoreHelper(str, 0, 0, "", "", &res)
	return res
}

func restoreHelper(str string, idx, segIdx int, seg, ip string, res *[]string) {
	if idx == len(str) && segIdx == 3 && isValidIPSegment(seg) {
		*res = append(*res, ip+seg)
		return
	}

	if idx < len(str) && segIdx <= 3 {
		if isValidIPSegment(seg + string(str[idx])) {
			// current segment append current character is a valid IP segment
			restoreHelper(str, idx+1, segIdx, seg+string(str[idx]), ip, res)
		}
		if len(seg) > 0 && // current segment must not be empty
			segIdx < 3 { // all segments must less than 3
			restoreHelper(str, idx+1, segIdx+1, string(str[idx]), ip+seg+".", res)
		}
	}
}

// isValidIPSegment checks if the given IP segment is valid.
// If it is not start with zero and the value is less than or equal to 255.
func isValidIPSegment(seg string) bool {
	i, err := strconv.Atoi(seg)
	if err != nil {
		return false
	}
	return i <= 255 && seg[0] != '0'
}
