package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if unicode.IsDigit(rune(abbr[j])) {
			// get the number, add number to i
			if abbr[j] == '0' {
				return false
			}
			num := 0
			for j < len(abbr) && unicode.IsDigit(rune(abbr[j])) {
				num = num*10 + int(abbr[j]-'0')
				j++
			}
			if num == 0 {
				return false
			}
			i += num
		} else {
			if word[i] != abbr[j] {
				return false
			}
			i++
			j++
		}
	}
	return i == len(word) && j == len(abbr)
}

func ValidWordAbbreviation(word string, abbr string) bool {
	abbrArr := ""
	currNumber := ""
	for _, ch := range abbr {
		if unicode.IsDigit(ch) {
			if ch == '0' {
				return false
			}
			currNumber += string(ch)
		} else {
			if currNumber != "" {
				n, _ := strconv.Atoi(currNumber)
				for i := 0; i < n; i++ {
					abbrArr += "*"
				}
				abbrArr += string(ch)
				currNumber = ""
			} else {
				abbrArr += string(ch)
			}
		}
	}

	if currNumber != "" {
		n, _ := strconv.Atoi(currNumber)
		for i := 0; i < n; i++ {
			abbrArr += "*"
		}
	}

	for idx, v := range word {
		if v != rune(abbrArr[idx]) && abbrArr[idx] != '*' {

			return false
		}
	}

	return true
}

func isDigit(ch byte) bool {
	return ch >= 30 && ch <= 39
}

func main() {
	fmt.Println(validWordAbbreviation("internationalization", "i12iz4n"))
	fmt.Println(validWordAbbreviation("apple", "a1pl1"))
}
