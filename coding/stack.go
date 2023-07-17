package coding

import "strconv"

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
