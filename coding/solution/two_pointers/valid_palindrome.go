package main

import (
	"regexp"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func isPalindrome(inputString string) bool {
	chars := make([]uint8, 0)

	for i := 0; i < len(s); i++ {
		if inputString[i] >= 65 && inputString[i] <= 90 {
			chars = append(chars, inputString[i])
		}
		if inputString[i] >= 97 && inputString[i] <= 122 {
			chars = append(chars, inputString[i]-32)
		}
		if inputString[i] >= 48 && inputString[i] <= 57 {
			chars = append(chars, inputString[i])
		}
	}

	i, j := 0, len(inputString)-1
	for i < j {

		if inputString[i] != inputString[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {

	println(isPalindrome("A man, a plan, a canal: Panama"))

}
