package coding

import (
	"math"
	"sort"
)

// ---------- 排序数组中的两个数字之和 ----------
// 固定一个值，使用二分查找在有序数组中找到另一个值，两者之和为 k
// 时间复杂度：O(nlogn),空间复杂度：O(1)
func twoSumV1(nums []int, k int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		target := k - nums[i]
		if idx, find := sort.Find(len(nums), func(j int) int {
			return target - nums[j]
		}); find {
			res = append(res, i, idx)
			break
		}
	}
	return res
}

// 空间换时间，使用哈希表保存有序数组中的每一个值，
// 固定一个值，通过查找哈希表的方式判断数组中是否存在另一个值，两者之和为 k
// 时间复杂度：O(n),空间复杂度：O(n)
func twoSumV2(nums []int, k int) []int {
	var res []int
	set := make(map[int]int)
	for i, v := range nums {
		set[v] = i
	}
	for i := 0; i < len(nums); i++ {
		target := k - nums[i]
		if idx, ok := set[target]; ok {
			res = append(res, i, idx)
			break
		}
	}
	return res
}

// 使用双指针，从有序数组的首尾同时开始查找，比较两数和与 k 的大小,
// 来控制左右指针的移动，直到两指针相遇或找到目标值为止
// 时间复杂度：O(n),空间复杂度：O(1)
func twoSumV3(nums []int, k int) []int {
	var res []int
	l, r := 0, len(nums)-1
	for l < r {
		target := nums[l] + nums[r]
		if target == k {
			res = append(res, l, r)
			break
		} else if nums[l]+nums[r] < k {
			l++
		} else {
			r--
		}
	}
	return res
}

// ---------- 数组中和为 0 的 3 个数字 ----------
// 1. 数组要有序 O(nlogn)
// 2. 先固定一个值，剩下的问题就是有序数组中找两数和为目标值 O(n^2)
// 注意：固定第一个值的时候，可能存在多组与另外两个值相加和为 0
func threeSum(nums []int) [][]int {
	var result [][]int
	if len(nums) > 2 {
		sort.Ints(nums)
		for i := 0; i < len(nums)-2; {
			twoSum(nums, i, &result)
			tmp := nums[i]
			for i < len(nums) && nums[i] == tmp {
				// 过滤掉重复 i 值，因为已经排序，所以重复的值都连在一起
				i++
			}
		}
	}
	return result
}

func twoSum(nums []int, i int, res *[][]int) {
	j, k := i+1, len(nums)-1
	for j < k {
		if nums[i]+nums[j]+nums[k] == 0 {
			*res = append(*res, []int{nums[i], nums[j], nums[k]})
			tmp := nums[j]
			for nums[j] == tmp && j < k {
				// 过滤掉重复 j 值，因为已经排序，所以重复的值都连在一起
				j++
			}
		} else if nums[i]+nums[j]+nums[k] < 0 {
			j++
		} else {
			k--
		}
	}
}

// ---------- 和大于或等于 k 的最短子数组 ----------
// l 和 r 作为滑动窗口的两端
//
// 假设数组的长度为 n，尽管代码中有两个嵌套的循环，该解法的时间复杂度仍然是 O(n)。
// 在这两个循环中，变量 l 和 r 都是只增加不减少，因此总的执行次数是 O(n),
// - r: 从 0 增加到 n-1
// - l: 从 0 最多增加到 n-1
// 或者说因为整个子数组的长度相对于原数组长度不在一个数量级，计算时间复杂度时可以不考虑
func minSubArrayLen(nums []int, k int) int {
	length := math.MaxFloat64
	for l, r, sum := 0, 0, 0; r < len(nums); r++ {
		sum += nums[r] // 移动 r 是 sum 增加
		for l <= r && sum >= k {
			length = math.Min(length, float64(r-l+1))
			sum -= nums[l] // 移动 l 是 sum 减少
			l++
		}
	}
	return int(length)
}

// ---------- 乘积小于 k 的子数组 ----------
// 目标是求出所有数字乘积小于k的子数组的个数，一旦向右移动指针P1到某个位置时子数组的乘积小于 k，
// 就不需要再向右移动指针 l。因为只要保持指针 r 不动，向右移动指针 l 形成的所有子数组的数字乘积就一定小于 k。
// 此时两个指针之间有多少个数字，就找到了多少个数字乘积小于 k 的子数组。
func numSubArrayProductLessThanK(nums []int, k int) int {
	var count int
	for l, r, product := 0, 0, 1; r < len(nums); r++ {
		product *= nums[r]
		for l <= r && product >= k {
			product /= nums[l]
			l++
		}
		if r >= l {
			count += r - l + 1
		}
	}
	return count
}

// 使用双指针解决子数组之和的面试题有一个前提条件——数组中的所有数字都是正数。如果数组中的数字有正数、负数和零，
// 那么双指针的思路并不适用，这是因为当数组中有负数时在子数组中添加数字不一定能增加子数组之和，从子数组中删除数字也不一定能减少子数组之和。
//
// 换一种思路求子数组之和。
// 假设整个数组的长度为 n，它的某个子数组的第 1 个数字的下标是 i，最后一个数字的下标是 j。
// 为了计算子数组之和，需要先做预处理，计算从数组下标为 0 的数字开始到以每个数字为结尾的子数组之和。
// 预处理只需要从头到尾扫描一次，求出:
// - 从下标 0 开始到下标 0 结束的子数组之和 S0，
// - 从下标 0 开始到下标 1 结束的子数组之和 S1，
// - 以此类推，直到求出
// - 从下标 0 开始到最后一个数字的子数组之和 S(n-1)。
//
// 因此，从下标为 i 开始到下标为 j 结束的子数组的和就是 Sj-S(i-1)。

