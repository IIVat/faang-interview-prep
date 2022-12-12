package main

import (
	"math"
)

func minWindow(s, t string) string {
	rCount, window := make(map[byte]int, 0), make(map[byte]int, 0)

	for idx, _ := range t {
		rCount[t[idx]]++
	}

	current, required := 0, len(rCount)

	// Setting up a variable containing the result's starting and ending point with default values and a length variable
	res, resLen := []int{-1, -1}, math.MaxInt32

	l := 0

	for r, _ := range s {

		ch := s[r]
		window[ch]++

		if _, ok := rCount[ch]; ok && window[ch] == rCount[ch] {
			current++
		}

		// Sliding Window in working
		for current == required {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = (r - l + 1)
			}

			window[s[l]] -= 1
			if _, ok := rCount[s[l]]; ok && window[s[l]] < rCount[s[l]] {
				current--
			}
			l++
		}
	}

	l, r := res[0], res[1]

	if resLen == math.MaxInt32 {
		return ""
	}

	return s[l : r+1]
}

func main() {
	println(minWindow("ABXYZJKLSNFC", "ABC"))
}
