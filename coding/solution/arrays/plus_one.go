package main

import "fmt"

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if (digits[i]+1)/10 == 0 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}

	return digits
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{1}))
	fmt.Println(plusOne([]int{1, 2, 3}))
}
