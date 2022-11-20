package main

import "fmt"

func validWordSquare(words []string) bool {
	for i := 0; i < len(words); i++ {

		for j := 0; j < len(words[i]); j++ {
			if j >= len(words) || i >= len(words[j]) {
				return false
			}
			if words[i][j] != words[j][i] {
				return false
			}
		}

	}

	return true
}

func main() {
	fmt.Println(validWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))
	fmt.Println(validWordSquare([]string{"abcd",
		"bnrt",
		"crmy",
		"dtye"}))
	fmt.Println(validWordSquare([]string{"ball",
		"area",
		"read",
		"lady"}))

	fmt.Println(validWordSquare([]string{"abcd", "bnrt", "crm", "dt"}))

}
