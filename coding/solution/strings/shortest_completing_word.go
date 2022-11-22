package main

import (
	"fmt"
)

func clearString(str string) [26]int {
	acc := [26]int{}
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		if char >= 'a' && char <= 'z' {
			acc[char-'a']++
		}

	}
	return acc

}

func shortestCompletingWord(licensePlate string, words []string) string {
	alphas := clearString(licensePlate)

	tmp := alphas

	shortest := -1

	for idx, word := range words {

		for _, ch := range word {
			if tmp[ch-'a'] > 0 {
				tmp[ch-'a']--
			}

		}

		isValidWord := true
		for _, v := range tmp {
			if v != 0 {
				isValidWord = false
				break
			}
		}

		if isValidWord {
			if shortest == -1 || len(word) < len(words[shortest]) {
				shortest = idx
			}
		}

		tmp = alphas
	}

	if shortest > -1 {
		return words[shortest]
	}

	return ""
}

func main() {

	// lpCleaned := strings.ToLower(clearString("1s3 PSt"))
	// alphas := [26]int{}

	// for _, v := range lpCleaned {
	// 	alphas[v-'a']++
	// }

	// tmp := alphas

	// tmp[0]--

	// fmt.Println(alphas)
	// fmt.Println(tmp)

	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	fmt.Println(shortestCompletingWord("1s3 456", []string{"look", "pets", "stew", "show"}))

}
