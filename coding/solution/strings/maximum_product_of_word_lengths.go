package main

import (
	"sort"
)

func maxProduct(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	length, maxProd := len(words), 0
	mask := make([]int, length)

	for i := 0; i < length; i++ {
		for _, ch := range words[i] {
			mask[i] |= 1 << (ch - 'a')
		}
		for j := 0; j < i; j++ {
			if (mask[i] & mask[j]) == 0 {
				maxProd = max(maxProd, len(words[i])*len(words[j]))

			}
		}

	}

	return maxProd

}

func maxProductSimple(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	maxProd := 0

	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			isSharing := false
			for k := 0; k < len(words[i]); k++ {
				for m := 0; m < len(words[j]); m++ {
					if words[i][k] == words[j][m] {
						isSharing = true
						break
					}
				}
			}
			if !isSharing {
				maxProd = max(maxProd, len(words[i])*len(words[j]))
			}
		}
	}

	return maxProd
}

func sortWord(w string) string {
	r := []rune(w)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
func main() {
	println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}
