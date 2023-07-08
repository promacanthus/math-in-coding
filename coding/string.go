package coding

import (
	"math"
)

// ---------- 字符串中的变位词 ----------
// 变位词：组成各个单词的字母及每个字母出现的次数完全相同，只是字母排列的顺序不同。
// 1. 变位词长度相同
// 2. 组成变位词的字母集合相同
// 3. 每个字母出现的次数相同
//
// 每次移动这两个指针时，相当于在原来的子字符串的最右边添加一个新的字符，并且从原来子字符串中删除最左边的字符。
// - 每当在子字符串中添加一个字符时，就把哈希表中对应位置的值减 1。
// - 每当在子字符串中删除一个字符时，就把哈希表中对应位置的值加 1。
func checkInclusion(str1, str2 string) bool {
	if len(str2) < len(str1) {
		return false
	}

	count := make([]int, 26)
	for i := 0; i < len(str1); i++ {
		// 记录 str1 的字母的同时把 str2 的第一个子串先检查掉
		// 后面每次只要移动整个子串，中间所有的字母就都比较过了
		count[str1[i]-'a']++
		count[str2[i]-'a']-- // 遍历第一个子串
	}

	// i 表示子串的最右边
	for i := len(str1); i < len(str2); i++ {
		// i 增大表示子串向右一定
		count[str2[i]-'a']--           // 对应的字母要检查，即 --
		count[str2[i-len(str1)]-'a']++ // 已经检查过的字母对应的位置要还原，即 ++
		if isZero(count) {
			return true
		}
	}
	return false
}

func isZero(count []int) bool {
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

// ---------- 不含重复字符的最长子字符串 ----------
// 使用一个哈希表统计子字符串中字符出现的次数。
func longestSubstringV1(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := make([]int, 256) // 假设只包含 ASCII 码表中的字符
	l, r, length := -1, 0, 1
	for ; r < len(s); r++ {
		count[s[r]-'a']++
		for hasDuplicates(count) {
			l++
			count[s[l]-'a']--
		}
		length = int(math.Max(float64(length), float64(r-l)))
	}
	return length
}

func hasDuplicates(count []int) bool {
	for _, v := range count {
		if v > 1 {
			return true
		}
	}
	return false
}

func longestSubstringV2(s string) int {
	if len(s) == 0 {
		return 0
	}

	count := make([]int, 256)
	l, r, length, countDup := -1, 0, 1, 0
	for ; r < len(s); r++ {
		count[s[r]-'a']++
		if count[s[r]-'a'] == 2 {
			countDup++
		}
		for countDup > 0 {
			l++
			count[s[l]-'a']--
			if count[s[l]-'a'] == 1 {
				countDup--
			}
		}
		length = int(math.Max(float64(length), float64(r-l)))
	}
	return length
}

// ---------- 包含所有字符的最短字符串 ----------
// 用一个哈希表来统计字符串 t 中每个字符出现的次数。
// 1. 扫描字符串 t，每扫描到一个字符，就把该字符在哈希表中对应的值加 1。
// 2. 扫描字符串s，每扫描一个字符，就检查哈希表中是否包含该字符。
//   - 如果哈希表中没有该字符，则说明该字符不是字符串 t 中的字符，可以忽略不计
//   - 如果哈希表中存在该字符，则把该字符在哈希表中的对应值减 1。
//
// 3. 如果字符串 s 中包含字符串 t 的所有字符，那么哈希表中最终所有的值都应该小于或等于 0。
func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	set := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		set[t[i]]++
	}

	l, r, minL, minR := 0, 0, 0, 0
	length := math.MaxInt
	for r < len(s) {
		if !contains(set) {
			if _, ok := set[s[r]]; ok {
				// s 的子串中有 t 的字母
				set[s[r]]--
			}
			r++
		} else {
			if r-l < length {
				length = r - l
				minL = l
				minR = r
			}
			if _, ok := set[s[l]]; ok {
				set[s[l]]++
			}
			l++
		}
	}

	if length < math.MaxInt {
		return s[minL:minR]
	}
	return ""
}

func contains(set map[byte]int) bool {
	sum := 0
	for _, v := range set {
		sum += v
	}
	return sum <= 0
}

// ---------- 有效的回文 ----------
// 只考虑字母和数字，并忽略大小写
// 双指针从两端开始比较
func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isLetterOrDigit(s[i]) {
			i++
		} else if !isLetterOrDigit(s[j]) {
			j--
		} else {
			if toLower(s[i]) != toLower(s[j]) {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func isLetterOrDigit(b byte) bool {
	return b >= 'a' && b <= 'z' || // 是小写字母
		b >= 'A' && b <= 'Z' || // 是大写字母
		b >= '0' && b <= '9' // 是数字
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 'a' - 'A' // 大写转小写
	}
	return b
}

// ---------- 最多删除一个字符得到回文 ----------
// 双指针，从两端开始比较，遇到不相等是，删除一个字符再判断
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < len(s)/2 {
		if s[i] != s[j] {
			break
		}
		i++
		j--
	}
	return i == len(s)/2 || // 字符串本身就是回文
		isPalindromeForIndex(s, i+1, j) || // 删除不相等字符中，左边的那个，判断剩下字符是否回文
		isPalindromeForIndex(s, i, j-1) // 删除不相等字符中，右边的那个，判断剩下字符是否回文
}

func isPalindromeForIndex(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// ---------- 回文子字符串的个数 ----------
// 双指针从字符串的中心开始向两端延伸
func countSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(s); i++ {
		count += countPalindrome(s, i, i)   // 长度为奇数的回文子字符串的对称中心
		count += countPalindrome(s, i, i+1) // 长度为偶数的回文子字符串的对称中心
	}
	return count
}

func countPalindrome(s string, start, end int) int {
	count := 0
	for start >= 0 && end < len(s) && s[start] == s[end] {
		count++
		start--
		end++
	}
	return count
}
