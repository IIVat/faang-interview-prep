package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	parentheses := make(map[rune]rune)

	parentheses[']'] = '['
	parentheses[')'] = '('
	parentheses['}'] = '{'

	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else if parentheses[v] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}

	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("[]{}()"))
	fmt.Println(isValid("[{}()"))
	fmt.Println(isValid("[{}]([])"))
	// <[> =>  <[{> => [ =>
}
