package main

import "math"

func happyNumber(num int) bool {
	if num >= math.MaxInt {
		return false
	}

	slowPointer := num
	fastPointer := sumOfSquaresDigits(num)

	for {

		if fastPointer == 1 {
			return true
		}
		if fastPointer == slowPointer {
			return false
		}

		println(fastPointer)

		slowPointer = sumOfSquaresDigits(slowPointer)
		fastPointer = sumOfSquaresDigits(sumOfSquaresDigits(fastPointer))
	}

}

func sumOfSquaresDigits(n int) int {
	d := n
	sum := 0

	for d != 0 {
		tmp := d % 10
		d = d / 10
		sum += (tmp * tmp)
	}
	return sum
}

func main() {
	println(happyNumber(1))
}
