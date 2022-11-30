package main

import "math"

func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool)
	dp := make([]bool, len(s)+1)
	maxLen := math.MinInt

	for _, w := range wordDict {
		set[w] = true
		maxLen = max(maxLen, len(w))
	}

	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if (i - j) > maxLen {
				continue
			}
			dp[i] = dp[j] && set[s[j:i]]
			if dp[i] {
				break
			}
		}
	}

	return dp[len(s)]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	println(wordBreak("applepenapple", []string{"apple", "pen"}))
	println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
