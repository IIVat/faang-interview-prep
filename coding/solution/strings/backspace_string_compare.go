package main

import "fmt"

func backspaceCompare(s string, t string) bool {
	return applyBackspace(s) == applyBackspace(t)
}

func applyBackspace(s string) string {
	res := ""
	for _, ch := range s {

		if ch == rune('#') {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = res + string(ch)
		}
	}

	return res
}

//O(m+n) - time
//O(1) -space

func backspaceCompareO1Space(s string, t string) bool {
	return backspace(s) == backspace(t)
}

func backspace(s string) string {
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == '#' {
			s = s[:i-1] + s[i+1:]
			i -= 2
		} else if s[i] == '#' {
			s = s[1:]
			i -= 1
		}
	}
	return s
}

func main() {
	fmt.Println(backspaceCompare("y#fo##f",
		"y#f#o##f"))

	fmt.Println("ab#c"[3])
}
package main

import "fmt"


func validWordSquare(words []string) bool {
	return true
}

func main()  {
	fmt.Println(validWordSquare([
		"abcd",
		"bnrt",
		"crm",
		"dt"
	  ]))
}
