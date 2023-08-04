package coding

import (
	"math"
	"math/rand"
)

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// ---------- 查找插入位置 ----------
func searchInsertV1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target && nums[i-1] < target {
			return i
		}
	}
	return -1
}

func searchInsertV2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			// If mid is zero, all values are greater than target.
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return len(nums) // All values in the array are less than target.
}

// ---------- 山峰数组的顶部 ----------
func peakIndexMountainArrayV1(nums []int) int {
	max, index := math.MinInt, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return index
}

func peakIndexMountainArrayV2(nums []int) int {
	// 数组首尾的值 不可能是山峰，直接忽略
	left, right := 1, len(nums)-2
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			// 峰顶：数值比左边大，比右边大
			return mid
		}
		if nums[mid] > nums[mid-1] {
			// 左峰：数值比左边大，比右边小
			// 峰顶在右边
			left = mid + 1
		} else {
			// 右峰：数值比左边小，比右边大
			// 峰顶在左边
			right = mid - 1
		}
	}
	return -1
}

// ---------- 排序数组中只出现一次的数字 ----------
func singleNonDuplicateV1(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

// 将数组中的数字每两个分为一组。先找出位于中间的一组，确定这一组的两个数字是否相同。
// -1 如果两个数字相同，那么那个只出现一次的数字一定在它的后面，因此接着查找它的后半部分。
// -2 如果两个数字不相同，那么接着检查这一组是不是第1组两个数字不相同的分组。
//
//	-2.1 如果是第1组，那么这一组的第1个数字就是只出现一次的数字。
//	-2.2 如果不是第1组，那么第1组一定在它的前面，因此接着查找它的前半部分。
func singleNonDuplicateV2(nums []int) int {
	left, right := 0, len(nums)/2
	for left <= right {
		mid := (left + right) / 2
		i := mid * 2
		if i < len(nums)-1 &&
			nums[i] != nums[i+1] { // i 和他后面的一个不相等
			if mid == 0 ||
				nums[i-2] == nums[i-1] { // i 的前两个相等
				return nums[i] // 所以 i 就是唯一的那个数字  对应2.1
			}
			right = mid - 1 // 对应2.2
		} else {
			left = mid + 1 // 对应1
		}
	}
	return nums[len(nums)-1] // 所有数字都出现2次，除了最后一个数字
}

// ---------- 按权重生成随机数字 ----------
func pickIndexV1(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	index := make([]int, sum)
	for i := 0; i < sum; {
		for j, num := range nums {
			for k := 0; k < num; k++ {
				index[i] = j
				i++
			}
		}
	}
	return index[rand.Intn(sum)]
}

func pickIndexV2(nums []int) int {
	var sum int
	var sums []int
	for _, num := range nums {
		sum += num
		sums = append(sums, sum)
	}
	n := rand.Intn(sum)
	left, right := 0, len(nums)
	for left <= right {
		mid := (left + right) / 2
		if sums[mid] > n {
			if mid == 0 || sums[mid-1] <= n {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// ---------- 平方根 ----------
func mySqrt(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		// Use division to prevent overflows.
		if mid <= n/mid {
			if (mid + 1) > n/(mid+1) {
				return mid
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	// if n == 0, left = 1 and right = 0, return 0
	return 0
}

// ---------- 狒狒吃香蕉 ----------
func minEatingSpeed(piles []int, H int) int {
	max := math.MinInt
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	left, right := 1, max
	for left <= right {
		mid := (left + right) / 2
		hours := getHours(piles, mid)
		if hours <= H {
			if mid == 1 || // 最慢的速度就是 1
				getHours(piles, mid-1) > H { // mid 是最慢的速度了
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func getHours(piles []int, speed int) int {
	var hours int
	for _, pile := range piles {
		hours += int(math.Ceil(float64(pile) / float64(speed)))
	}
	return hours
}
