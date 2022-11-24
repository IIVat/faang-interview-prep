package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	var stack = []string{}
	var curr = ""
	var num = 0

	for _, v := range s {
		if v == '[' {
			stack = append(stack, curr)
			stack = append(stack, fmt.Sprint((num)))
			num = 0
			curr = ""
		}

		if unicode.IsLetter(v) {
			curr += string(v)
		}

		if unicode.IsDigit(v) {
			num = num*10 + int(v-'0')
		}

		if v == ']' {
			len := len(stack)
			n, _ := strconv.Atoi(stack[len-1])

			str := stack[len-2]
			stack = stack[:len-2]
			curr = str + strings.Repeat(curr, n)

		}

	}

	return curr
}

func main() {
	println(decodeString("3[a]2[bc]"))
	// println(decodeString("ab"))
	// println(decodeString("3[a2[c]]"))
	println(decodeString("3[2[3[c]]]"))
	// println(decodeString("100[leetcode]"))
	// println(strings.Repeat("a", 3))
}