// ---------- 和为 k 的子数组 ----------
// 在从头到尾逐个扫描数组中的数字时求出前 i 个数字之和，并且将和保存下来。数组的前 i 个数字之和记为 x。
// 如果存在一个 j（j＜i），数组的前 j 个数字之和为 x-k，那么数组中从第 j+1 个数字开始到第 i 个数字结束的子数组之和为 k。
//
// 计算和为 k 的子数组的个数。
// 当扫描到数组的第 i 个数字并求得前 i 个数字之和是 x 时，需要知道在 i 之前存在多少个 j 并且前 j 个数字之和等于 x-k。
// 对每个 i，不但要保存前 i 个数字之和，还要保存每个和出现的次数。用一个哈希表，哈希表的键是前 i 个数字之和，值为每个和出现的次数。
func subArraySum(nums []int, k int) int {
	set := make(map[int]int)
	set[0] = 1
	sum, count := 0, 0
	for _, num := range nums {
		sum += num              // 计算前 i 个数之和
		count += set[sum-k]     // 在前 i 个数中，出现了多少次和为 sum-k
		set[sum] = set[sum] + 1 // 保存前 i 个数的和及出现的次数
	}
	return count
}

// ---------- 0 和 1 个数相同的子数组 ----------
// 曲线救国：
// 首输入数组中所有的 0 都替换成 -1，题目就变成求包含相同数目的 -1 和 1 的最长子数组的长度。
// 在一个只包含数字 1 和 -1 的数组中，如果子数组中 -1 和 1 的数目相同，那么子数组的所有数字之和就是 0，
// 因此题目就变成求数字之和为 0 的最长子数组的长度。
//
// 在扫描数组时累加已经扫描过的数字之和。如果数组中前 i 个数字之和为 m，前 j 个数字（j>i）之和也为 m，
// 那么从第 i+1 个数字到第 j 个数字的子数组的数字之和为 0，这个和为 0 的子数组的长度是 j-i。
func findMaxLength(nums []int) int {
	// set：
	// - key 是前 i 个元素的和
	// - value 是第 i 个元素的下标
	set := make(map[int]int)
	set[0] = -1
	sum, length := 0, float64(0)
	for i, num := range nums {
		if num == 0 {
			sum += -1
		} else {
			sum += 1
		}
		if _, ok := set[sum]; ok {
			length = math.Max(float64(length), float64(i-set[sum]))
		} else {
			set[sum] = i
		}
	}
	return int(length)
}

// ---------- 左右两边子数组的和相等 ----------
// 如果从数组的第 1 个数字开始扫描并逐一累加扫描到的数字，当扫描到第 i 个数字的时候，
// 就可以知道累加到第 i 个数字的和，这个和减去第 i 个数字就是累加到第 i-1 个数字的和。
// 同时，要知道数组中的所有数字之和，只需要从头到尾扫描一次数组就可以。
func pivotIndex(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	sum := 0
	for i, num := range nums {
		sum += num
		if sum-num == total-sum {
			return i
		}
	}
	return -1
}

// ---------- 二维子矩阵的数字之和 ----------
// 任何一个子矩阵的和，可以通过如下公式计算得到：
// 矩阵表示形式：（最上角|右下角）
// (r1,c1|r2,c2) = (0,0|r2,c2) - (0,0|r1-1,c2) - (0,0|r2,c1-1)  + (0,0|r1-1,c1-1)
//
// 预处理的时候，先求出从(0,0)到每个右下角左边的子矩阵的数字之和，
// 使用一个辅助矩阵来保存结果，为了简化计算，防止数组越界，行列分别多加 1，
// 使得矩阵最上面和最左边多一个空行和一列空列，相当于整个矩阵沿着右下的对角线移动了一格。
// 计算子矩阵的和，可以分为两部分的总和
// - 第一部分：(0,0|i,j)这个子矩阵的和
// - 第二部分：i+1这一行每个元素的累加
//
// 这样，最后计算结果时候的公式的最表要稍微调整一下：
// (r1,c1|r2,c2) = (0,0|r2+1,c2+1) - (0,0|r1,c2+) - (0,0|r2+1,c1)  + (0,0|r1,c1)
func numMatrix(matrix [][]int, r1, c1, r2, c2 int) int {
	rows := len(matrix)
	cols := len((matrix[0]))
	if rows == 0 || cols == 0 {
		return 0
	}

	// create a new matrix
	matrixSum := make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		matrixSum[i] = make([]int, cols+1)
	}
	// initialize the matrix
	for i := 0; i < rows; i++ {
		rowSum := 0
		for j := 0; j < cols; j++ {
			rowSum += matrix[i][j]
			matrixSum[i+1][j+1] = matrixSum[i][j+1] + rowSum
		}
	}
	return matrixSum[r2+1][c2+1] + matrixSum[r1][c1] - matrixSum[r1][c2+1] - matrixSum[r2+1][c1]
}
