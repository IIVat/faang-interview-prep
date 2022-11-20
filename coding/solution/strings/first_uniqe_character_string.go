package main

import "fmt"

func firstUniqChar(s string) int {
	//here can be map as well but somehow the solution a bit faster
	m := make([]int, 26)

	for _, v := range s {
		m[v-'a']++
	}

	for i, v := range s {
		if m[v-'a'] == 1 {
			return i
		}
	}

	return -1
}

func firstUniqCharNotOpt(s string) int {
	for i := 0; i < len(s); i++ {

		for j := 0; j < len(s); j++ {
			if s[i] == s[j] && i != j {

				break

			}
			if j == len(s)-1 {
				return i
			}
		}

	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("aacll"))

	fmt.Println("hello"[1:5])
}
