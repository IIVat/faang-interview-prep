package main

// space O(m*n)
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := fill(len1+1, len2+1)

	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {

			if i == 0 || j == 0 {
				dp[i][j] = i + j
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
			}
		}

	}

	return dp[len1][len2]
}

//space O(n)

func minDistanceN(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	dp := make([]int, len1+1)

	for i := 0; i < len1+1; i++ {
		dp[i] = i
	}

	for j := 1; j < len2+1; j++ {
		last := dp[0]
		dp[0] = j
		for i := 1; i < len1+1; i++ {
			current := dp[i]
			if word1[i-1] == word2[j-1] {
				dp[i] = last
			} else {
				dp[i] = min(dp[i-1], dp[i]) + 1
			}
			last = current
		}

	}

	for i := 0; i < len1+1; i++ {
		println(dp[i])
	}

	return dp[len1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func fill(len1 int, len2 int) [][]int {
	dp := make([][]int, len1)

	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}

	return dp
}

func main() {
	println(minDistanceN("leetcode", "etco"))
	// println(minDistance("sea", "eat"))
	// println(minDistanceN("a", "b"))
}
