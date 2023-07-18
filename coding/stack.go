package coding

import (
	"math"
	"strconv"
)

// ---------- 后缀表达式 ----------
// 后缀表达式，也称为逆波兰表达式，是一种数学表达式的表示方法，其中运算符在操作数的后面。
// 后缀表达式的优点是不需要括号来表示运算的优先级，计算过程也更加直观和简单。
// 例如：中缀表达式 （1+2)*3 的后缀表达式为 12+3*
// - 初始化一个空栈。
// - 从左到右扫描后缀表达式的每个元素。
// - 如果遇到操作数（数字），将其入栈。
// - 如果遇到运算符，从栈中弹出两个操作数，进行相应的运算，并将结果入栈。
// - 重复上述步骤，直到扫描完整个后缀表达式。
// - 栈中最后剩下的元素即为最终的计算结果。
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, calculate(num1, num2, token))
		default:
			i, _ := strconv.Atoi(token)
			stack = append(stack, i)
		}
	}
	return stack[len(stack)-1]
}

func calculate(num1, num2 int, operator string) int {
	var res int
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	}
	return res
}

// ---------- 小行星碰撞 ----------
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)
	for _, asteroid := range asteroids {
		for len(stack) != 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -asteroid {
			// 栈不为空，栈顶元素为正（向右飞），栈顶元素小于取反后的待入栈元素，说明会相撞且栈顶元素会消失
			// 继续判断下一个栈顶元素是否会消失
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 && asteroid < 0 && stack[len(stack)-1] == -asteroid {
			// 栈不为空，待入栈元素为负（向左飞），栈顶元素和待入栈元素的绝对值相等，说明两个元素都会消失
			stack = stack[:len(stack)-1]
		} else if asteroid > 0 || len(stack) == 0 || stack[len(stack)-1] < 0 {
			stack = append(stack, asteroid)
		}
	}
	return stack
}

// ---------- 每日温度 ----------
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	var stack []int
	for i, t := range temperatures {
		for len(stack) != 0 && t > temperatures[stack[len(stack)-1]] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[prev] = i - prev
		}
		stack = append(stack, i)
	}
	return res
}

// ---------- 直方图最大矩形面积 ----------
// 排列组合，两重循环，固定第一个柱子之后，把所有的面积都计算出来。
// 记录当前最小的高度和最大的面积，最后得到结果。
func largestRectangleAreaV1(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		min := heights[i]
		for j := i; j < len(heights); j++ {
			min = int(math.Min(float64(min), float64(heights[j])))
			area := min * (j - i + 1)
			maxArea = int(math.Max(float64(area), float64(maxArea)))
		}
	}
	return maxArea
}

// 先找到直方图中高度最低的位置，然后计算面积有三种情况：
// 1. 面积最大的矩形，穿过高度最低的柱子
// 2. 面积最大的矩形，在高度最低的柱子左边
// 3. 面积最大的矩形，在高度最低的柱子右边
// 用分治的思想，递归的方式，将求解过程不断拆分。
// 时间复杂度：O(nlogn)，空间复杂度：O(logn)
func largestRectangleAreaV2(heights []int) int {
	return helper(heights, 0, len(heights))
}

func helper(heights []int, start, end int) int {
	if start == end {
		return 0
	}
	if start+1 == end {
		return heights[start]
	}

	// 寻找当前部分中，高度最低的柱子
	minIndex := start
	for i := start + 1; i < end; i++ {
		if heights[i] < heights[minIndex] {
			minIndex = i
		}
	}

	area := heights[minIndex] * (end - start)
	left := helper(heights, start, minIndex)
	right := helper(heights, minIndex+1, end)
	res := math.Max(math.Max(float64(area), float64(left)), float64(right))
	return int(res)
}

// 使用一个栈，来保存直方图中柱子的高度，保证柱子是按照高度递增的，栈中保存柱子的下标。
// 1. 从左到右逐一扫描数组中的每根柱子。
// 2. 如果当前柱子的高度大于位于栈顶的柱子的高度，那么将该柱子的下标入栈；
// 3. 否则，将位于栈顶的柱子的下标出栈，并且计算以位于栈顶的柱子为顶的最大矩形面积。
//
// 以某根柱子为顶的最大矩形，一定是从该柱子向两侧延伸直到遇到比它矮的柱子，
// 这个最大矩形的高是该柱子的高，最大矩形的宽是两侧比它矮的柱子中间的间隔。
// 时间复杂度：O(n)，空间复杂度：O(n)
func largestRectangleAreaV3(heights []int) int {
	stack := make([]int, 0)
	stack = append(stack, -1)

	var maxArea int
	for i := 0; i < len(heights); i++ {
		for stack[len(stack)-1] != -1 && heights[stack[len(stack)-1]] >= heights[i] {
			height := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width := i - stack[len(stack)-1] - 1
			maxArea = int(math.Max(float64(maxArea), float64(height*width)))
		}
		stack = append(stack, i)
	}

	for stack[len(stack)-1] != -1 {
		height := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		width := len(heights) - stack[len(stack)-1] - 1
		maxArea = int(math.Max(float64(maxArea), float64(height*width)))
	}
	return maxArea
}

// ---------- 矩阵中的最大矩形 ----------
// 如果分别以图矩阵的每行为基线，就可以得到 4 个由数字 1 的格子组成的直方图。
// 在将矩阵转换成多个直方图之后，就可以计算并比较每个直方图的最大矩形面积，
// 所有直方图中的最大矩形也是整个矩阵中的最大矩形。
func maximalRectangle(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for _, row := range matrix {
		for i := 0; i < len(row); i++ {
			if row[i] == 0 {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}
		maxArea = int(math.Max(float64(maxArea), float64(largestRectangleAreaV2(heights))))
	}
	return maxArea
}
