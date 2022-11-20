package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	bytes := []byte(s)

	i, j := 0, len(s)-1
	for i < j {
		if isVowel(bytes[i]) && isVowel(bytes[j]) {
			bytes[i], bytes[j] = bytes[j], bytes[i]
			i++
			j--
		} else if !isVowel(bytes[i]) {
			i++
		} else if !isVowel(bytes[j]) {
			j--
		}
	}
	return string(bytes)

}

func isVowel(s byte) bool {
	switch s {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(reverseVowels("aA"))
}
