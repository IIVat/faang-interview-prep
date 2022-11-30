package main

func expressiveWords(s string, words []string) int {
	result := 0

	for _, v := range words {

		if isExtensibleWord(v, s) {
			result++
		}
	}

	return result
}

func isExtensibleWord(queryWord string, searchWord string) bool {
	lenQW := len(queryWord)
	lenS := len(searchWord)
	i := 0
	j := 0

	for i < lenQW && j < lenS {

		if queryWord[i] != searchWord[j] {
			return false
		}

		countS := getRepeatedLength(searchWord, j)
		countQW := getRepeatedLength(queryWord, i)

		if (countQW != countS && countS < 3) || countQW > countS {
			return false
		}

		j += countS
		i += countQW
	}

	return i == lenQW && j == lenS
}

func getRepeatedLength(str string, i int) int {
	j := i
	for j < len(str) && str[j] == str[i] {
		j++
	}
	return j - i
}

func main() {
	println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
	println(expressiveWords("zzzzzyyyyy", []string{"zzyy", "zy", "zyy"}))
	println(expressiveWords("abcd", []string{"abc"}))
	// println(isExtensibleWord("hello", "heeelooo"))
	println(expressiveWords("hihelohello", []string{"hello", "hi", "helo"}))

}
