package coding

import (
	"math"
	"strconv"
)

// ---------- 整数除法 -----------
// 注意点：
// 1. 数据溢出问题，-2^31 / -1
// 2. 符号问题，负数转整数也会溢出，都转为负数一起处理
// 3. 被除数比除数大太多的性能问题，直接用除数的倍数去比较
func divide(dividend, divisor int) int {
	// int 的范围是 -2^31 ~ (2^31)-1, -2^31 / -1 会溢出
	if dividend == math.MinInt && divisor == -1 {
		return math.MaxInt
	}

	// 最小的整数是 -2^31，最大的整数是 (2^31)-1
	// 如果将 -2^31 转换为正数则会导致溢出，将任意正数转换为负数都不会溢出，
	// 先将正数都转换成负数，用减法计算两个负数的除法，然后根据需要调整商的正负号。
	negative := 2
	if dividend > 0 {
		negative--
		dividend = -dividend
	}
	if divisor > 0 {
		negative--
		divisor = -divisor
	}

	result := divideCore(dividend, divisor)
	if negative == 1 {
		return -result
	}
	return result
}

func divideCore(dividend, divisor int) int {
	result := 0
	for dividend <= divisor {
		value := divisor
		quotient := 1
		// 当被除数远大于除数时，性能优化:
		// - 不断计算 value 的倍数，即除数的倍数
		// - 比较被除数和 value 的大小，直到被除数小于 value 的值
		// - 被除数直接减去 value，即，减去那么多倍的除数
		for value >= math.MinInt/2 && dividend <= value+value {
			quotient += quotient
			value += value
		}
		result += quotient
		dividend -= value
	}
	return result
}

// ---------- 二进制加法 ----------
// 如果将字符串转为十进制数字相加后再转为二进制字符串可能会溢出。
func addBinary(a, b string) string {
	var (
		result []byte
		res    string
		carry  byte
	)

	i := len(a) - 1
	j := len(b) - 1

	// 二进制最低位在最右侧
	for i >= 0 || j >= 0 {
		digitA := a[i] - '0'
		i--
		digitB := b[j] - '0'
		j--
		sum := digitA + digitB + carry
		if sum >= 2 {
			sum -= 2
			carry = 1
		} else {
			carry = 0
		}
		result = append(result, sum)
	}

	if carry == 1 {
		result = append(result, 1)
	}

	// 反转数组
	for i, j := 0, len(result)-1; i < len(result)/2; i++ {
		result[i], result[j] = result[j], result[i]
	}

	// 将二进制的数字转为对应的字符串，直接用`string()`相当于输出 ASCII 表中对应位的符号
	for _, b := range result {
		res += strconv.FormatUint(uint64(b), 2)
	}
	return res
}

// ---------- 前 n 个数字，二进制形式中 1 的个数 ----------
// 计算整数 i 的二进制形式中 1 的个数，比较高效的一种方式是 `i&(i-1)`。
func countBitsV1(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		j := i
		for j != 0 {
			result[i]++
			j = j & (j - 1)
		}
	}
	return result
}

// 整数 i 的二进制形式中 1 的个数比 `i&(i-1)` 多一个 1
func countBitsV2(n int) []int {
	result := make([]int, n+1)
	for i := 1; i <= n; i++ {
		result[i] = result[i&(i-1)] + 1
	}
	return result
}

// 对于正整数 i，
// - i 是偶数，相当于 `i/2` 左移一位的结果，因此 1 的个数相同
// - i 是奇数，相当于 `i/2` 左移一位后加 1 的结果，因此 i 中 1 的个数多一个
//
// 例如，3 的二进制形式是 11
// - 6 是偶数，二进制形式为 110, 3 << 1
// - 7 是奇数，二进制形式为 111, (3 << 1) + 1
//
// 二进制中，最低一位表示奇偶性,
// - i 是偶数，最低一位为 0，`i&1` = 0
// - i 是奇数，最低一位为 1，`i&1` = 1
func countBitsV3(n int) []int {
	result := make([]int, n+1)
	for i := 0; i <= n; i++ {
		result[i] = result[i>>1] + (i & 1)
	}
	return result
}

// ---------- 只出现一次的数字 ----------
// 任何一个数字异或它自己的结果都是 0。
// 将数组中所有数字的同一位置的数位相加。这些出现了 3 次的数字的任意第 i 个数位之和都能被 3 整除。
//
// - 如果数组中所有数字的第 i 个数位相加之和能被 3 整除，那么只出现一次的数字的第 i 个数位一定是 0
// - 如果数组中所有数字的第 i 个数位相加之和被 3 除余1，那么只出现一次的数字的第 i 个数位一定是 1
//
// 这里的出现 3 次和出现 1 次，可以推广到出现 n 和 m 次，m 不能被 n 整除。
func singleNumber(numbers []int) int {
	bitSum := make([]int, 32)
	for _, num := range numbers {
		for i := 0; i < 32; i++ {
			// 数组 0 存字符最高位
			// (num >> (31 - i)) & 1 获取 num 从左边起第 i 位
			bitSum[i] += (num >> (31 - i)) & 1
		}
	}
	result := 0
	for i := 0; i < 32; i++ {
		// bitSum[0] 存放的是最左边的最高位，通过 for 循环，
		// 不断将 result 右移一位，将每一位移动到对应的位置
		result = (result << 1) + (bitSum[i] % 3)
	}
	return result
}

// ---------- 单词长度的最大乘积 ----------
func maxProductV1(words []string) int {
	// 使用一个 2D 数组保存每一个单词和出现的字母
	// 26 个小写英文字母分别对应数组下标 0-25
	flags := make([][]bool, len(words))
	for i := 0; i < len(words); i++ {
		flags[i] = make([]bool, 26)
		for _, c := range words[i] {
			flags[i][c-'a'] = true
		}
	}

	var result float64
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			k := 0
			for ; k < 26; k++ {
				if flags[i][k] && flags[j][k] {
					// 说明两个单词有重复的字母
					break
				}
			}
			if k == 26 {
				// 两个单词没有重复的字母，计算当前单词长度的乘积
				prod := len(words[i]) * len(words[j])
				result = math.Max(float64(result), float64(prod))
			}
		}
	}
	return int(result)
}

// 只用位运算判断两个字符是否含有相同的字母更快。
func maxProductV2(words []string) int {
	// 用一个 int 的每一位来表示字符串是否有改字母
	// 26 个小写英文字母分别对应 int 的下标 0-25
	flags := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		for _, c := range words[i] {
			// 用或操作将某一位置为 1
			flags[i] |= 1 << (c - 'a')
		}
	}
	var result float64
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			// 如果两个字符串中包含相同的字符，那么它们对应的整数相同的某个数位都为 1，两个整数的与运算将不会等于 0。
			// 如果两个字符串没有相同的字符，那么它们对应的整数的与运算的结果等于 0。
			if (flags[i] & flags[j]) == 0 {
				prod := len(words[i]) * len(words[j])
				result = math.Max(result, float64(prod))
			}
		}
	}
	return int(result)
}
