package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	maxLength := max(len(num1), len(num2))

	// allocate an extra byte for carry, we will remove it if carry is 0
	sum := make([]byte, maxLength+1)

	carry := byte(0)

	for i := 0; i < maxLength; i++ {
		a, b := byte(0), byte(0)

		if i < len(num1) {
			a = atoi(num1[len(num1)-i-1])
		}
		if i < len(num2) {
			b = atoi(num2[len(num2)-i-1])
		}

		s := a + b + carry
		carry = s / 10
		sum[len(sum)-i-1] = itoa(s % 10)
	}

	if carry > 0 {
		sum[0] = '1'
	} else {
		sum = sum[1:]
	}

	return string(sum)
}

// atoi is usually used in strongly typed languages and opens as ASCII to Integer
// because in ASCII characters are represented by number in the ASCII table we can convert
// the digit character into it's number by substracting the number of the first digit in the table '0'
func atoi(a byte) byte {
	fmt.Println(a - '0')
	return a - '0'
}

// Integer to ASCII
// We're performing the reverse of atoi by adding the char code of '0'
func itoa(a byte) byte {
	return '0' + a
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(addStrings("120", "30"))
}

//my solution

func addStrings(num1 string, num2 string) string {
	res := ""
	tmp := 0
	maxLength := max(len(num1), len(num2))

	for i := 0; i < maxLength; i++ {

		var a int
		if i < len(num1) {
			a, _ = strconv.Atoi(string(num1[len(num1)-i-1]))
		}
		var b int

		if i < len(num2) {
			b, _ = strconv.Atoi(string(num2[len(num2)-i-1]))
		}

		sum := a + b + tmp
		tmp = sum / 10

		res = fmt.Sprint(sum%10) + res
	}

	if tmp == 0 {
		return res
	}

	return fmt.Sprint(tmp) + res
}
