package binary

import "math"

// NOTICE:
//  1. 二进制转换为十进制：请将每个二进制位乘以2的幂，然后将结果相加
//  2. 十进制转换为二进制：请将该数字除以2，直到商为0为止。每次除法的余数将是二进制位的值。最后，将这些值从右到左排列起来，以获得二进制数。

// BinaryToDecimal takes in a binary integer and converts it to decimal.
// It uses a loop to iterate through each digit of the binary number, starting from the least significant digit.
// The remainder of the current digit is calculated using the modulo operator.
// The binary number is then divided by 10 to remove the current digit.
// The decimal equivalent of the current digit is calculated using the formula 2^i * remainder, where i is the position of the digit.
// The decimal equivalents of all the digits are added together to get the final decimal number.
// The function returns the decimal equivalent of the input binary number.
func BinaryToDecimal(binary int64) int64 {
	decimal, i, remainder := int64(0), int64(0), int64(0)
	for binary != 0 {
		remainder = binary % 10 // 获取当前最低位
		binary /= 10            // 移除当前最低位
		decimal += remainder * int64(math.Pow(2, float64(i)))
		i++
	}
	return decimal
}

// DecimalToBinary takes an input decimal number and converts it to binary.
// It uses a loop to repeatedly divide the decimal number by 2 and keep track of the remainders.
// The remainders are then used to construct the binary number by multiplying each remainder by
// a power of 10 and adding it to the binary number. The function returns the binary number.
func DecimalToBinary(decimal int64) int64 {
	binary, i, remainder := int64(0), int64(1), int64(0)
	for decimal != 0 {
		remainder = decimal % 2 // 取当前最低位
		decimal /= 2            // 移除当前最低位
		binary += remainder * i
		i *= 10
	}
	return binary
}

// LeftShift takes two parameters - num and m, both of type int64.
// It performs a left shift operation on the num variable by m bits and returns the result.
// A left shift operation moves the bits of a number to the left by a certain number of positions,
// effectively multiplying the number by 2 to the power of the shift amount.
func LeftShift(num, m int64) int64 {
	return num << m
}

// RightShift takes two parameters: num and m, both of type int64.
// It performs a right shift operation on num by m bits and returns the result.
// A right shift operation moves the bits of a number to the right by a specified number of positions,
// effectively dividing the number by 2 to the power of the shift amount.
func RightShift(num, m int64) int64 {
	return num >> m
}

// OR takes two int64 arguments a and b, and returns the result of performing a bitwise OR operation on them.
// The bitwise OR operation compares the binary representation of each argument and returns a new number where
// each bit is set to 1 if either of the corresponding bits in the input numbers is 1.
func OR(a, b int64) int64 {
	return a | b
}

// AND takes two int64 arguments, "a" and "b", and returns their bitwise AND operation result.
// The "&" symbol is used to perform the AND operation. This function can be used to check
// if certain bits are set in both "a" and "b".
func AND(a, b int64) int64 {
	return a & b
}

// XOR takes two int64 values as input and returns their bitwise XOR (exclusive OR) operation.
// The XOR operator returns a 1 in each bit position where the corresponding bits of either
// but not both operands are 1. This function is a simple and efficient way to perform bitwise operations on integers.
func XOR(a, b int64) int64 {
	return a ^ b
}
